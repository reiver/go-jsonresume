package jsonresume

import (
	gojson "encoding/json"
)

// rawProject is used as an intermediate step when unmarshaling [AnyProject] or [Project].
//
// The fields that require this intermediate step are those typed as non-empty interfaces
// (such as []activitypub.ProtoLink) which cannot be directly unmarshaled by the JSON unmarshaler.
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray] or similar.
type rawProject struct {
	ID          gojson.RawMessage `json:"id"`
	AtType      gojson.RawMessage `json:"@type"`
	Description gojson.RawMessage `json:"description"`
	EndDate     gojson.RawMessage `json:"endDate"`
	Entity      gojson.RawMessage `json:"entity"`
	Highlights  gojson.RawMessage `json:"highlights"`
	Keywords    gojson.RawMessage `json:"keywords"`
	Name        gojson.RawMessage `json:"name"`
	Roles       gojson.RawMessage `json:"roles"`
	StartDate   gojson.RawMessage `json:"startDate"`
	Type        gojson.RawMessage `json:"type"`
	URL         gojson.RawMessage `json:"url"`
}
