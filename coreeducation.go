package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreEducation struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Area        activitypub.Strings     `json:"area,omitempty"`
	Courses     activitypub.Strings     `json:"courses,omitempty"`
	EndDate     nul.Nullable[string]    `json:"endDate,omitempty"`
	Institution nul.Nullable[string]    `json:"institution,omitempty"`
	Score       nul.Nullable[string]    `json:"score,omitempty"`
	StartDate   nul.Nullable[string]    `json:"startDate,omitempty"`
	StudyType   activitypub.Strings     `json:"studyType,omitempty"`
	URL         []activitypub.ProtoLink `json:"url,omitempty,jsonld.compact" jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}
