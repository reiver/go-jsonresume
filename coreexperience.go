package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreExperience struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	EndDate      nul.Nullable[string] `json:"endDate"`
	Highlights   activitypub.Strings  `json:"highlights"`
	Organization nul.Nullable[string] `json:"organization"`
	Position     activitypub.Strings  `json:"position"`
	StartDate    nul.Nullable[string] `json:"startDate"`
}
