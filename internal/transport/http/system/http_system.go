package system

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/deniskelin/billing-gokit/internal/endpoint/system"
	"github.com/deniskelin/billing-gokit/internal/transport/http/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/valyala/bytebufferpool"
	"gitlab.tada.team/tada-back/billing/proto/apistatus/pb"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewServer initializes a new gRPC server
func NewServer(endpoints system.Endpoints, options []kithttp.ServerOption) http.Handler {
	options = append(options, kithttp.ServerErrorEncoder(common.EncodeErrorResponse) /*, kithttp.ServerErrorHandler(logger)*/)
	r := chi.NewRouter()
	r.Use(middleware.NoCache)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID) // todo change for custom
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Get("/version", kithttp.NewServer(endpoints.APIVersion, decodeAPIVersionRequest, common.EncodeResponse, options...).ServeHTTP)
	r.Get("/info", kithttp.NewServer(endpoints.Info, decodeInfoRequest, common.EncodeResponse, options...).ServeHTTP)
	return r
}

func decodeAPIVersionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := &pb.APIVersionRequest{}
	buf := &bytes.Buffer{}
	_, err := io.Copy(buf, r.Body) // buf.ReadFrom(r.Body)
	if err != nil {
		return nil, err
	}
	if buf.Len() > 0 {
		err = protojson.Unmarshal(buf.Bytes(), request)
		//err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			return nil, err
		}
	}
	return request, nil
}

func decodeInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := &pb.InfoRequest{}
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	_, err := io.Copy(buf, r.Body) // buf.ReadFrom(r.Body)
	if err != nil {
		return nil, err
	}
	if buf.Len() > 0 {
		err = protojson.Unmarshal(buf.Bytes(), request)
		//err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			return nil, err
		}
	}
	return request, nil
}
