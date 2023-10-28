package apps

import (
	"context"
	"fmt"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/kumin/AndPadDating/configs"
	"github.com/kumin/AndPadDating/handler/http/middleware"
	http_handler "github.com/kumin/AndPadDating/handler/http/v1"
	"github.com/rs/zerolog/log"
)

type HttpServer struct {
	port   int
	server *gin.Engine
}

func NewHttpServer(
	configs *configs.ServerConfiguration,
	userHandler *http_handler.UserHandler,
	matchingHandler *http_handler.MatchingHandler,
	feedHandler *http_handler.FeedHandler,
) *HttpServer {
	router := gin.Default()
	// UserAPI
	router.Use(middleware.ValidateLogin())
	userGroup := router.Group("/v1/user")
	userGroup.POST("", userHandler.CreateUser)
	userGroup.GET("/:id", userHandler.GetUser)
	userGroup.PUT("/:id", userHandler.UpdateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)

	// MatchingAPI
	matchingGroup := router.Group("/v1/matching")
	matchingGroup.POST("", matchingHandler.CreateMatching)
	matchingGroup.GET("/whoilike/:userid", matchingHandler.WhoILike)
	matchingGroup.GET("/wholikeme/:partnerid", matchingHandler.WhoLikeMe)
	matchingGroup.GET("/list/:userid", matchingHandler.ListMatching)

	//FeedAPI
	feedGroup := router.Group("v1/feed")
	feedGroup.GET("/:userid", feedHandler.GetFeed)

	return &HttpServer{
		port:   configs.Port,
		server: router,
	}
}

func (h *HttpServer) Start(ctx context.Context) error {
	log.Printf("Server is listening on port:%d", h.port)
	return endless.ListenAndServe(fmt.Sprintf("localhost:%d", h.port), h.server)
}
