package jsonresume

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreReference struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Reference nul.Nullable[string] `json:"reference"`
}
