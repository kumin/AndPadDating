package apps

import (
	"github.com/google/wire"
	"github.com/kumin/BityDating/configs"
	http_handler "github.com/kumin/BityDating/handler/http/v1"
	"github.com/kumin/BityDating/repos/provider"
	"github.com/kumin/BityDating/services"
)

var ServerGraphSet = wire.NewSet(
	configs.ConfigGraphSet,
	http_handler.HttpHandlerGraphSet,
	services.ServiceGraphSet,
	provider.MysqlGraphSet,
	NewHttpServer,
)
