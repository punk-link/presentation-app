package static

import (
	releaseServices "main/services/releases"

	base "github.com/punk-link/gin-generic-http-templates"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type ReleaseController struct {
	service *releaseServices.ReleaseService
}

func NewReleaseController(injector *do.Injector) (*ReleaseController, error) {
	service := do.MustInvoke[*releaseServices.ReleaseService](injector)

	return &ReleaseController{
		service: service,
	}, nil
}

func (t *ReleaseController) Get(ctx *gin.Context) {
	hash := ctx.Param("hash")

	result, err := t.service.Get(hash)
	base.OkOrNotFoundTemplate(ctx, "release.go.tmpl", result, err)
}
