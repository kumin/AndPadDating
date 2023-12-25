package instrumentation

import (
	"github.com/google/wire"
	"github.com/kumin/BityDating/configs"
)

var GraphSet = wire.NewSet(
	configs.ConfigGraphSet,
	NewMetricServer,
)
