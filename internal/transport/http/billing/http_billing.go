package billing

import (
	"context"
	"github.com/deniskelin/billing-gokit/internal/endpoint/billing"
	"github.com/deniskelin/billing-gokit/internal/transport/http/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/valyala/bytebufferpool"
	pb "gitlab.tada.team/tada-back/billing/proto/billing-gw/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"net/http"
)

func NewServer(endpoints billing.Endpoints, options []kithttp.ServerOption) http.Handler {
	options = append(options, kithttp.ServerErrorEncoder(common.EncodeErrorResponse) /*, kithttp.ServerErrorHandler(logger)*/)
	r := chi.NewRouter()
	r.Use(middleware.NoCache)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	r.Route("/accounts", func(r chi.Router) {
		r.Post("/", kithttp.NewServer(endpoints.CreatePersonalAccount, decodeCreatePersonalAccountRequest, common.EncodeResponse, options...).ServeHTTP)

		//r.With(paginate).Get("/", listArticles)                           // GET /articles
		//r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017
		//
		//r.Post("/", createArticle)       // POST /articles
		//r.Get("/search", searchArticles) // GET /articles/search
		//
		//// Regexp url parameters:
		//r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug) // GET /articles/home-is-toronto
		//
		//// Subrouters:
		//r.Route("/{articleID}", func(r chi.Router) {
		//	r.Use(ArticleCtx)
		//	r.Get("/", getArticle)       // GET /articles/123
		//	r.Put("/", updateArticle)    // PUT /articles/123
		//	r.Delete("/", deleteArticle) // DELETE /articles/123
		//})

	})
	return r
}

func decodeCreatePersonalAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := &pb.CreatePersonalAccountRequest{}
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
