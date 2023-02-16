package middlewares

import "github.com/gin-gonic/gin"

type MetricsHandler interface {
	HandleMetrics(instrumentationName string) gin.HandlerFunc
}
