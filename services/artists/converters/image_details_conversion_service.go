package converters

import (
	contracts "github.com/punk-link/presentation-contracts"
)

func toImageDetailsMap(details *contracts.ImageDetails) map[string]any {
	return map[string]any{
		"AltText": details.AltText,
		"Height":  details.Height,
		"Url":     details.Url,
		"Width":   details.Width,
	}
}
