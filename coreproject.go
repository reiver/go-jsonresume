package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-nul"
)

type CoreProject struct {
	Description nul.Nullable[string] `json:"description"`
	EndDate     nul.Nullable[string] `json:"endDate"`
	Entity      nul.Nullable[string] `json:"entity"`
	Highlights  activitypub.Strings  `json:"highlights"`
	Keywords    activitypub.Strings  `json:"keywords"`
	Roles       activitypub.Strings  `json:"roles"`
	StartDate   nul.Nullable[string] `json:"startDate"`
}
