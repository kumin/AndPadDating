package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kumin/BityDating/monitor"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func MeterAPI(histogram metric.Float64Histogram) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		latency := time.Since(start)
		histogram.Record(ctx.Request.Context(), float64(latency.Milliseconds()),
			metric.WithAttributes(
				attribute.Key(monitor.ApiUrl).String(ctx.Request.RequestURI),
				attribute.Key(monitor.StatusCode).Int(ctx.Writer.Status()),
			))
	}
}
