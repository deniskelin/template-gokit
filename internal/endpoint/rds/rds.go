package rds

import (
	"context"

	"github.com/deniskelin/billing-gokit/internal/service/rds"
	pb "github.com/deniskelin/billing-gokit/proto/rds"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	GetBillingByEnvAndID          endpoint.Endpoint
	GetBillingEnvRouteByAccountID endpoint.Endpoint
	SetRouteLabel                 endpoint.Endpoint
	GetRouteLabel                 endpoint.Endpoint
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s rds.IService) Endpoints {
	return Endpoints{
		GetBillingByEnvAndID:          makeGetBillingByEnvAndIDEndpoint(s),
		GetBillingEnvRouteByAccountID: makeGetBillingEnvRouteByAccountIDEndpoint(s),
		SetRouteLabel:                 makeSetRouteLabelEndpoint(s),
		GetRouteLabel:                 makeGetRouteLabelEndpoint(s),
	}
}

func makeGetBillingByEnvAndIDEndpoint(s rds.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*pb.GetBillingByEnvAndIDRequest) // type assertion
		if !ok {
			return nil, errors.New("wrong assertion type")
		}
		billings, err := s.GetBillingByEnvAndID(ctx, req.GetAccountId(), req.GetEnv())
		if err != nil {
			return nil, err
		}
		return &pb.GetBillingByEnvAndIDResponse{
			Billing: billings,
		}, nil
	}
}

func makeGetBillingEnvRouteByAccountIDEndpoint(s rds.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*pb.GetBillingEnvRouteByAccountIDRequest) // type assertion
		if !ok {
			return nil, errors.New("wrong assertion type")
		}
		billings, err := s.GetBillingEnvRouteByAccountID(ctx, req.GetAccountId())
		if err != nil {
			return nil, err
		}
		result := make([]*pb.GetBillingEnvRouteByAccountIDResponseElement, 0, len(billings))
		for i := range billings {
			result = append(result, &pb.GetBillingEnvRouteByAccountIDResponseElement{
				AccountId:     billings[i].NumberCode,
				IAccount:      billings[i].IAccount,
				EnvId:         billings[i].EnvID,
				BillingSource: billings[i].BillingSource,
				RouteLabel:    billings[i].RouteLabel,
			})
		}
		return &pb.GetBillingEnvRouteByAccountIDResponse{
			BillingEnvRoute: result,
		}, nil
	}
}

func makeSetRouteLabelEndpoint(s rds.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*pb.SetRouteLabelRequest) // type assertion
		if !ok {
			return nil, errors.New("wrong assertion type")
		}
		err := s.SetRouteLabel(ctx, req.GetAccountId(), req.GetEnv(), req.GetBillingSource(), req.GetRouteLabel())
		if err != nil {
			return nil, err
		}

		return &pb.SetRouteLabelResponse{}, nil
	}
}

func makeGetRouteLabelEndpoint(s rds.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*pb.GetRouteLabelRequest) // type assertion
		if !ok {
			return nil, errors.New("wrong assertion type")
		}
		rl, err := s.GetRouteLabel(ctx, req.GetAccountId(), req.GetEnv(), req.GetBillingSource())
		if err != nil {
			return nil, err
		}
		result := make([]*pb.GetRouteLabelResponseElement, 0, len(rl))
		for i := range rl {
			result = append(result, &pb.GetRouteLabelResponseElement{
				AccountId:  rl[i].AccountID,
				RouteLabel: rl[i].RouteLabel,
			})
		}
		return &pb.GetRouteLabelResponse{
			RouteLabel: result,
		}, nil
	}
}
