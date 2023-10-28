package apps

import (
	"github.com/google/wire"
	"github.com/kumin/AndPadDating/configs"
	http_handler "github.com/kumin/AndPadDating/handler/http/v1"
	"github.com/kumin/AndPadDating/repos/provider"
	"github.com/kumin/AndPadDating/services"
)

var ServerGraphSet = wire.NewSet(
	configs.ConfigGraphSet,
	http_handler.HttpHandlerGraphSet,
	services.ServiceGraphSet,
	provider.MysqlGraphSet,
	NewHttpServer,
)
