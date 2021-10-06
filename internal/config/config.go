package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-envconfig"
)

const (
	HeaderContentTypeKey  = "Content-Type"
	HeaderContentTypeJSON = "application/json; charset=utf-8"
	ServerName            = "Not all heroes wear capes 💪"
)

func NewConfig() (*Configuration, error) {
	appEnv := os.Getenv("APP_ENV")
	_ = godotenv.Load(".env", fmt.Sprintf(".env.%s", appEnv))
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
	Log         LogConfig         `env:",prefix=LOG_"`
	Runtime     RuntimeConfig     `env:",prefix=RUNTIME_"`
	Cache       CacheConfig       `env:",prefix=CAHCE_"`
	Listen      ListenConfig      `env:",prefix=LISTEN_"`
	HTTP        HTTPConfig        `env:",prefix=HTTP_"`
	GRPC        GRPCConfig        `env:",prefix=GRPC_"`
	HealthCheck HealthCheckConfig `env:",prefix=HEALTHCHECK_"`
	Metrics     MetricsConfig     `env:",prefix=HEALTHCHECK_"`
	RWDB        DBConfig          `env:",prefix=RWDB_"`
	RDB         DBConfig          `env:",prefix=RDB_"`
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

type CacheConfig struct {
	Type             string        `env:"TYPE,default=dummy"`
	ConnectionString string        `env:"CONNECTION_STRING"` // ex. redis://<user>:<password>@<host>:<port>/<db_number>
	TTL              time.Duration `env:"TTL,default=60s"`
}

type ListenConfig struct {
	Network string `env:"NETWORK,default=tcp"`
	Address string `env:"ADDRESS,default=:8080"`
}

type HTTPConfig struct {
	CORSEnabled                bool          `env:"CORS_ENABLED,default=false"`
	RequestLoggingEnabled      bool          `env:"REQUEST_LOGGING_ENABLED,default=false"`
	ResponseTimeLoggingEnabled bool          `env:"RESPONSE_TIME_LOGGING_ENABLED,default=false"`
	ReadTimeout                time.Duration `env:"READ_TIMEOUT,default=30s"`
	WriteTimeout               time.Duration `env:"WRITE_TIMEOUT,default=30s"`
	IdleTimeout                time.Duration `env:"IDLE_TIMEOUT,default=30s"`
	MaxRequestBodySize         int           `env:"MAX_REQUEST_BODY_SIZE,default=4194304"`
}

type GRPCConfig struct {
	RequestLoggingEnabled      bool          `env:"REQUEST_LOGGING_ENABLED,default=false"`
	ResponseTimeLoggingEnabled bool          `env:"RESPONSE_TIME_LOGGING_ENABLED,default=false"`
	ReadTimeout                time.Duration `env:"READ_TIMEOUT,default=30s"`
	WriteTimeout               time.Duration `env:"WRITE_TIMEOUT,default=30s"`
	IdleTimeout                time.Duration `env:"IDLE_TIMEOUT,default=30s"`
	MaxRequestBodySize         int           `env:"MAX_REQUEST_BODY_SIZE,default=4194304"`
}

type HealthCheckConfig struct {
	GoroutineThreshold int `env:"GOROUTINE_THRESHOLD,default=20"`
}

type MetricsConfig struct {
	Path      string `env:"PATH,default=/metrics"`
	Namespace string `env:"NAMESPACE,default=rds"`
}

type DBConfig struct {
	DriverName               string        `env:"DRIVER_NAME,required"` // postgres || pgx
	ConnectionString         string        `env:"CONNECTION_STRING,required"`
	MaxOpenConnection        int           `env:"MAX_OPEN_CONNECTION,default=25"`
	MaxIdleConnection        int           `env:"MAX_IDLE_CONNECTION,default=5"`
	MaxIdleConnectionTimeout time.Duration `env:"MAX_IDLE_TIMEOUT,default=300s"`
}
