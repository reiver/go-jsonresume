package jsonresume

import (
	gojson "encoding/json"
)

// rawExperience is used as an intermediate step when unmarshaling [AnyExperience] or [Experience].
//
// The fields that require this intermediate step are those typed as non-empty interfaces
// (such as []activitypub.ProtoLink) which cannot be directly unmarshaled by the JSON unmarshaler.
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray] or similar.
type rawExperience struct {
	ID           gojson.RawMessage `json:"id"`
	Type         gojson.RawMessage `json:"type"`
	Description  gojson.RawMessage `json:"description"`
	EndDate      gojson.RawMessage `json:"endDate"`
	Highlights   gojson.RawMessage `json:"highlights"`
	Location     gojson.RawMessage `json:"location"`
	Name         gojson.RawMessage `json:"name"`
	Organization gojson.RawMessage `json:"organization"`
	Position     gojson.RawMessage `json:"position"`
	StartDate    gojson.RawMessage `json:"startDate"`
	Summary      gojson.RawMessage `json:"summary"`
	URL          gojson.RawMessage `json:"url"`
}
