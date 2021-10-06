package common

import (
	"context"
	"net/http"

	"github.com/deniskelin/billing-gokit/internal/config"
	"github.com/goccy/go-json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Errorer interface {
	error() error
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(config.HeaderContentTypeKey, config.HeaderContentTypeJSON)
	if e, ok := response.(Errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		EncodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	if resp, ok := response.(proto.Message); ok {
		buf, err := protojson.Marshal(resp)
		if err != nil {
			return err
		}
		_, err = w.Write(buf)
		if err != nil {
			return err
		}
	} else {
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		if err := enc.EncodeWithOption(response, json.UnorderedMap()); err != nil {
			return err
		}
	}

	return nil
}

func EncodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set(config.HeaderContentTypeKey, config.HeaderContentTypeJSON)
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}

func codeFrom(err error) int {
	switch err {
	//case order.ErrOrderNotFound:
	//	return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
