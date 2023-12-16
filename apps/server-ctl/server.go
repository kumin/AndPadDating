package apps

import (
	"context"
	"fmt"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/kumin/BityDating/configs"
	"github.com/kumin/BityDating/handler/http/middleware"
	http_handler "github.com/kumin/BityDating/handler/http/v1"
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
	authHandler *http_handler.AuthHandler,
) *HttpServer {
	router := gin.Default()
	// UserAPI
	userGroup := router.Group("/v1/user", middleware.ValidateToken())
	//userGroup.POST("", userHandler.CreateUser)
	userGroup.GET("/:id", userHandler.GetUser)
	userGroup.PUT("/:id", userHandler.UpdateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)
	userGroup.POST("/:id/upload", userHandler.UploadFile)

	// MatchingAPI
	matchingGroup := router.Group("/v1/matching", middleware.ValidateToken())
	matchingGroup.POST("", middleware.ValidateToken(), matchingHandler.CreateMatching)
	matchingGroup.GET("/whoilike/:userid", matchingHandler.WhoILike)
	matchingGroup.GET("/wholikeme/:userid", matchingHandler.WhoLikeMe)
	matchingGroup.GET("/list/:userid", matchingHandler.ListMatching)

	//FeedAPI
	feedGroup := router.Group("/v1/feed", middleware.ValidateToken())
	feedGroup.GET("/:userid", feedHandler.GetFeed)

	//AuthAPI
	authGroup := router.Group("/v1/auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)

	return &HttpServer{
		port:   configs.Port,
		server: router,
	}
}

func (h *HttpServer) Start(ctx context.Context) error {
	log.Printf("Server is listening on port:%d", h.port)
	return endless.ListenAndServe(fmt.Sprintf("localhost:%d", h.port), h.server)
}
