package sharelinks

import (
	"fmt"
	"net/url"

	presentationContractConstants "github.com/punk-link/presentation-contracts/constants"
)

func Generate(socialNetworkId string, ref string, text string) string {
	switch socialNetworkId {
	case presentationContractConstants.FACEBOOK:
		return buildFacebookLink(ref)
	case presentationContractConstants.INSTAGRAM:
		return buildFacebookLink(ref)
	case presentationContractConstants.TELEGRAM:
		return buildTelegramLink(ref, text)
	case presentationContractConstants.TWITTER:
		return buildTwitterLink(ref, text)
	case presentationContractConstants.VK:
		return buildVkLink(ref, text)
	}

	return ref
}

func addTextPreview(ref string, text string) string {
	if text == "" {
		return ref
	}

	encodedText := url.QueryEscape(text)

	return fmt.Sprintf("%s&text=%s", ref, encodedText)
}

func buildFacebookLink(ref string) string {
	return fmt.Sprintf("https://www.facebook.com/sharer/sharer.php?u=%s", ref)
}

func buildTelegramLink(ref string, text string) string {
	result := fmt.Sprintf("https://t.me/share/url?url=%s", ref)
	return addTextPreview(result, text)
}

func buildTwitterLink(ref string, text string) string {
	result := fmt.Sprintf("https://twitter.com/share?url=%s", ref)
	return addTextPreview(result, text)
}

func buildVkLink(ref string, text string) string {
	result := fmt.Sprintf("https://vk.com/share.php?url=%s", ref)
	return addTextPreview(result, text)
}
