package jsonresume

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreLanguage struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Fluency  nul.Nullable[string] `json:"fluency"`
	Language nul.Nullable[string] `json:"language"`
}
