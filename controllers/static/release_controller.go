package static

import (
	artistServices "main/services/artists"

	templates "github.com/punk-link/gin-generic-http-templates"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type ReleaseController struct {
	service artistServices.ReleaseServer
}

func NewReleaseController(injector *do.Injector) (*ReleaseController, error) {
	service := do.MustInvoke[artistServices.ReleaseServer](injector)

	return &ReleaseController{
		service: service,
	}, nil
}

func (t *ReleaseController) Get(ctx *gin.Context) {
	hash := ctx.Param("hash")

	result, err := t.service.Get(hash)
	templates.OkOrNotFoundTemplate(ctx, "release.go.tmpl", "global/404.go.tmpl", result, err)
}
