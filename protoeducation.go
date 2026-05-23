package jsonresume

import (
	gobytes "bytes"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-json"
)

// ProtoEducation represents something that is any type 'Education', including sub-types.
//
// See also:
//
//	• [AnyEducation]
//	• [Education]
//	• [EducationID]
type ProtoEducation interface {
	activitypub.ProtoObject
	ProtoEducation() AnyEducation
}

var (
	_ ProtoEducation = Education{}
	_ ProtoEducation = EducationID{}
)

func ProtoEducationUnmarshalJSON(bytes []byte) (ProtoEducation, error) {
	if len(bytes) <= 0 {
		return nil, ErrBytesEmpty
	}

	if gobytes.Equal(null, bytes) {
		return nil, nil
	}

	switch bytes[0] {
	case '"':
		var dst EducationID
		err := json.Unmarshal(bytes, &dst)
		if nil != err {
			err = erorr.Wrap(err, "failed to json-unmarshal json string",
				field.FormattedString("destination-type", "%T", dst),
			)
			return nil, err
		}

		return dst, nil
	case '{':
		var dst Education
		err := json.Unmarshal(bytes, &dst)
		if nil != err {
			err = erorr.Wrap(err, "failed to json-unmarshal json object",
				field.FormattedString("destination-type", "%T", dst),
			)
			return nil, err
		}

		return dst, nil
	default:
		var err error = ErrTypeUnsupported
		err = erorr.Wrap(err, "failed to json-unmarshal")
		return nil, err
	}
}
