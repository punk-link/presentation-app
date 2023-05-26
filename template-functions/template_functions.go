package templateFunctions

import (
	platformContracts "github.com/punk-link/platform-contracts"
	presentationContractConstants "github.com/punk-link/presentation-contracts/constants"
)

func GetPlatformIconPath(platformId string) string {
	switch platformId {
	case platformContracts.AppleMusic:
		return "AppleMusic.svg"
	case platformContracts.Deezer:
		return "Deezer.svg"
	case platformContracts.Spotify:
		return "Spotyfy.svg"
	case platformContracts.Uma:
		return "VKMusic.svg"
	case platformContracts.YandexMusic:
		return "YandexMusic.svg"
	case platformContracts.YouTubeMusic:
		return "YouTube.svg"
	}

	return ""
}

func GetPlatformName(platformId string) string {
	switch platformId {
	case platformContracts.AppleMusic:
		return "Apple Music"
	case platformContracts.Deezer:
		return "Deezer"
	case platformContracts.Spotify:
		return "Spotyfy"
	case platformContracts.Uma:
		return "VK Music"
	case platformContracts.YandexMusic:
		return "Yandex Music"
	case platformContracts.YouTubeMusic:
		return "YouTube Music"
	}

	return ""
}

func GetReleaseTypeId(releaseType string) string {
	if releaseType == platformContracts.Album {
		return "content-album"
	} else if releaseType == platformContracts.Compilation {
		return "content-compilation"
	}

	return "content-single"
}

func GetSocialNetworkIconPath(socialNetworkId string) string {
	switch socialNetworkId {
	case presentationContractConstants.FACEBOOK:
		return "facebook"
	case presentationContractConstants.INSTAGRAM:
		return "instagram"
	case presentationContractConstants.TELEGRAM:
		return "telegram"
	case presentationContractConstants.TWITTER:
		return "twitter"
	case presentationContractConstants.VK:
		return "vk"
	}

	return socialNetworkId
}
