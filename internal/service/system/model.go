package system

import "time"

type SystemAPIVersionResponse struct {
	BuildDate   string
	BuildNumber string
	Version     string
	Hash        string
}

type SystemInfoResponse struct {
	Hostname                     string
	ServerTime                   time.Time
	UserAgent                    string
	RuntimeNumGoroutine          int64
	RuntimeMemStatsAlloc         uint64
	RuntimeMemStatsTotalAlloc    uint64
	RuntimeMemStatsSys           uint64
	RuntimeMemStatsLookups       uint64
	RuntimeMemStatsMallocs       uint64
	RuntimeMemStatsFrees         uint64
	RuntimeMemStatsHeapAlloc     uint64
	RuntimeMemStatsHeapSys       uint64
	RuntimeMemStatsHeapIdle      uint64
	RuntimeMemStatsHeapInUse     uint64
	RuntimeMemStatsHeapReleased  uint64
	RuntimeMemStatsHeapObjects   uint64
	RuntimeMemStatsStackInuse    uint64
	RuntimeMemStatsStackSys      uint64
	RuntimeMemStatsMSpanInuse    uint64
	RuntimeMemStatsMSpanSys      uint64
	RuntimeMemStatsMCacheInuse   uint64
	RuntimeMemStatsMCacheSys     uint64
	RuntimeMemStatsBuckHashSys   uint64
	RuntimeMemStatsGCSys         uint64
	RuntimeMemStatsOtherSys      uint64
	RuntimeMemStatsNextGC        uint64
	RuntimeMemStatsLastGC        uint64
	RuntimeMemStatsPauseTotalNS  uint64
	RuntimeMemStatsNumGC         uint32
	RuntimeMemStatsNumForcedGC   uint32
	RuntimeMemStatsGCCPUFraction float64
	RuntimeMemStatsEnableGC      bool
	RuntimeMemStatsDebugGC       bool
	RuntimeNumCPU                int64
	RequestID                    string
	RemoteAddr                   string
	RequestHeaders               map[string]string
}
