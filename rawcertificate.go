package jsonresume

import (
	gojson "encoding/json"
)

// rawCertificate is used as an intermediate step when unmarshaling [AnyCertificate] or [Certificate].
//
// The fields that require this intermediate step are those typed as non-empty interfaces
// (such as []activitypub.ProtoLink) which cannot be directly unmarshaled by the JSON unmarshaler.
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray] or similar.
type rawCertificate struct {
	ID     gojson.RawMessage `json:"id"`
	Type   gojson.RawMessage `json:"type"`
	Date   gojson.RawMessage `json:"date"`
	Name   gojson.RawMessage `json:"name"`
	Issuer gojson.RawMessage `json:"issuer"`
	URL    gojson.RawMessage `json:"url"`
}
