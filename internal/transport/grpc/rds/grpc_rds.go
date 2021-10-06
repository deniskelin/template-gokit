package rds

import (
	"context"

	rdEp "github.com/deniskelin/billing-gokit/internal/endpoint/rds"
	pb "github.com/deniskelin/billing-gokit/proto/rds"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type RPCServer struct {
	getBillingByEnvAndID          kitgrpc.Handler
	getBillingEnvRouteByAccountID kitgrpc.Handler
	setRouteLabel                 kitgrpc.Handler
	getRouteLabel                 kitgrpc.Handler
	//pb.UnimplementedRDSServer
	pb.UnsafeRDSServer
}

// NewServer initializes a new gRPC server
func NewServer(endpoints rdEp.Endpoints, options []kitgrpc.ServerOption) pb.RDSServer {
	// TODO ADD SERVER OPTIONS
	return &RPCServer{
		getBillingByEnvAndID:          kitgrpc.NewServer(endpoints.GetBillingByEnvAndID, decodeGetBillingByEnvAndIDRequest, encodeGetBillingByEnvAndIDResponse, options...),
		getBillingEnvRouteByAccountID: kitgrpc.NewServer(endpoints.GetBillingEnvRouteByAccountID, decodeGetBillingEnvRouteByAccountIDRequest, encodeGetBillingEnvRouteByAccountIDResponse, options...),
		setRouteLabel:                 kitgrpc.NewServer(endpoints.SetRouteLabel, decodeSetRouteLabelRequest, encodeSetRouteLabelResponse, options...),
		getRouteLabel:                 kitgrpc.NewServer(endpoints.GetRouteLabel, decodeGetRouteLabelRequest, encodeGetRouteLabelResponse, options...),
	}
}

func (s *RPCServer) GetBillingByEnvAndID(ctx context.Context, req *pb.GetBillingByEnvAndIDRequest) (*pb.GetBillingByEnvAndIDResponse, error) {
	_, resp, err := s.getBillingByEnvAndID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetBillingByEnvAndIDResponse), nil
}

func (s *RPCServer) GetBillingEnvRouteByAccountID(ctx context.Context, req *pb.GetBillingEnvRouteByAccountIDRequest) (*pb.GetBillingEnvRouteByAccountIDResponse, error) {
	_, resp, err := s.getBillingEnvRouteByAccountID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetBillingEnvRouteByAccountIDResponse), nil
}

func (s *RPCServer) SetRouteLabel(ctx context.Context, req *pb.SetRouteLabelRequest) (*pb.SetRouteLabelResponse, error) {
	_, resp, err := s.getBillingEnvRouteByAccountID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SetRouteLabelResponse), nil
}

func (s *RPCServer) GetRouteLabel(ctx context.Context, req *pb.GetRouteLabelRequest) (*pb.GetRouteLabelResponse, error) {
	_, resp, err := s.getBillingEnvRouteByAccountID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetRouteLabelResponse), nil
}

func decodeGetBillingByEnvAndIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	//req := request.(*pb.GetBillingByEnvAndIDRequest)
	return request, nil
}

func encodeGetBillingByEnvAndIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	//resp := response.(*pb.GetBillingByEnvAndIDResponse)
	return response, nil
}

func decodeGetBillingEnvRouteByAccountIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	//req := request.(*pb.GetBillingEnvRouteByAccountIDRequest)
	return request, nil
}

func encodeGetBillingEnvRouteByAccountIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	//resp := response.(*pb.GetBillingEnvRouteByAccountIDResponse)
	return response, nil
}

func decodeSetRouteLabelRequest(_ context.Context, request interface{}) (interface{}, error) {
	//req := request.(*pb.SetRouteLabelRequest)
	return request, nil
}

func encodeSetRouteLabelResponse(_ context.Context, response interface{}) (interface{}, error) {
	//resp := response.(*pb.SetRouteLabelResponse)
	return response, nil
}

func decodeGetRouteLabelRequest(_ context.Context, request interface{}) (interface{}, error) {
	//req := request.(*pb.GetRouteLabelRequest)
	return request, nil
}

func encodeGetRouteLabelResponse(_ context.Context, response interface{}) (interface{}, error) {
	//resp := response.(*pb.GetRouteLabelResponse)
	return response, nil
}
