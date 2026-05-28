package jsonresume

import (
	gojson "encoding/json"
)

// rawLanguage is used as an intermediate step when unmarshaling [AnyLanguage] or [Language].
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.Unmarshal] or similar.
type rawLanguage struct {
	ID       gojson.RawMessage `json:"id"`
	Type     gojson.RawMessage `json:"type"`
	Fluency  gojson.RawMessage `json:"fluency"`
	Language gojson.RawMessage `json:"language"`
}
