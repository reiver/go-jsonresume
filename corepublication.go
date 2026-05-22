package jsonresume

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CorePublication struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Publisher   nul.Nullable[string] `json:"publisher"`
	ReleaseDate nul.Nullable[string] `json:"releaseDate"`
}
