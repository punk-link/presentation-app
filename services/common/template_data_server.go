package common

type TemplateDataServer interface {
	AddRequestUrl(pageUrl string, data map[string]any) map[string]any
	Enrich(pageTitle string, data map[string]any) map[string]any
}
