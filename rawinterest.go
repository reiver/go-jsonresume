package jsonresume

import (
	gojson "encoding/json"
)

// rawInterest is used as an intermediate step when unmarshaling [AnyInterest] or [Interest].
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.Unmarshal] or similar.
type rawInterest struct {
	ID       gojson.RawMessage `json:"id"`
	Type     gojson.RawMessage `json:"type"`
	Name     gojson.RawMessage `json:"name"`
	Keywords gojson.RawMessage `json:"keywords"`
}
