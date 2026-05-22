package jsonresume

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreLocation struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Address     nul.Nullable[string] `json:"address"`
	City        nul.Nullable[string] `json:"city"`
	CountryCode nul.Nullable[string] `json:"countryCode"`
	PostalCode  nul.Nullable[string] `json:"postalCode"`
	Region      nul.Nullable[string] `json:"region"`
}
