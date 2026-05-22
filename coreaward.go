package jsonresume

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreAward struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Awarder nul.Nullable[string] `json:"awarder"`
	Date    nul.Nullable[string] `json:"date"`
	Title   nul.Nullable[string] `json:"title"`
}
