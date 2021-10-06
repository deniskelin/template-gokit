package system

import (
	"context"
	"os"
	"runtime"
	"time"

	"github.com/deniskelin/billing-gokit/pkg/version"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

// IService provides RDS service.
type IService interface {
	APIVersion(ctx context.Context) (*SystemAPIVersionResponse, error)
	Info(ctx context.Context) (*SystemInfoResponse, error)
}

type Service struct {
	logger zerolog.Logger
}

func NewService(logger zerolog.Logger) IService {
	return &Service{logger: logger}
}

func (s *Service) APIVersion(ctx context.Context) (*SystemAPIVersionResponse, error) {
	reqID, _ := ctx.Value(middleware.RequestIDKey).(string)
	s.logger.Info().Str("request-id", reqID).Msg("request processed")
	return &SystemAPIVersionResponse{
		BuildDate:   version.AppVersion.BuildDate,
		BuildNumber: version.AppVersion.Build,
		Version:     version.AppVersion.Version,
		Hash:        version.AppVersion.CommitHash,
	}, nil
}

func (s *Service) Info(ctx context.Context) (*SystemInfoResponse, error) {
	reqID, _ := ctx.Value(middleware.RequestIDKey).(string)
	s.logger.Info().Str("request-id", reqID).Msg("request processed")

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	hname, err := os.Hostname()
	if err != nil {
		hname = "unknown"
	}

	return &SystemInfoResponse{
		Hostname:                     hname,
		ServerTime:                   time.Now(),
		UserAgent:                    "",
		RuntimeNumGoroutine:          int64(runtime.NumGoroutine()),
		RuntimeMemStatsAlloc:         memStats.Alloc,
		RuntimeMemStatsTotalAlloc:    memStats.TotalAlloc,
		RuntimeMemStatsSys:           memStats.Sys,
		RuntimeMemStatsLookups:       memStats.Lookups,
		RuntimeMemStatsMallocs:       memStats.Mallocs,
		RuntimeMemStatsFrees:         memStats.Frees,
		RuntimeMemStatsHeapAlloc:     memStats.HeapAlloc,
		RuntimeMemStatsHeapSys:       memStats.HeapSys,
		RuntimeMemStatsHeapIdle:      memStats.HeapIdle,
		RuntimeMemStatsHeapInUse:     memStats.HeapInuse,
		RuntimeMemStatsHeapReleased:  memStats.HeapReleased,
		RuntimeMemStatsHeapObjects:   memStats.HeapObjects,
		RuntimeMemStatsStackInuse:    memStats.StackInuse,
		RuntimeMemStatsStackSys:      memStats.StackSys,
		RuntimeMemStatsMSpanInuse:    memStats.MSpanInuse,
		RuntimeMemStatsMSpanSys:      memStats.MSpanSys,
		RuntimeMemStatsMCacheInuse:   memStats.MCacheInuse,
		RuntimeMemStatsMCacheSys:     memStats.MCacheSys,
		RuntimeMemStatsBuckHashSys:   memStats.BuckHashSys,
		RuntimeMemStatsGCSys:         memStats.GCSys,
		RuntimeMemStatsOtherSys:      memStats.OtherSys,
		RuntimeMemStatsNextGC:        memStats.NextGC,
		RuntimeMemStatsLastGC:        memStats.LastGC,
		RuntimeMemStatsPauseTotalNS:  memStats.PauseTotalNs,
		RuntimeMemStatsNumGC:         memStats.NumGC,
		RuntimeMemStatsNumForcedGC:   memStats.NumForcedGC,
		RuntimeMemStatsGCCPUFraction: memStats.GCCPUFraction,
		RuntimeMemStatsEnableGC:      memStats.EnableGC,
		RuntimeMemStatsDebugGC:       memStats.DebugGC,
		RuntimeNumCPU:                int64(runtime.NumCPU()),
		RequestID:                    reqID,
		RemoteAddr:                   "",
		RequestHeaders:               nil,
	}, nil
}
