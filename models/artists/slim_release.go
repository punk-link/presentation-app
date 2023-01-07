package artists

import (
	"time"

	contracts "github.com/punk-link/presentation-contracts"
)

type SlimRelease struct {
	Slug         string
	ImageDetails *contracts.ImageDetails
	Name         string
	ReleaseDate  time.Time
	Type         string
}
