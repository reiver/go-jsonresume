package jsonresume

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreCertificate struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Date   nul.Nullable[string] `json:"date"`
	Issuer nul.Nullable[string] `json:"issuer"`
}
