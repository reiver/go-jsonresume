package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-nul"
)

type CoreExperience struct {
	EndDate      nul.Nullable[string] `json:"endDate"`
	Highlights   activitypub.Strings  `json:"highlights"`
	Organization nul.Nullable[string] `json:"organization"`
	Position     activitypub.Strings  `json:"position"`
	StartDate    nul.Nullable[string] `json:"startDate"`
}
