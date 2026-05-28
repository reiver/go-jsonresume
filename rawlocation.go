package jsonresume

import (
	gojson "encoding/json"
)

// rawLocation is used as an intermediate step when unmarshaling [AnyLocation] or [Location].
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.Unmarshal] or similar.
type rawLocation struct {
	ID          gojson.RawMessage `json:"id"`
	Type        gojson.RawMessage `json:"type"`
	Address     gojson.RawMessage `json:"address"`
	City        gojson.RawMessage `json:"city"`
	CountryCode gojson.RawMessage `json:"countryCode"`
	PostalCode  gojson.RawMessage `json:"postalCode"`
	Region      gojson.RawMessage `json:"region"`
}
