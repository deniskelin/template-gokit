package main

import (
	"net"
	"net/http"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/deniskelin/billing-gokit/internal/config"
	epBilling "github.com/deniskelin/billing-gokit/internal/endpoint/billing"
	epSystem "github.com/deniskelin/billing-gokit/internal/endpoint/system"
	"github.com/deniskelin/billing-gokit/internal/healthchecker"
	svcBilling "github.com/deniskelin/billing-gokit/internal/service/billing"
	svcSystem "github.com/deniskelin/billing-gokit/internal/service/system"
	tpGRPC "github.com/deniskelin/billing-gokit/internal/transport/grpc"
	tpGRPCBilling "github.com/deniskelin/billing-gokit/internal/transport/grpc/billing"
	tpGRPCSys "github.com/deniskelin/billing-gokit/internal/transport/grpc/system"
	tpHTTP "github.com/deniskelin/billing-gokit/internal/transport/http"
	tpHTTPBilling "github.com/deniskelin/billing-gokit/internal/transport/http/billing"
	tpHTTPSys "github.com/deniskelin/billing-gokit/internal/transport/http/system"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	pbapisys "gitlab.tada.team/tada-back/billing/proto/apistatus/pb"
	pbBillingGW "gitlab.tada.team/tada-back/billing/proto/billing-gw/pb"
	googlegrpc "google.golang.org/grpc"

	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
)

// initRuntime initialize runtime params
func initRuntime(cpu, threads int, logger zerolog.Logger) {
	if cpu == 0 {
		cpu = runtime.NumCPU()
		runtime.GOMAXPROCS(runtime.NumCPU())
	} else {
		runtime.GOMAXPROCS(cpu)
	}
	logger.Info().Msgf("set to use %d CPUs", cpu)
	if threads == 0 {
		threads = 10000
	} else {
		debug.SetMaxThreads(threads)
	}
	logger.Info().Msgf("set to use maximum %d threads", threads)
}

func initBillingGWServiceEndpoint(appConfig *config.Configuration, apiLogger zerolog.Logger) epBilling.Endpoints {
	srv := svcBilling.NewService(appConfig, &apiLogger)
	return epBilling.MakeEndpoints(srv)
}

func initSystemServiceEndpoint(_ *config.Configuration, apiLogger zerolog.Logger) epSystem.Endpoints {
	svcs := svcSystem.NewService(apiLogger)
	return epSystem.MakeEndpoints(svcs)
}

func initKitGRPC(appConfig *config.Configuration, billingEP epBilling.Endpoints, systemEP epSystem.Endpoints, netLogger zerolog.Logger, listenErr chan error) (*googlegrpc.Server, net.Listener) {
	var serverOptions []kitgrpc.ServerOption
	grpcBillingServer := tpGRPCBilling.NewServer(billingEP, serverOptions)
	grpcSysServer := tpGRPCSys.NewServer(systemEP)
	grpcServer := googlegrpc.NewServer()
	pbBillingGW.RegisterBillingAPIServer(grpcServer, grpcBillingServer)
	pbapisys.RegisterSystemServer(grpcServer, grpcSysServer)

	listenerGrpc, err := net.Listen(appConfig.GRPC.Network, appConfig.GRPC.Address)
	if err != nil {
		netLogger.Fatal().Err(err).Msg("failed to init net.Listen for grpc")
	} else {
		netLogger.Info().Msg("successful net.Listen for grpc init")
	}

	go tpGRPC.RunGRPCServer(grpcServer, listenerGrpc, netLogger, listenErr)
	time.Sleep(10 * time.Millisecond)
	return grpcServer, listenerGrpc
}

func initHTTPRouter(_ *config.Configuration, apiLogger zerolog.Logger) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.NoCache)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	var pongResponse = []byte("pong")

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(pongResponse); err != nil {
			apiLogger.Fatal().Err(err).Msg("error of writing response")
		}
		return
	})

	return router
}

func initHealthChecker(config *config.Configuration, router *chi.Mux) {
	healthChecker := healthchecker.NewHealthChecker()
	//healthChecker.GetHealthChecker().AddReadinessCheck("Net Listener Started", healthcheck.TCPDialCheck(config.NetListen.Address, 50*time.Millisecond))
	// Our app is not ready if we can't resolve our upstream dependency in DNS.
	healthChecker.GetHealthChecker().AddLivenessCheck("HTTP Net Listener Started", healthcheck.TCPDialCheck(config.HTTP.Address, 50*time.Millisecond))
	healthChecker.GetHealthChecker().AddLivenessCheck("GRPC Net Listener Started", healthcheck.TCPDialCheck(config.GRPC.Address, 50*time.Millisecond))
	// Our app is not happy if we've got more than 100 goroutines running.
	healthChecker.GetHealthChecker().AddLivenessCheck("Goroutine Threshold", healthcheck.GoroutineCountCheck(25))

	router.Mount("/", healthChecker.Handler())
}

func initMetrics(config *config.Configuration, router *chi.Mux) {
	router.Get(config.Metrics.Path, promhttp.Handler().ServeHTTP)
}

func initKitHTTP(appConfig *config.Configuration, billingEP epBilling.Endpoints, systemEP epSystem.Endpoints, netLogger zerolog.Logger, listenErr chan error, router *chi.Mux) (*http.Server, net.Listener) {

	var serverOptions []kithttp.ServerOption

	router.Mount("/system", tpHTTPSys.NewServer(systemEP, serverOptions))
	router.Mount("/v1", tpHTTPBilling.NewServer(billingEP, serverOptions))
	if webDebugEnabled {
		router.Mount("/dbg", ProfilerHandler())
	}

	httpServer := &http.Server{
		Handler:      router,
		TLSConfig:    nil,
		ReadTimeout:  appConfig.HTTP.ReadTimeout,
		WriteTimeout: appConfig.HTTP.WriteTimeout,
		IdleTimeout:  appConfig.HTTP.IdleTimeout,
	}

	listenerHttp, err := net.Listen(appConfig.HTTP.Network, appConfig.HTTP.Address)
	if err != nil {
		netLogger.Fatal().Err(err).Msg("failed to init net.Listen for http")
	} else {
		netLogger.Info().Msg("successful net.Listen for http init")
	}

	go tpHTTP.RunHTTPServer(httpServer, listenerHttp, netLogger, listenErr)
	time.Sleep(10 * time.Millisecond)
	return httpServer, listenerHttp
}
