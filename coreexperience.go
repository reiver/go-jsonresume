package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreExperience struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Description  nul.Nullable[string]    `json:"description,omitempty"`
	EndDate      nul.Nullable[string]    `json:"endDate,omitempty"`
	Highlights   activitypub.Strings     `json:"highlights"`
	Location     nul.Nullable[string]    `json:"location,omitempty"`
	Name         nul.Nullable[string]    `json:"name,omitempty"         jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Organization nul.Nullable[string]    `json:"organization,omitempty"`
	Position     activitypub.Strings     `json:"position"`
	StartDate    nul.Nullable[string]    `json:"startDate,omitempty"`
	Summary      nul.Nullable[string]    `json:"summary,omitempty"      jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	URL          []activitypub.ProtoLink `json:"url"          jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}
