package static

import (
	artistServices "main/services/artists"

	templates "github.com/punk-link/gin-generic-http-templates"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type ArtistController struct {
	service artistServices.ArtistServer
}

func NewArtistController(injector *do.Injector) (*ArtistController, error) {
	service := do.MustInvoke[artistServices.ArtistServer](injector)

	return &ArtistController{
		service: service,
	}, nil
}

func (t *ArtistController) Get(ctx *gin.Context) {
	hash := ctx.Param("hash")

	result, err := t.service.Get(hash)
	templates.OkOrNotFoundTemplate(ctx, "artist.go.tmpl", "global/404.go.tmpl", result, err)
}
