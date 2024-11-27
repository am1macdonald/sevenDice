package apiConfig

import (
	"github.com/am1macdonald/sevenDice/internal/router"
)

type ApiConfig struct {
	Router *router.Router
}

func New(router *router.Router) *ApiConfig {
	return &ApiConfig{
		Router: router,
	}
}
