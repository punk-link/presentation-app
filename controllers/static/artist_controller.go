package static

import (
	base "main/controllers"
	artistServices "main/services/artists"
	commonServices "main/services/common"

	templates "github.com/punk-link/gin-generic-http-templates"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type ArtistController struct {
	dataService commonServices.TemplateDataServer
	service     artistServices.ArtistServer
}

func NewArtistController(injector *do.Injector) (*ArtistController, error) {
	dataService := do.MustInvoke[commonServices.TemplateDataServer](injector)
	service := do.MustInvoke[artistServices.ArtistServer](injector)

	return &ArtistController{
		dataService: dataService,
		service:     service,
	}, nil
}

func (t *ArtistController) Get(ctx *gin.Context) {
	hash := ctx.Param("hash")

	result, err := t.service.Get(hash)

	requestUrl := base.GetRequestUrl(ctx)
	enrichedResult := t.dataService.AddRequestUrl(requestUrl, result)
	templates.OkOrNotFoundTemplate(ctx, "artist.go.tmpl", "global/404.go.tmpl", enrichedResult, err)
}
