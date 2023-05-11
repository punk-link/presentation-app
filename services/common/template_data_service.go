package common

import (
	"fmt"
	"main/constants"
	commonModels "main/models/common"

	"github.com/samber/do"
)

type TemplateDataService struct {
	options *commonModels.TemplateOptions
}

func NewTemplateDataService(injector *do.Injector) (TemplateDataServer, error) {
	options := do.MustInvoke[*commonModels.TemplateOptions](injector)
	return &TemplateDataService{
		options: options,
	}, nil
}

func (t *TemplateDataService) AddRequestUrl(pageUrl string, data map[string]any) map[string]any {
	data["PageUrl"] = pageUrl

	return data
}

func (t *TemplateDataService) Enrich(pageTitle string, data map[string]any) map[string]any {
	data["PageTitle"] = fmt.Sprintf("%s – %s", pageTitle, t.options.PageTitleTemplate)
	data["ManagementAppEndpoint"] = t.options.ManagementAppEndpoint
	data["SharableSocialNetworkIds"] = t.addSocialNetworkIds()

	return data
}

// TODO: move to contracts
func (t *TemplateDataService) addSocialNetworkIds() []string {
	return []string{
		constants.FACEBOOK,
		//constants.INSTAGRAM,
		constants.TELEGRAM,
		constants.TWITTER,
		constants.VK,
	}
}
