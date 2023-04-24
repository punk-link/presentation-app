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
	// TODO: add a dynamic domain name to the title
	data["PageTitle"] = fmt.Sprintf("%s – Synonym.fans", pageTitle)
	data["ManagementAppEndpoint"] = t.options.ManagementAppEndpoint
	data["SocialNatworkIds"] = t.addSocialNetworkIds()

	return data
}

func (t *TemplateDataService) addSocialNetworkIds() map[string]any {
	return map[string]any{
		"Facebook":  constants.FACEBOOK,
		"Instagram": constants.INSTAGRAM,
		"Telegram":  constants.TELEGRAM,
		"Twitter":   constants.TWITTER,
		"Vk":        constants.VK,
	}
}
