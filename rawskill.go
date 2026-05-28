package jsonresume

import (
	gojson "encoding/json"
)

// rawSkill is used as an intermediate step when unmarshaling [AnySkill] or [Skill].
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.Unmarshal] or similar.
type rawSkill struct {
	ID       gojson.RawMessage `json:"id"`
	Type     gojson.RawMessage `json:"type"`
	Keywords gojson.RawMessage `json:"keywords"`
	Level    gojson.RawMessage `json:"level"`
	Name     gojson.RawMessage `json:"name"`
}
