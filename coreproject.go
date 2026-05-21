package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-nul"
)

type CoreProject struct {
	Description nul.Nullable[string] `json:"description"`
	EndDate     nul.Nullable[string] `json:"endDate"`
	Highlights  activitypub.Strings  `json:"highlights"`
	StartDate   nul.Nullable[string] `json:"startDate"`
}
