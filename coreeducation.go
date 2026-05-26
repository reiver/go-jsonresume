package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreEducation struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Area        activitypub.Strings     `json:"area"`
	Courses     activitypub.Strings     `json:"courses"`
	EndDate     nul.Nullable[string]    `json:"endDate,omitempty"`
	Institution nul.Nullable[string]    `json:"institution,omitempty"`
	Score       nul.Nullable[string]    `json:"score,omitempty"`
	StartDate   nul.Nullable[string]    `json:"startDate,omitempty"`
	StudyType   activitypub.Strings     `json:"studyType"`
	URL         []activitypub.ProtoLink `json:"url"         jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}
