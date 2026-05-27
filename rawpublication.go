package jsonresume

import (
	gojson "encoding/json"
)

// rawPublication is used as an intermediate step when unmarshaling [AnyPublication] or [Publication].
//
// The fields that require this intermediate step are those typed as non-empty interfaces
// (such as []activitypub.ProtoLink) which cannot be directly unmarshaled by the JSON unmarshaler.
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray] or similar.
type rawPublication struct {
	ID          gojson.RawMessage `json:"id"`
	Type        gojson.RawMessage `json:"type"`
	Name        gojson.RawMessage `json:"name"`
	Publisher   gojson.RawMessage `json:"publisher"`
	ReleaseDate gojson.RawMessage `json:"releaseDate"`
	Summary     gojson.RawMessage `json:"summary"`
	URL         gojson.RawMessage `json:"url"`
}
