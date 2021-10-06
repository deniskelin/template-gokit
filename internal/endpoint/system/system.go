package system

import (
	"context"
	"github.com/deniskelin/billing-gokit/internal/service/system"

	pb "github.com/deniskelin/billing-gokit/proto/apistatus"
	"github.com/deniskelin/billing-gokit/proto/rds"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	APIVersion endpoint.Endpoint
	Info       endpoint.Endpoint
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s system.IService) Endpoints {
	return Endpoints{
		APIVersion: makeAPIVersionEndpoint(s),
		Info:       makeInfoEndpoint(s),
	}
}

func makeAPIVersionEndpoint(s system.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//req, ok := request.(*pb.SystemAPIVersionRequest) // type assertion
		//if !ok {
		//	return nil, errors.New("wrong assertion type")
		//}
		apiVersion, err := s.APIVersion(ctx)
		if err != nil {
			return nil, err
		}
		return &pb.APIVersionResponse{
			BuildDate:   apiVersion.BuildDate,
			BuildNumber: apiVersion.BuildNumber,
			Version:     apiVersion.Version,
			Hash:        apiVersion.Hash,
		}, nil
	}
}

func makeInfoEndpoint(s system.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//req, ok := request.(*pb.SystemInfoRequest) // type assertion
		//if !ok {
		//	return nil, errors.New("wrong assertion type")
		//}
		sysInfo, err := s.Info(ctx)
		if err != nil {
			return nil, err
		}

		return &pb.InfoResponse{
			Hostname:                     sysInfo.Hostname,
			ServerTime:                   timestamppb.New(sysInfo.ServerTime),
			UserAgent:                    sysInfo.UserAgent,
			RuntimeNumGoroutine:          sysInfo.RuntimeNumGoroutine,
			RuntimeMemStatsAlloc:         sysInfo.RuntimeMemStatsAlloc,
			RuntimeMemStatsTotalAlloc:    sysInfo.RuntimeMemStatsTotalAlloc,
			RuntimeMemStatsSys:           sysInfo.RuntimeMemStatsSys,
			RuntimeMemStatsLookups:       sysInfo.RuntimeMemStatsLookups,
			RuntimeMemStatsMallocs:       sysInfo.RuntimeMemStatsMallocs,
			RuntimeMemStatsFrees:         sysInfo.RuntimeMemStatsFrees,
			RuntimeMemStatsHeapAlloc:     sysInfo.RuntimeMemStatsHeapAlloc,
			RuntimeMemStatsHeapSys:       sysInfo.RuntimeMemStatsHeapSys,
			RuntimeMemStatsHeapIdle:      sysInfo.RuntimeMemStatsHeapIdle,
			RuntimeMemStatsHeapInUse:     sysInfo.RuntimeMemStatsHeapInUse,
			RuntimeMemStatsHeapReleased:  sysInfo.RuntimeMemStatsHeapReleased,
			RuntimeMemStatsHeapObjects:   sysInfo.RuntimeMemStatsHeapObjects,
			RuntimeMemStatsStackInuse:    sysInfo.RuntimeMemStatsStackInuse,
			RuntimeMemStatsStackSys:      sysInfo.RuntimeMemStatsStackSys,
			RuntimeMemStatsMSpanInuse:    sysInfo.RuntimeMemStatsMSpanInuse,
			RuntimeMemStatsMSpanSys:      sysInfo.RuntimeMemStatsMSpanSys,
			RuntimeMemStatsMCacheInuse:   sysInfo.RuntimeMemStatsMCacheInuse,
			RuntimeMemStatsMCacheSys:     sysInfo.RuntimeMemStatsMCacheSys,
			RuntimeMemStatsBuckHashSys:   sysInfo.RuntimeMemStatsBuckHashSys,
			RuntimeMemStatsGCSys:         sysInfo.RuntimeMemStatsGCSys,
			RuntimeMemStatsOtherSys:      sysInfo.RuntimeMemStatsOtherSys,
			RuntimeMemStatsNextGC:        sysInfo.RuntimeMemStatsNextGC,
			RuntimeMemStatsLastGC:        sysInfo.RuntimeMemStatsLastGC,
			RuntimeMemStatsPauseTotalNs:  sysInfo.RuntimeMemStatsPauseTotalNS,
			RuntimeMemStatsNumGC:         sysInfo.RuntimeMemStatsNumGC,
			RuntimeMemStatsNumForcedGC:   sysInfo.RuntimeMemStatsNumForcedGC,
			RuntimeMemStatsGCCpuFraction: sysInfo.RuntimeMemStatsGCCPUFraction,
			RuntimeMemStatsEnableGC:      sysInfo.RuntimeMemStatsEnableGC,
			RuntimeMemStatsDebugGC:       sysInfo.RuntimeMemStatsDebugGC,
			RuntimeNumCpu:                sysInfo.RuntimeNumCPU,
			RequestId:                    sysInfo.RequestID,
			RemoteAddr:                   sysInfo.RemoteAddr,
			RequestHeaders:               sysInfo.RequestHeaders,
		}, nil
	}
}
