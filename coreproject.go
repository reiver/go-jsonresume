package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreProject struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Description nul.Nullable[string]    `json:"description,omitempty"`
	EndDate     nul.Nullable[string]    `json:"endDate,omitempty"`
	Entity      nul.Nullable[string]    `json:"entity,omitempty"`
	Highlights  activitypub.Strings     `json:"highlights,omitempty"`
	Keywords    activitypub.Strings     `json:"keywords,omitempty"`
	Name        nul.Nullable[string]    `json:"name,omitempty"               jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Roles       activitypub.Strings     `json:"roles,omitempty"`
	StartDate   nul.Nullable[string]    `json:"startDate,omitempty"`
	URL         []activitypub.ProtoLink `json:"url,omitempty,jsonld.compact" jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}
