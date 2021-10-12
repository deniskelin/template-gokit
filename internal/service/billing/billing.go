package billing

import (
	"context"
	"github.com/deniskelin/billing-gokit/internal/config"
	log "github.com/deniskelin/billing-gokit/pkg/logger"
	"github.com/go-chi/chi/v5/middleware"
	pb_billing_api "gitlab.tada.team/tada-back/billing/proto/billing-api/pb"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type IService interface {
	CreatePersonalAccount(ctx context.Context, accountID string) (*pb_billing_api.CreatePersonalAccountResponse, error)
}

type Service struct {
	logger    *zerolog.Logger
	appConfig *config.Configuration
	client    pb_billing_api.PersonalAccountClient
}

func NewService(cfg *config.Configuration, logger *zerolog.Logger) IService {
	connection, err := grpc.Dial(cfg.Billing.ServiceGrpcConnectionString,
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(cfg.GRPC.MaxRequestBodySize),
			grpc.MaxCallSendMsgSize(cfg.GRPC.MaxRequestBodySize),
		), grpc.WithInsecure())
	if err != nil {
		logger.Fatal().Err(err).Msg("connection failed")
	}
	billingProtoClient := pb_billing_api.NewPersonalAccountClient(connection)

	return &Service{
		logger:    logger,
		appConfig: cfg,
		client:    billingProtoClient,
	}
}

func (bs *Service) CreatePersonalAccount(ctx context.Context, accountID string) (*pb_billing_api.CreatePersonalAccountResponse, error) {
	reqID, _ := ctx.Value(middleware.RequestIDKey).(string)
	logger := bs.logger.With().Str("request-id", reqID).Logger()

	request := &pb_billing_api.CreatePersonalAccountRequest{
		Name: accountID,
	}
	logger.With()
	logger = log.NewFieldLogger(logger, "request", request)

	response, err := bs.client.CreatePersonalAccount(ctx, request)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get billing by env and id")
		return nil, err
	}
	logger.Info().Interface("response", response).Msg("success")
	return response, nil
}
