package jsonresume

import (
	"github.com/reiver/go-nul"
)

type CoreLocation struct {
	Address     nul.Nullable[string] `json:"address"`
	City        nul.Nullable[string] `json:"city"`
	CountryCode nul.Nullable[string] `json:"countryCode"`
	PostcalCode nul.Nullable[string] `json:"postalCode"`
	Region      nul.Nullable[string] `json:"region"`
}
