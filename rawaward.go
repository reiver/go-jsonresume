package jsonresume

import (
	gojson "encoding/json"
)

// rawAward is used as an intermediate step when unmarshaling [AnyAward] or [Award].
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.Unmarshal] or similar.
type rawAward struct {
	ID      gojson.RawMessage `json:"id"`
	Type    gojson.RawMessage `json:"type"`
	Awarder gojson.RawMessage `json:"awarder"`
	Date    gojson.RawMessage `json:"date"`
	Summary gojson.RawMessage `json:"summary"`
	Title   gojson.RawMessage `json:"title"`
}
