package jsonresume

import (
	gojson "encoding/json"
)

// rawMeta is used as an intermediate step when unmarshaling [AnyMeta] or [Meta].
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.Unmarshal] or similar.
type rawMeta struct {
	ID           gojson.RawMessage `json:"id"`
	Type         gojson.RawMessage `json:"type"`
	Canonical    gojson.RawMessage `json:"canonical"`
	LastModified gojson.RawMessage `json:"lastModified"`
	Version      gojson.RawMessage `json:"version"`
}
