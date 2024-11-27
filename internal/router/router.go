package router

import (
	"github.com/am1macdonald/sevenDice/internal/database"
	"github.com/am1macdonald/sevenDice/internal/mailer"
	"github.com/am1macdonald/sevenDice/internal/templates"
	"github.com/valkey-io/valkey-go"
	"net/http"
	"strings"
)

type Router struct {
	cache           *valkey.Client
	db              *database.Queries
	mailer          *mailer.Mailer
	templateHandler *templates.TemplateHandler
	mux             *http.ServeMux
}

func NewRouter(cache *valkey.Client, db *database.Queries, mailer *mailer.Mailer, t *templates.TemplateHandler) *Router {
	return &Router{
		cache:           cache,
		db:              db,
		mailer:          mailer,
		templateHandler: t,
		mux:             http.NewServeMux(),
	}
}

func (router *Router) HandleRoute(w http.ResponseWriter, r *http.Request) {
	isApi := strings.Split(strings.TrimLeft(r.URL.Path, "/"), "/")[0] == "api"
	println(isApi)
	if isApi {
		router.handleAPIRoutes(w, r)
		return
	}
	router.handlePageRoutes(w, r)
}
