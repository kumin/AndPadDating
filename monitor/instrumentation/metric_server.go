package instrumentation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fvbock/endless"
	"github.com/kumin/BityDating/configs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

type MetricServer struct {
	port int
}

func NewMetricServer(
	cfg *configs.MetricServerCfg,
) *MetricServer {
	return &MetricServer{
		port: cfg.Port,
	}
}

func (m *MetricServer) Start(ctx context.Context) error {
	log.Info().Msgf("Metric server is startd on port %d", m.port)
	http.Handle("/metrics", promhttp.Handler())
	return endless.ListenAndServe(fmt.Sprintf(":%d", m.port), nil)
}
