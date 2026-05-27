package jsonresume

import (
	gojson "encoding/json"
)

// rawResume is used as an intermediate step in [AnyResume.UnmarshalJSON] and [Resume.UnmarshalJSON].
//
// The JSON for [AnyResume] and [Resume] has a number of JSON resume fields:
//
//	• "awards"
//	• "basics"
//	• "certificates"
//	• "education"
//	• "interests"
//	• "languages"
//	• "meta"
//	• "projects"
//	• "publications"
//	• "references"
//	• "skills"
//	• "volunteer"
//	• "work"
//
// As well as some ActivityPub related fields:
//
//	• "id"
//	• "type"
//
// The original raw JSON []byte for a resume is first loaded into rawResume.
//
// This has the effect of splitting that resume raw JSON into the fields:
//
//	• "awards"       → rawResume.Awards
//	• "basics"       → rawResume.Basics
//	• "certificates" → rawResume.Certificates
//	• "education"    → rawResume.Education
//	• "interests"    → rawResume.Interests
//	• "languages"    → rawResume.Languages
//	• "meta"         → rawResume.Meta
//	• "projects"     → rawResume.Projects
//	• "publications" → rawResume.Publications
//	• "references"   → rawResume.References
//	• "skills"       → rawResume.Skills
//	• "volunteer"    → rawResume.Volunteer
//	• "work"         → rawResume.Work
//
// And:
//
//	• "id"           → rawResume.ID
//	• "type"         → rawResume.Type
//
// Each of these contains the raw JSON just for that field.
//
// This is then used in the next step that [jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray] provides.
type rawResume struct {
	ID           gojson.RawMessage `json:"id"`
	Type         gojson.RawMessage `json:"type"`
	Awards       gojson.RawMessage `json:"awards"`
	Basics       gojson.RawMessage `json:"basics"`
	Certificates gojson.RawMessage `json:"certificates"`
	Education    gojson.RawMessage `json:"education"`
	Interests    gojson.RawMessage `json:"interests"`
	Languages    gojson.RawMessage `json:"languages"`
	Meta         gojson.RawMessage `json:"meta"`
	Projects     gojson.RawMessage `json:"projects"`
	Publications gojson.RawMessage `json:"publications"`
	References   gojson.RawMessage `json:"references"`
	Skills       gojson.RawMessage `json:"skills"`
	Volunteer    gojson.RawMessage `json:"volunteer"`
	Work         gojson.RawMessage `json:"work"`
}
