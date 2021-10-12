package billing

import (
	"context"
	"github.com/deniskelin/billing-gokit/internal/service/billing"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	pb "gitlab.tada.team/tada-back/billing/proto/billing-gw/pb"
	"strconv"
)

type Endpoints struct {
	CreatePersonalAccount endpoint.Endpoint
}

func MakeEndpoints(s billing.IService) Endpoints {
	return Endpoints{
		CreatePersonalAccount: makeCreatePersonalAccountEndpoint(s),
	}
}

func makeCreatePersonalAccountEndpoint(s billing.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*pb.CreatePersonalAccountRequest) // type assertion
		if !ok {
			return nil, errors.New("wrong assertion type")
		}
		personalAccount, err := s.CreatePersonalAccount(ctx, req.GetAccountId())
		if err != nil {
			return nil, err
		}
		return &pb.CreatePersonalAccountResponse{
			Success:   true,
			AccountId: strconv.Itoa(int(personalAccount.Id)),
		}, nil
	}
}
