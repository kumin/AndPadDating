//go:build wireinject
// +build wireinject

package httphandlertest

import (
	"github.com/google/wire"
	"github.com/kumin/BityDating/configs"
	http_handler "github.com/kumin/BityDating/handler/http/v1"
	"github.com/kumin/BityDating/repos/provider"
	"github.com/kumin/BityDating/services"
)

type HttpHandler struct {
	userHandler     *http_handler.UserHandler
	matchingHandler *http_handler.MatchingHandler
	authHandler     *http_handler.AuthHandler

	userService *services.UserService
}

func NewHttpHandler(
	userHandler *http_handler.UserHandler,
	matchingHandler *http_handler.MatchingHandler,
	authHandler *http_handler.AuthHandler,
	userService *services.UserService,
) *HttpHandler {
	return &HttpHandler{
		userHandler:     userHandler,
		matchingHandler: matchingHandler,
		authHandler:     authHandler,
		userService:     userService,
	}
}

var HandlerGraphSet = wire.NewSet(
	configs.ConfigGraphSet,
	services.ServiceGraphSet,
	provider.MysqlGraphSet,
	http_handler.HttpHandlerGraphSet,
	wire.NewSet(NewHttpHandler),
)

func BuildHttpHandler() (*HttpHandler, error) {
	wire.Build(HandlerGraphSet)
	return nil, nil
}
