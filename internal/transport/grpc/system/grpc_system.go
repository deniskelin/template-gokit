package system

import (
	"context"

	"github.com/deniskelin/billing-gokit/internal/endpoint/system"
	gt "github.com/go-kit/kit/transport/grpc"
	"gitlab.tada.team/tada-back/billing/proto/apistatus/pb"
)

type Server struct {
	systemAPIVersion gt.Handler
	systemInfo       gt.Handler
	pb.UnimplementedSystemServer
}

// NewServer initializes a new gRPC server
func NewServer(endpoints system.Endpoints) pb.SystemServer {
	return &Server{
		systemAPIVersion: gt.NewServer(endpoints.APIVersion, decodeAPIVersionRequest, encodeAPIVersionResponse),
		systemInfo:       gt.NewServer(endpoints.Info, decodeInfoRequest, encodeInfoResponse),
	}
}

func (s *Server) SystemAPIVersion(ctx context.Context, req *pb.APIVersionRequest) (*pb.APIVersionResponse, error) {
	_, resp, err := s.systemAPIVersion.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.APIVersionResponse), nil
}

func decodeAPIVersionRequest(_ context.Context, request interface{}) (interface{}, error) {
	//req := request.(*pb.SystemAPIVersionRequest)
	//return endpoints.SystemAPIVersionReq{NumA: req.NumA, NumB: req.NumB}, nil
	return request, nil
}

func encodeAPIVersionResponse(_ context.Context, response interface{}) (interface{}, error) {
	//resp := response.(endpoints.SystemAPIVersionResp)
	//return &pb.SystemAPIVersionResponse{Result: resp.Result}, nil
	return response, nil
}

func (s *Server) SystemInfo(ctx context.Context, req *pb.InfoRequest) (*pb.InfoResponse, error) {
	_, resp, err := s.systemInfo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.InfoResponse), nil
}

func decodeInfoRequest(_ context.Context, request interface{}) (interface{}, error) {
	//req := request.(*pb.SystemInfoRequest)
	//return endpoints.SystemInfoReq{NumA: req.NumA, NumB: req.NumB}, nil
	return request, nil
}

func encodeInfoResponse(_ context.Context, response interface{}) (interface{}, error) {
	//resp := response.(endpoints.SystemInfoResp)
	//return &pb.SystemInfoResponse{Result: resp.Result}, nil
	return response, nil
}
