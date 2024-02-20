package http_handler

import (
	"github.com/google/wire"
)

var HttpHandlerGraphSet = wire.NewSet(
	NewUserHandler,
	NewMatchingHandler,
	NewFeedHandler,
	NewAuthHandler,
	NewAlbumHandler,
	NewWalletHandler,
)
