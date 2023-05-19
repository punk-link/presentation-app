package templateFunctions

import (
	platformContracts "github.com/punk-link/platform-contracts"
	presentationContractConstants "github.com/punk-link/presentation-contracts/constants"
)

func GetPlatformIconPath(platformId string) string {
	switch platformId {
	case platformContracts.AppleMusic:
		return "AppleMusic.webp"
	case platformContracts.Deezer:
		return "Deezer.webp"
	case platformContracts.Spotify:
		return "Spotyfy.webp"
	case platformContracts.Uma:
		return "VKMusic.webp"
	case platformContracts.YandexMusic:
		return "YandexMusic.webp"
	case platformContracts.YouTubeMusic:
		return "YouTube.webp"
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
		return "fb"
	case presentationContractConstants.INSTAGRAM:
		return "insta"
	case presentationContractConstants.TELEGRAM:
		return "telegram"
	case presentationContractConstants.TWITTER:
		return "twitter"
	case presentationContractConstants.VK:
		return "vk"
	}

	return socialNetworkId
}
