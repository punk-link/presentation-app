package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRequestUrl(ctx *gin.Context) string {
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}

	return fmt.Sprintf("%s://%s", scheme, ctx.Request.Host+ctx.Request.RequestURI)
}
