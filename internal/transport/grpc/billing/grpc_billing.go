package billing

import (
	"context"
	billingEp "github.com/deniskelin/billing-gokit/internal/endpoint/billing"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	pb "gitlab.tada.team/tada-back/billing/proto/billing-gw/pb"
)

type RPCServer struct {
	createPersonalAccount kitgrpc.Handler
	pb.UnsafeBillingAPIServer
}

func NewServer(endpoints billingEp.Endpoints, options []kitgrpc.ServerOption) pb.BillingAPIServer {
	return &RPCServer{
		createPersonalAccount: kitgrpc.NewServer(endpoints.CreatePersonalAccount, decodeRequest, encodeResponse, options...),
	}
}

func (s *RPCServer) CreatePersonalAccount(ctx context.Context, req *pb.CreatePersonalAccountRequest) (*pb.CreatePersonalAccountResponse, error) {
	_, resp, err := s.createPersonalAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreatePersonalAccountResponse), nil
}

func decodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func encodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
