package jsonresume

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreLocation struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Address     nul.Nullable[string] `json:"address,omitempty"`
	City        nul.Nullable[string] `json:"city,omitempty"`
	CountryCode nul.Nullable[string] `json:"countryCode,omitempty"`
	PostalCode  nul.Nullable[string] `json:"postalCode,omitempty"`
	Region      nul.Nullable[string] `json:"region,omitempty"`
}
