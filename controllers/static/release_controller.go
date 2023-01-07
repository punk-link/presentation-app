package static

import (
	artistServices "main/services/artists"

	templates "github.com/punk-link/gin-generic-http-templates"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type ReleaseController struct {
	service *artistServices.ReleaseService
}

func NewReleaseController(injector *do.Injector) (*ReleaseController, error) {
	service := do.MustInvoke[*artistServices.ReleaseService](injector)

	return &ReleaseController{
		service: service,
	}, nil
}

func (t *ReleaseController) Get(ctx *gin.Context) {
	hash := ctx.Param("hash")

	result, err := t.service.Get(hash)
	templates.OkOrNotFoundTemplate(ctx, "release.go.tmpl", result, err)
}
