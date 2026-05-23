package jsonresume

import (
	gobytes "bytes"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-json"
)

// ProtoCertificate represents something that is any type 'Certificate', including sub-types.
//
// See also:
//
//	• [AnyCertificate]
//	• [Certificate]
//	• [CertificateID]
type ProtoCertificate interface {
	activitypub.ProtoObject
	ProtoCertificate() AnyCertificate
}

var (
	_ ProtoCertificate = Certificate{}
	_ ProtoCertificate = CertificateID{}
)

func ProtoCertificateUnmarshalJSON(bytes []byte) (ProtoCertificate, error) {
	if len(bytes) <= 0 {
		return nil, ErrBytesEmpty
	}

	if gobytes.Equal(null, bytes) {
		return nil, nil
	}

	switch bytes[0] {
	case '"':
		var dst CertificateID
		err := json.Unmarshal(bytes, &dst)
		if nil != err {
			err = erorr.Wrap(err, "failed to json-unmarshal json string",
				field.FormattedString("destination-type", "%T", dst),
			)
			return nil, err
		}

		return dst, nil
	case '{':
		var dst Certificate
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
