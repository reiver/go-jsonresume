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
	Highlights  activitypub.Strings     `json:"highlights"`
	Keywords    activitypub.Strings     `json:"keywords"`
	Name        nul.Nullable[string]    `json:"name,omitempty"        jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Roles       activitypub.Strings     `json:"roles"`
	StartDate   nul.Nullable[string]    `json:"startDate,omitempty"`
	URL         []activitypub.ProtoLink `json:"url"         jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}
