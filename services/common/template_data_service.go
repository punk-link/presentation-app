package common

import "github.com/samber/do"

type TemplateDataService struct {
}

func NewTemplateDataService(injector *do.Injector) (*TemplateDataService, error) {
	return &TemplateDataService{}, nil
}

func (t *TemplateDataService) Enrich(pageTitle string, data map[string]any) map[string]any {
	data["PageTitle"] = pageTitle
	data["ManagementAppEndpoint"] = "https://app.punk.link"

	return data
}
