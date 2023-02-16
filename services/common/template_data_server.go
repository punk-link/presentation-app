package common

type TemplateDataServer interface {
	Enrich(pageTitle string, data map[string]any) map[string]any
}
