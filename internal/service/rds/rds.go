package rds

import (
	"context"

	"github.com/deniskelin/billing-gokit/internal/config"

	"github.com/deniskelin/billing-gokit/pkg/cache"
	"github.com/deniskelin/billing-gokit/pkg/rds"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

// IService provides RDS service.
type IService interface {
	GetBillingByEnvAndID(ctx context.Context, accountID uint64, env string) ([]string, error)
	GetBillingEnvRouteByAccountID(ctx context.Context, accountID uint64) ([]rds.BillingInfo, error)
	SetRouteLabel(ctx context.Context, accountID uint64, envID, billingSource, routeLabel string) error
	GetRouteLabel(ctx context.Context, accountID uint64, envID, billingSource string) ([]rds.RouteLabel, error)
}

type Service struct {
	logger    zerolog.Logger
	rdsClient rds.IClient
	cache     cache.ICache
	appConfig *config.Configuration
}

func NewService(logger zerolog.Logger, rwdb rds.DBWriter, rdb rds.DBReader, iCache cache.ICache, appConfig *config.Configuration) IService {
	return &Service{
		logger:    logger,
		rdsClient: rds.NewClient(rwdb, rdb),
		cache:     iCache,
		appConfig: appConfig,
	}
}

func (ls *Service) GetBillingByEnvAndID(ctx context.Context, accountID uint64, envID string) ([]string, error) {
	reqID, _ := ctx.Value(middleware.RequestIDKey).(string)
	logger := ls.logger.With().Str("request-id", reqID).Logger()
	// todo ADD CACHE
	result, err := ls.rdsClient.GetBillingByEnvAndID(accountID, envID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get billing by env and id")
		return nil, err
	}
	return result, nil
}

func (ls *Service) GetBillingEnvRouteByAccountID(ctx context.Context, accountID uint64) ([]rds.BillingInfo, error) {
	reqID, _ := ctx.Value(middleware.RequestIDKey).(string)
	logger := ls.logger.With().Str("request-id", reqID).Logger()
	// todo ADD CACHE
	result, err := ls.rdsClient.GetBillingEnvRouteByAccountID(accountID, nil)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get billing, env and route by account ID")
		return nil, err
	}
	return result, nil
}

func (ls *Service) SetRouteLabel(ctx context.Context, accountID uint64, envID, billingSource, routeLabel string) error {
	reqID, _ := ctx.Value(middleware.RequestIDKey).(string)
	logger := ls.logger.With().Str("request-id", reqID).Logger()
	// todo ADD CACHE FLUSH
	err := ls.rdsClient.SetRouteLabel(accountID, envID, billingSource, routeLabel)
	if err != nil {
		logger.Error().Err(err).Msg("failed to set route label")
		return err
	}
	return nil
}

func (ls *Service) GetRouteLabel(ctx context.Context, accountID uint64, envID, billingSource string) ([]rds.RouteLabel, error) {
	reqID, _ := ctx.Value(middleware.RequestIDKey).(string)
	logger := ls.logger.With().Str("request-id", reqID).Logger()
	// todo ADD CACHE
	rl, err := ls.rdsClient.GetRouteLabel(accountID, envID, billingSource)
	if err != nil {
		logger.Error().Err(err).Msg("failed to set route label")
		return nil, err
	}
	return rl, nil
}
