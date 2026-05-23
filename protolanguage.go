package jsonresume

import (
	gobytes "bytes"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-json"
)

// ProtoLanguage represents something that is any type 'Language', including sub-types.
//
// See also:
//
//	• [AnyLanguage]
//	• [Language]
//	• [LanguageID]
type ProtoLanguage interface {
	activitypub.ProtoObject
	ProtoLanguage() AnyLanguage
}

var (
	_ ProtoLanguage = Language{}
	_ ProtoLanguage = LanguageID{}
)

func ProtoLanguageUnmarshalJSON(bytes []byte) (ProtoLanguage, error) {
	if len(bytes) <= 0 {
		return nil, ErrBytesEmpty
	}

	if gobytes.Equal(null, bytes) {
		return nil, nil
	}

	switch bytes[0] {
	case '"':
		var dst LanguageID
		err := json.Unmarshal(bytes, &dst)
		if nil != err {
			err = erorr.Wrap(err, "failed to json-unmarshal json string",
				field.FormattedString("destination-type", "%T", dst),
			)
			return nil, err
		}

		return dst, nil
	case '{':
		var dst Language
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
