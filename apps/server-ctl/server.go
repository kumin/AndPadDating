package apps

import (
	"context"
	"fmt"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/kumin/BityDating/configs"
	"github.com/kumin/BityDating/handler/http/middleware"
	http_handler "github.com/kumin/BityDating/handler/http/v1"
	"github.com/kumin/BityDating/monitor"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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
	albumHandler *http_handler.AlbumHandler,
	walletHandler *http_handler.WalletHandler,
) (*HttpServer, error) {
	// instrument
	latencyHistogram, err := monitor.LatencyHistorgram()
	if err != nil {
		return nil, err
	}
	router := gin.Default()
	router.Use(otelgin.Middleware("bity_dating"))
	// UserAPI
	userGroup := router.Group("/v1/user", middleware.ValidateToken())
	// userGroup.POST("", userHandler.CreateUser)
	userGroup.GET("/:id", userHandler.GetUser)
	userGroup.PUT("/:id", userHandler.UpdateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)
	userGroup.POST("/:id/avatar", userHandler.SetAvatar)

	// Album API
	albumGroup := router.Group(
		"/v1/album",
		middleware.MeterAPI(latencyHistogram),
		middleware.ValidateToken(),
	)
	albumGroup.POST("/upone", albumHandler.CreateOne)
	albumGroup.POST("/upmany", albumHandler.CreateMany)
	albumGroup.GET("/useralbum", albumHandler.GetUserAlbum)

	// MatchingAPI
	matchingGroup := router.Group("/v1/matching", middleware.ValidateToken())
	matchingGroup.POST("", matchingHandler.CreateMatching)
	matchingGroup.GET("/whoilike/:userid", matchingHandler.WhoILike)
	matchingGroup.GET("/wholikeme/:userid", matchingHandler.WhoLikeMe)
	matchingGroup.GET("/list/:userid", matchingHandler.ListMatching)

	// FeedAPI
	feedGroup := router.Group("/v1/feed", middleware.ValidateToken())
	feedGroup.GET("/:userid", feedHandler.GetFeed)

	// WalletAPI
	walletGroup := router.Group(
		"/v1/wallet",
		middleware.MeterAPI(latencyHistogram),
		middleware.ValidateToken(),
	)
	walletGroup.POST("", walletHandler.CreateTransaction)
	walletGroup.GET("list/:userid", walletHandler.ListTransactions)

	// AuthAPI
	authGroup := router.Group("/v1/auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)

	return &HttpServer{
		port:   configs.Port,
		server: router,
	}, nil
}

func (h *HttpServer) Start(ctx context.Context) error {
	log.Info().Msgf("Server is listening on port:%d", h.port)
	otelHandler := otelhttp.NewHandler(h.server, "/",
		otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents))
	return endless.ListenAndServe(fmt.Sprintf(":%d", h.port), otelHandler)
}
