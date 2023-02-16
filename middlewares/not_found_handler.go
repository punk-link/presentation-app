package middlewares

import "github.com/gin-gonic/gin"

type NotFoundHandler interface {
	HandlePageNotFound() gin.HandlerFunc
}
