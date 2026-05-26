package jsonresume

import (
	gobytes "bytes"
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-jsonld"
)

// rawResume is used as an intermediate step in [Resume.UnmarshalJSON].
//
// The JSON for [Resume] has a number of fields:
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
// Each of these contains the raw JSON just for that field.
//
// This is then used in the next step that [protoSliceUnmarshalJSON] provides.
type rawResume struct {
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

// rawSliceUnmarshalJSON is used as an intermediate step in [Resume.UnmarshalJSON].
//
// The JSON for [Resume] has a number of fields whose values could be JSON arrays:
//
//	• "awards"
//	• "certificates"
//	• "education"
//	• "interests"
//	• "languages"
//	• "projects"
//	• "publications"
//	• "references"
//	• "skills"
//	• "volunteer"
//	• "work"
//
// rawSliceUnmarshalJSON takes a []byte that contains that (raw) JSON array, and into a []gojson.RawMessage.
//
// This is then used in the next step that [protoSliceUnmarshalJSON] provides.
func rawSliceUnmarshalJSON(bytes []byte, target *[]gojson.RawMessage) error {
	return jsonld.Unmarshal(bytes, target)
}

func protoSliceUnmarshalJSON[protoT activitypub.ProtoObject, Tid activitypub.ProtoObject, T activitypub.ProtoObject](bytes []byte) ([]protoT, error) {
	if len(bytes) <= 0 {
		return nil, ErrBytesEmpty
	}

	var result []protoT

	{
		var b byte = bytes[0]

		if '"' == b || '{' == b {
			t, err := protoUnmarshalJSON[protoT, Tid, T](bytes)
			if nil != err {
				return nil, err
			}

			result = append(result, t)

			return result, nil
		}
	}

        var target []gojson.RawMessage

        err := rawSliceUnmarshalJSON(bytes, &target)
        if nil != err {
                return nil, err
        }

        for _, datum := range target {

		t, err := protoUnmarshalJSON[protoT, Tid, T]([]byte(datum))
		if nil != err {
			return nil, err
		}

		result = append(result, t)
        }

	return result, nil
}

func protoUnmarshalJSON[protoT activitypub.ProtoObject, Tid activitypub.ProtoObject, T activitypub.ProtoObject](bytes []byte) (protoT, error) {
	if len(bytes) <= 0 {
		var nada protoT
		return nada, ErrBytesEmpty
	}

	if gobytes.Equal(null, bytes) {
		var nada protoT
		return nada, nil
	}

	switch bytes[0] {
	case '"':
		var dst Tid
		err := jsonld.Unmarshal(bytes, &dst)
		if nil != err {
			err = erorr.Wrap(err, "failed to json-unmarshal json string",
				field.FormattedString("destination-type", "%T", dst),
			)
			var nada protoT
			return nada, err
		}

		return any(dst).(protoT), nil
	case '{':
		var dst T
		err := jsonld.Unmarshal(bytes, &dst)
		if nil != err {
			err = erorr.Wrap(err, "failed to json-unmarshal json object",
				field.FormattedString("destination-type", "%T", dst),
			)
			var nada protoT
			return nada, err
		}

		return any(dst).(protoT), nil
	default:
		var err error = ErrTypeUnsupported
		err = erorr.Wrap(err, "failed to json-unmarshal")
		var nada protoT
		return nada, err
	}
}
