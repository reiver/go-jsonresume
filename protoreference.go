package jsonresume

import (
	gobytes "bytes"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-json"
)

// ProtoReference represents something that is any type 'Reference', including sub-types.
//
// See also:
//
//	• [AnyReference]
//	• [Reference]
//	• [ReferenceID]
type ProtoReference interface {
	activitypub.ProtoObject
	ProtoReference() AnyReference
}

var (
	_ ProtoReference = Reference{}
	_ ProtoReference = ReferenceID{}
)

func ProtoReferenceUnmarshalJSON(bytes []byte) (ProtoReference, error) {
	if len(bytes) <= 0 {
		return nil, ErrBytesEmpty
	}

	if gobytes.Equal(null, bytes) {
		return nil, nil
	}

	switch bytes[0] {
	case '"':
		var dst ReferenceID
		err := json.Unmarshal(bytes, &dst)
		if nil != err {
			err = erorr.Wrap(err, "failed to json-unmarshal json string",
				field.FormattedString("destination-type", "%T", dst),
			)
			return nil, err
		}

		return dst, nil
	case '{':
		var dst Reference
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
