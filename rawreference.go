package jsonresume

import (
	gojson "encoding/json"
)

// rawReference is used as an intermediate step when unmarshaling [AnyReference] or [Reference].
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.Unmarshal] or similar.
type rawReference struct {
	ID        gojson.RawMessage `json:"id"`
	Type      gojson.RawMessage `json:"type"`
	Name      gojson.RawMessage `json:"name"`
	Reference gojson.RawMessage `json:"reference"`
}
