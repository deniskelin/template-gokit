package rds

import (
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/deniskelin/billing-gokit/internal/endpoint/rds"
	"github.com/deniskelin/billing-gokit/internal/transport/http/common"
	pb "github.com/deniskelin/billing-gokit/proto/rds"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/valyala/bytebufferpool"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewServer initializes a new gRPC server
func NewServer(endpoints rds.Endpoints, options []kithttp.ServerOption) http.Handler {
	options = append(options, kithttp.ServerErrorEncoder(common.EncodeErrorResponse) /*, kithttp.ServerErrorHandler(logger)*/)
	r := chi.NewRouter()
	r.Use(middleware.NoCache)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID) // todo change for custom
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	// {envID:[0-9A-z-]+}

	r.Get("/envs/{envID}/accounts/{accountID}", kithttp.NewServer(endpoints.GetBillingByEnvAndID, decodeGetBillingByEnvAndIDRequest, common.EncodeResponse, options...).ServeHTTP)
	r.Get("/accounts/{accountID}", kithttp.NewServer(endpoints.GetBillingEnvRouteByAccountID, decodeGetBillingEnvRouteByAccountIDRequest, common.EncodeResponse, options...).ServeHTTP)
	r.Get("/route_label/{accountID}/{billingSource}/{envID}", kithttp.NewServer(endpoints.GetRouteLabel, decodeGetRouteLabelRequest, common.EncodeResponse, options...).ServeHTTP)
	r.Get("/route_label/{accountID}/{billingSource}", kithttp.NewServer(endpoints.GetRouteLabel, decodeGetRouteLabelRequest, common.EncodeResponse, options...).ServeHTTP)
	r.Post("/route_label", kithttp.NewServer(endpoints.SetRouteLabel, decodeSetRouteLabelRequest, common.EncodeResponse, options...).ServeHTTP)
	return r
}

func decodeGetBillingByEnvAndIDRequest(_ context.Context, r *http.Request) (interface{}, error) {

	accountID, err := strconv.Atoi(chi.URLParam(r, "accountID"))
	if err != nil {
		return nil, err
	}

	return &pb.GetBillingByEnvAndIDRequest{
		AccountId: uint64(accountID),
		Env:       chi.URLParam(r, "envID"),
	}, nil
}

func decodeGetBillingEnvRouteByAccountIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	accountID, err := strconv.Atoi(chi.URLParam(r, "accountID"))
	if err != nil {
		return nil, err
	}
	return &pb.GetBillingEnvRouteByAccountIDRequest{
		AccountId: uint64(accountID),
		Env:       nil,
	}, nil
}

func decodeGetRouteLabelRequest(_ context.Context, r *http.Request) (interface{}, error) {

	res := &pb.GetRouteLabelRequest{
		BillingSource: chi.URLParam(r, "billingSource"),
	}
	// {accountID}/{envID}/{billingSource}
	accountID, err := strconv.Atoi(chi.URLParam(r, "accountID"))
	if err != nil {
		return nil, err
	}
	res.AccountId = uint64(accountID)

	if chi.URLParam(r, "envID") != "" {
		*res.Env = chi.URLParam(r, "envID")
	}

	return res, nil
}

func decodeSetRouteLabelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := &pb.SetRouteLabelRequest{}
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	_, err := io.Copy(buf, r.Body) // buf.ReadFrom(r.Body)
	if err != nil {
		return nil, err
	}
	err = protojson.Unmarshal(buf.Bytes(), request)
	//err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, err
	}
	return request, nil
}
