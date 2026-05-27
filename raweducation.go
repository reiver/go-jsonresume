package jsonresume

import (
	gojson "encoding/json"
)

// rawEducation is used as an intermediate step when unmarshaling [AnyEducation] or [Education].
//
// The fields that require this intermediate step are those typed as non-empty interfaces
// (such as []activitypub.ProtoLink) which cannot be directly unmarshaled by the JSON unmarshaler.
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray] or similar.
type rawEducation struct {
	ID          gojson.RawMessage `json:"id"`
	Type        gojson.RawMessage `json:"type"`
	Area        gojson.RawMessage `json:"area"`
	Courses     gojson.RawMessage `json:"courses"`
	EndDate     gojson.RawMessage `json:"endDate"`
	Institution gojson.RawMessage `json:"institution"`
	Score       gojson.RawMessage `json:"score"`
	StartDate   gojson.RawMessage `json:"startDate"`
	StudyType   gojson.RawMessage `json:"studyType"`
	URL         gojson.RawMessage `json:"url"`
}
