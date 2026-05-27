package jsonresume

import (
	gojson "encoding/json"
)

// rawBasics is used as an intermediate step when unmarshaling [AnyBasics] or [Basics].
//
// The fields that require this intermediate step are those typed as non-empty interfaces
// (such as []activitypub.ProtoLink, []activitypub.ProtoImageOrProtoLink, []ProtoProfile, []ProtoLocation)
// which cannot be directly unmarshaled by the JSON unmarshaler.
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray] or similar.
type rawBasics struct {
	ID       gojson.RawMessage `json:"id"`
	Type     gojson.RawMessage `json:"type"`
	EMail    gojson.RawMessage `json:"email"`
	Image    gojson.RawMessage `json:"image"`
	Label    gojson.RawMessage `json:"label"`
	Location gojson.RawMessage `json:"location"`
	Name     gojson.RawMessage `json:"name"`
	Phone    gojson.RawMessage `json:"phone"`
	Profiles gojson.RawMessage `json:"profiles"`
	Summary  gojson.RawMessage `json:"summary"`
	URL      gojson.RawMessage `json:"url"`
}
