package main

import (
	"net"
	"net/http"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/deniskelin/billing-gokit/internal/config"
	epRDS "github.com/deniskelin/billing-gokit/internal/endpoint/rds"
	epSystem "github.com/deniskelin/billing-gokit/internal/endpoint/system"
	"github.com/deniskelin/billing-gokit/internal/healthchecker"
	svcRDS "github.com/deniskelin/billing-gokit/internal/service/rds"
	svcSystem "github.com/deniskelin/billing-gokit/internal/service/system"
	tpGRPC "github.com/deniskelin/billing-gokit/internal/transport/grpc"
	tpGRPCRDS "github.com/deniskelin/billing-gokit/internal/transport/grpc/rds"
	tpGRPCSys "github.com/deniskelin/billing-gokit/internal/transport/grpc/system"
	tpHTTP "github.com/deniskelin/billing-gokit/internal/transport/http"
	tpHTTPRDS "github.com/deniskelin/billing-gokit/internal/transport/http/rds"
	tpHTTPSys "github.com/deniskelin/billing-gokit/internal/transport/http/system"
	"github.com/deniskelin/billing-gokit/pkg/cache"
	"github.com/deniskelin/billing-gokit/pkg/cache/connector"
	"github.com/deniskelin/billing-gokit/pkg/rds"
	pbapisys "github.com/deniskelin/billing-gokit/proto/apistatus"
	pbrds "github.com/deniskelin/billing-gokit/proto/rds"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/heptiolabs/healthcheck"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/soheilhy/cmux"
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

func initCache(ctype, connectionString string) (cache.ICache, error) {
	return connector.NewCache(ctype, connectionString)
}

func initDBConnection(dbConfig *config.DBConfig) (rds.DB, error) {
	dbConnection, err := sqlx.Connect(dbConfig.DriverName, dbConfig.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer dbConnection.Close()

	dbConnection.SetMaxIdleConns(dbConfig.MaxIdleConnection)
	dbConnection.SetMaxOpenConns(dbConfig.MaxOpenConnection)
	dbConnection.SetConnMaxIdleTime(dbConfig.MaxIdleConnectionTimeout)

	return dbConnection, nil
}

func initRDSServiceEndpoint(rwdb, rdb rds.DB, iCache cache.ICache, appConfig *config.Configuration, apiLogger zerolog.Logger) epRDS.Endpoints {
	svcl := svcRDS.NewService(apiLogger, rwdb, rdb, iCache, appConfig)
	return epRDS.MakeEndpoints(svcl)
}

func initSystemServiceEndpoint(_ *config.Configuration, apiLogger zerolog.Logger) epSystem.Endpoints {
	svcs := svcSystem.NewService(apiLogger)
	return epSystem.MakeEndpoints(svcs)
}

func initCMux(l net.Listener, log zerolog.Logger, listenErr chan error) (cmux.CMux, net.Listener, net.Listener) {
	m := cmux.New(l)
	var grpcListener, httpListener net.Listener
	grpcListener = m.Match(cmux.HTTP2())
	//grpcListener = m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener = m.Match(cmux.HTTP1Fast())

	go func(m cmux.CMux, log zerolog.Logger, listenErr chan error) {
		//httpListener := m.Match(cmux.HTTP1Fast())
		log.Info().Msgf("cMux listener started")
		if err := m.Serve(); err != nil {
			listenErr <- err
		}
	}(m, log, listenErr)
	return m, grpcListener, httpListener
}

func initKitGRPC(_ *config.Configuration, l net.Listener, rdsEP epRDS.Endpoints, systemEP epSystem.Endpoints, netLogger zerolog.Logger, listenErr chan error) *googlegrpc.Server {
	//ocTracing := kitoc.HTTPServerTrace()
	//serverOptions := []kithttp.ServerOption{ocTracing}
	var serverOptions []kitgrpc.ServerOption
	grpcRDSServer := tpGRPCRDS.NewServer(rdsEP, serverOptions)
	grpcSysServer := tpGRPCSys.NewServer(systemEP)
	grpcServer := googlegrpc.NewServer()
	pbrds.RegisterRDSServer(grpcServer, grpcRDSServer)
	pbapisys.RegisterSystemServer(grpcServer, grpcSysServer)
	go tpGRPC.RunGRPCServer(grpcServer, l, netLogger, listenErr)
	time.Sleep(10 * time.Millisecond)
	return grpcServer
}

func initHTTPRouter(_ *config.Configuration) *chi.Mux {
	router := chi.NewRouter()

	// todo - replace with transport layer!
	router.Use(middleware.NoCache)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID) // todo change for custom
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	var pongResponse = []byte("pong")

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write(pongResponse)
		return
	})

	return router
}

func initHealthChecker(config *config.Configuration, router *chi.Mux) {
	healthChecker := healthchecker.NewHealthChecker()
	//healthChecker.GetHealthChecker().AddReadinessCheck("Net Listener Started", healthcheck.TCPDialCheck(config.NetListen.Address, 50*time.Millisecond))
	// Our app is not ready if we can't resolve our upstream dependency in DNS.
	healthChecker.GetHealthChecker().AddLivenessCheck("Net Listener Started", healthcheck.TCPDialCheck(config.Listen.Address, 50*time.Millisecond))
	// Our app is not happy if we've got more than 100 goroutines running.
	healthChecker.GetHealthChecker().AddLivenessCheck("Goroutine Threshold", healthcheck.GoroutineCountCheck(25))

	router.Mount("/", healthChecker.Handler())
}

func initMetrics(config *config.Configuration, router *chi.Mux) {
	router.Get(config.Metrics.Path, promhttp.Handler().ServeHTTP)
}

func initKitHTTP(appConfig *config.Configuration, l net.Listener, rdsEP epRDS.Endpoints, systemEP epSystem.Endpoints, netLogger zerolog.Logger, listenErr chan error, router *chi.Mux) *http.Server {
	//ocTracing := kitoc.HTTPServerTrace()
	//serverOptions := []kithttp.ServerOption{ocTracing}
	var serverOptions []kithttp.ServerOption

	router.Mount("/system", tpHTTPSys.NewServer(systemEP, serverOptions))
	router.Mount("/v1", tpHTTPRDS.NewServer(rdsEP, serverOptions))
	if webDebugEnabled {
		router.Mount("/dbg", ProfilerHandler())
	}
	//router.Post("/sendEvent", sendEventHandler.ServeHTTP)
	httpServer := &http.Server{
		Handler:      router,
		TLSConfig:    nil,
		ReadTimeout:  appConfig.HTTP.ReadTimeout,
		WriteTimeout: appConfig.HTTP.WriteTimeout,
		IdleTimeout:  appConfig.HTTP.IdleTimeout,
	}
	go tpHTTP.RunHTTPServer(httpServer, l, netLogger, listenErr)
	time.Sleep(10 * time.Millisecond)
	return httpServer
}
