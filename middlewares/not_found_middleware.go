package middlewares

import (
	commonServices "main/services/common"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"

	templates "github.com/punk-link/gin-generic-http-templates"
)

type NotFoundMiddleware struct {
	templateDataService commonServices.TemplateDataServer
}

func NewNotFoundMiddleware(injector *do.Injector) (NotFoundHandler, error) {
	templateDataService := do.MustInvoke[commonServices.TemplateDataServer](injector)

	return &NotFoundMiddleware{
		templateDataService: templateDataService,
	}, nil
}

func (t *NotFoundMiddleware) HandlePageNotFound() gin.HandlerFunc {
	data := t.templateDataService.Enrich("Page not found", make(map[string]any))

	return func(ctx *gin.Context) {
		templates.NotFoundTemplate(ctx, "global/404.go.tmpl", data)
	}
}
