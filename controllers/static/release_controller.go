package static

import (
	base "main/controllers"

	artistServices "main/services/artists"
	commonServices "main/services/common"

	templates "github.com/punk-link/gin-generic-http-templates"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type ReleaseController struct {
	dataService commonServices.TemplateDataServer
	service     artistServices.ReleaseServer
}

func NewReleaseController(injector *do.Injector) (*ReleaseController, error) {
	dataService := do.MustInvoke[commonServices.TemplateDataServer](injector)
	service := do.MustInvoke[artistServices.ReleaseServer](injector)

	return &ReleaseController{
		dataService: dataService,
		service:     service,
	}, nil
}

func (t *ReleaseController) Get(ctx *gin.Context) {
	hash := ctx.Param("hash")

	result, err := t.service.Get(hash)

	requestUrl := base.GetRequestUrl(ctx)
	enrichedResult := t.dataService.AddRequestUrl(requestUrl, result)
	templates.OkOrNotFoundTemplate(ctx, "release.go.tmpl", "global/404.go.tmpl", enrichedResult, err)
}
