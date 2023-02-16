package controllers

import (
	"fmt"
	commonServices "main/services/common"

	"github.com/gin-gonic/gin"
	templates "github.com/punk-link/gin-generic-http-templates"
	"github.com/samber/do"
)

type StatusController struct {
	service commonServices.HealthChecker
}

func NewStatusController(injector *do.Injector) (*StatusController, error) {
	service := do.MustInvoke[commonServices.HealthChecker](injector)

	return &StatusController{
		service: service,
	}, nil
}

func (t *StatusController) CheckHealth(ctx *gin.Context) {
	err := t.service.Check()
	if err != nil {
		templates.BadRequest(ctx, fmt.Sprintf("Degraded: %v", err))
		return
	}

	templates.Ok(ctx, "OK")
}

func (t *StatusController) Ok(ctx *gin.Context) {
	templates.Ok(ctx, "OK")
}
