package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-envconfig"
)

const (
	HeaderContentTypeKey  = "Content-Type"
	HeaderContentTypeJSON = "application/json; charset=utf-8"
)

func NewConfig() (*Configuration, error) {

	var envFiles []string
	if _, err := os.Stat(".env"); err == nil {
		log.Println("found .env file, adding it to env config files list")
		envFiles = append(envFiles, ".env")
	}
	if os.Getenv("APP_ENV") != "" {
		appEnvName := fmt.Sprintf(".env.%s", os.Getenv("APP_ENV"))
		if _, err := os.Stat(appEnvName); err == nil {
			log.Println("found", appEnvName, "file, adding it to env config files list")
			envFiles = append(envFiles, appEnvName)
		}
	}
	if len(envFiles) > 0 {
		err := godotenv.Overload(envFiles...)
		if err != nil {
			return nil, errors.Wrapf(err, "error while opening env config: %s", err)
		}
	}
	cfg := &Configuration{}
	ctx := context.Background()

	err := envconfig.Process(ctx, cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing env config: %s", err)
	}
	return cfg, nil
}

// Configuration is basic structure that contains configuration
type Configuration struct {
	Log         LogConfig            `env:",prefix=LOG_"`
	Runtime     RuntimeConfig        `env:",prefix=RUNTIME_"`
	HTTP        HTTPConfig           `env:",prefix=HTTP_"`
	GRPC        GRPCConfig           `env:",prefix=GRPC_"`
	HealthCheck HealthCheckConfig    `env:",prefix=HEALTHCHECK_"`
	Metrics     MetricsConfig        `env:",prefix=HEALTHCHECK_"`
	Billing     BillingServiceConfig `env:",prefix=BILLING_"`
}

type LogConfig struct {
	Level             string        `env:"LEVEL,default=info"`
	Batch             bool          `env:"BATCH,default=false"`
	BatchSize         int           `env:"BATCH_SIZE,default=1000"`
	BatchPollInterval time.Duration `env:"BATCH_POLL_INTERVAL,default=5s"`
}

type RuntimeConfig struct {
	UseCPUs    int `env:"USE_CPUS,default=0"`
	MaxThreads int `env:"MAX_THREADS,default=0"`
}

type ListenConfig struct {
	Network     string `env:"NETWORK,default=tcp"`
	HttpAddress string `env:"HTTP_ADDRESS,default=:8080"`
	GrpcAddress string `env:"GRPC_ADDRESS,default=:8081"`
}

type HTTPConfig struct {
	CORSEnabled                bool          `env:"CORS_ENABLED,default=false"`
	RequestLoggingEnabled      bool          `env:"REQUEST_LOGGING_ENABLED,default=false"`
	ResponseTimeLoggingEnabled bool          `env:"RESPONSE_TIME_LOGGING_ENABLED,default=false"`
	ReadTimeout                time.Duration `env:"READ_TIMEOUT,default=30s"`
	WriteTimeout               time.Duration `env:"WRITE_TIMEOUT,default=30s"`
	IdleTimeout                time.Duration `env:"IDLE_TIMEOUT,default=30s"`
	MaxRequestBodySize         int           `env:"MAX_REQUEST_BODY_SIZE,default=4194304"`
	Network                    string        `env:"NETWORK,default=tcp"`
	Address                    string        `env:"ADDRESS,default=:8080"`
}

type GRPCConfig struct {
	RequestLoggingEnabled      bool          `env:"REQUEST_LOGGING_ENABLED,default=false"`
	ResponseTimeLoggingEnabled bool          `env:"RESPONSE_TIME_LOGGING_ENABLED,default=false"`
	ReadTimeout                time.Duration `env:"READ_TIMEOUT,default=30s"`
	WriteTimeout               time.Duration `env:"WRITE_TIMEOUT,default=30s"`
	IdleTimeout                time.Duration `env:"IDLE_TIMEOUT,default=30s"`
	MaxRequestBodySize         int           `env:"MAX_REQUEST_BODY_SIZE,default=4194304"`
	Network                    string        `env:"NETWORK,default=tcp"`
	Address                    string        `env:"ADDRESS,default=:18080"`
}

type HealthCheckConfig struct {
	GoroutineThreshold int `env:"GOROUTINE_THRESHOLD,default=20"`
}

type MetricsConfig struct {
	Path      string `env:"PATH,default=/metrics"`
	Namespace string `env:"NAMESPACE,default=api-gw"`
}

type BillingServiceConfig struct {
	ServiceGrpcConnectionString string `env:"SERVICE_GRPC_CONNECTION_STRING"`
}
