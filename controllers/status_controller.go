package controllers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	templates "github.com/punk-link/gin-generic-http-templates"
	contracts "github.com/punk-link/presentation-contracts"
	"github.com/samber/do"
)

type StatusController struct {
	grpcClient contracts.PresentationClient
}

func NewStatusController(injector *do.Injector) (*StatusController, error) {
	grpcClient := do.MustInvoke[contracts.PresentationClient](injector)

	return &StatusController{
		grpcClient: grpcClient,
	}, nil
}

func (t *StatusController) CheckHealth(ctx *gin.Context) {
	_, err := t.grpcClient.CheckHealth(context.Background(), &contracts.HealthCheckRequest{})
	if err != nil {
		templates.BadRequest(ctx, fmt.Sprintf("Degraded: %v", err))
		return
	}

	templates.Ok(ctx, "OK")
}
