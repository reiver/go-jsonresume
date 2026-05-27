package jsonresume

import (
	gojson "encoding/json"
)

// rawProfile is used as an intermediate step when unmarshaling [AnyProfile] or [Profile].
//
// The fields that require this intermediate step are those typed as non-empty interfaces
// (such as []activitypub.ProtoLink) which cannot be directly unmarshaled by the JSON unmarshaler.
//
// The raw JSON for each field is captured here, then dispatched in a second step
// using [jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray] or similar.
type rawProfile struct {
	ID       gojson.RawMessage `json:"id"`
	Type     gojson.RawMessage `json:"type"`
	Network  gojson.RawMessage `json:"network"`
	UserName gojson.RawMessage `json:"username"`
	URL      gojson.RawMessage `json:"url"`
}
