package controllers

import (
	"github.com/gin-gonic/gin"
	templates "github.com/punk-link/gin-generic-http-templates"
	"github.com/samber/do"
)

type StatusController struct {
}

func NewStatusController(injector *do.Injector) (*StatusController, error) {
	return &StatusController{}, nil
}

func (t *StatusController) CheckHealth(ctx *gin.Context) {
	templates.Ok(ctx, "OK")
}
