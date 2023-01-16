package templateFunctions

import platformContracts "github.com/punk-link/platform-contracts"

func GetPlatformIconPath(platformId string) string {
	switch platformId {
	case platformContracts.Spotify:
		return "Spotyfy.png"
	case platformContracts.Deezer:
		return ""
	}

	return ""
}

func GetPlatformName(platformId string) string {
	switch platformId {
	case platformContracts.Spotify:
		return "Spotyfy"
	case platformContracts.Deezer:
		return "Deezer"
	}

	return ""
}
