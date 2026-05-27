package jsonresume

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreMeta struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Canonical    nul.Nullable[string] `json:"canonical,omitempty"`
	LastModified nul.Nullable[string] `json:"lastModified,omitempty"`
	Version      nul.Nullable[string] `json:"version,omitempty"`
}
