package jsonresume

import (
	gobytes "bytes"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-json"
)

// ProtoSkill represents something that is any type 'Skill', including sub-types.
//
// See also:
//
//	• [AnySkill]
//	• [Skill]
//	• [SkillID]
type ProtoSkill interface {
	activitypub.ProtoObject
	ProtoSkill() AnySkill
}

var (
	_ ProtoSkill = Skill{}
	_ ProtoSkill = SkillID{}
)

func ProtoSkillUnmarshalJSON(bytes []byte) (ProtoSkill, error) {
	if len(bytes) <= 0 {
		return nil, ErrBytesEmpty
	}

	if gobytes.Equal(null, bytes) {
		return nil, nil
	}

	switch bytes[0] {
	case '"':
		var dst SkillID
		err := json.Unmarshal(bytes, &dst)
		if nil != err {
			err = erorr.Wrap(err, "failed to json-unmarshal json string",
				field.FormattedString("destination-type", "%T", dst),
			)
			return nil, err
		}

		return dst, nil
	case '{':
		var dst Skill
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
