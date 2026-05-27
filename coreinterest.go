package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreInterest struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Name     nul.Nullable[string] `json:"name,omitempty"     jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Keywords activitypub.Strings  `json:"keywords,omitempty"`
}
