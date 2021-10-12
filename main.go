package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/deniskelin/billing-gokit/internal/config"
	"github.com/deniskelin/billing-gokit/pkg/logger"
	"github.com/deniskelin/billing-gokit/pkg/version"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func main() {
	appConfig, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	profiler := initDebugger()
	defer func() {
		stopDebugger(profiler)
	}()
	if webDebugEnabled {
		appConfig.Log.Level = logger.Debug
	}

	var baseLogger zerolog.Logger
	var loggerCloser io.WriteCloser
	if appConfig.Log.Batch {
		baseLogger, loggerCloser, err = logger.NewDiodeLogger(os.Stdout, appConfig.Log.Level, appConfig.Log.BatchSize, appConfig.Log.BatchPollInterval)
	} else {
		baseLogger, loggerCloser, err = logger.NewLogger(os.Stdout, appConfig.Log.Level)
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if loggerCloser != nil {
			err = loggerCloser.Close()
			if err != nil {
				log.Fatalf("error acquired while closing log writer: %+v", err)
			}
		}
	}()

	baseLogger = baseLogger.With().
		Str("app_version", version.AppVersion.Version).
		Str("app_build", version.AppVersion.Build).
		Str("app_commit_hash", version.AppVersion.CommitHash).Logger()

	apiLogger := logger.NewComponentLogger(baseLogger, "api", 2)
	coreLogger := logger.NewComponentLogger(baseLogger, "core", 2)
	netLogger := logger.NewComponentLogger(baseLogger, "net", 2)

	defer func() {
		coreLogger.Info().Msg("application stopped")
	}()

	coreLogger.Info().Msg("system initialization started")

	initRuntime(appConfig.Runtime.UseCPUs, appConfig.Runtime.MaxThreads, coreLogger)

	listenErr := make(chan error, 1)

	billingEP := initBillingGWServiceEndpoint(appConfig, apiLogger)
	systemEP := initSystemServiceEndpoint(appConfig, apiLogger)
	chiRouter := initHTTPRouter(appConfig, apiLogger)
	initMetrics(appConfig, chiRouter)
	initHealthChecker(appConfig, chiRouter)
	grpcServer, grpcListener := initKitGRPC(appConfig, billingEP, systemEP, netLogger, listenErr)
	defer func() {
		_ = grpcListener.Close()
	}()
	httpServer, httpListener := initKitHTTP(appConfig, billingEP, systemEP, netLogger, listenErr, chiRouter)
	defer func() {
		_ = httpListener.Close()
	}()
	runApp(grpcServer, httpServer, coreLogger, listenErr)

}

func runApp(grpcServer *grpc.Server, httpServer *http.Server, coreLogger zerolog.Logger, listenErr chan error) {

	var shutdownCh = make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var err error
	var runningApp = true

	for runningApp {
		select {
		// handle error channel
		case err = <-listenErr:
			if err != nil {
				// LogWithTrace(app.GetServerLogger(), err).Errorf("received grpc server error: %s", err)
				coreLogger.Error().Err(err).Msg("received listener error")
				shutdownCh <- os.Kill
			}
		// handle os system signal
		case sig := <-shutdownCh:
			coreLogger.Info().Msgf("shutdown signal received: %s", sig.String())
			// httpServerCancelFunc context.CancelFunc
			ctxTimeout, timeoutCancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
			err = httpServer.Shutdown(ctxTimeout) // may returns ErrServerClosed
			defer timeoutCancelFunc()
			if err != nil {
				coreLogger.Error().Err(err).Msg("received http Shutdown error")
			}
			grpcServer.GracefulStop()
			coreLogger.Info().Msg("server loop stopped")
			runningApp = false
			break
		}
	}
}
