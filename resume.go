package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Resume is for marshaling a JSON Resume with a fixed type of "Resume".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Resume", "cv:Resume", or "https://w3id.org/fep/6158#Resume".
//
// For unmarshaling that accepts any type value, use [AnyResume] instead.
//
// See also:
//
//	• [AnyResume]
//	• [CoreResume]
//	• [ProtoResume]
//	• [ResumeID]
//	• [TypeResume]
type Resume struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Resume"`

	CoreResume
}

func (receiver Resume) ProtoNode() activitypub.AnyNode {
	const _type string = TypeResume

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Resume) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeResume

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Resume) ProtoObject() activitypub.AnyObject {
	const _type string = TypeResume

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Resume) ProtoResume() AnyResume {
	const _type string = TypeResume

	return AnyResume{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreResume: receiver.CoreResume,
	}
}

func (receiver *Resume) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawResume
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal resume")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			var typeValue string
			err := gojson.Unmarshal(bb, &typeValue)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume type")
				return err
			}

			switch typeValue {
			case TypeResume, CompactTypeResume, ExpandedTypeResume:
				// OK
			default:
				return erorr.Errorf("jsonresume: unexpected type for resume: %q", typeValue)
			}
		}
	}

	return receiver.CoreResume.unmarshalRawResume(raw, &receiver.ID)
}
