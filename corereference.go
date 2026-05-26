package jsonresume

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreReference struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Name      nul.Nullable[string] `json:"name,omitempty"      jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Reference nul.Nullable[string] `json:"reference,omitempty"`
}
