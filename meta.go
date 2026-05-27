package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Meta implements an embedded item in the JSON Resume "meta" field.
//
// In a JSON Resume document this might look like:
//
//	"meta": {
//		"canonical": "https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json",
//		"version": "v1.0.0",
//		"lastModified": "2017-12-24T15:53:00"
//	}
//
// Meta is an implementation of the JSON Resume meta object.
//
// Example usage:
//
//	var cv jsonresume.Resume
//
//	// ...
//
//	cv.Meta = jsonresume.Meta{
//		Canonical:    nul.Something("https://example.com/resume.json"),
//		Version:      nul.Something("v1.0.0"),
//		LastModified: nul.Something("2017-12-24T15:53:00"),
//	}
//
// Meta is for marshaling with a fixed type of "Meta".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Meta", "cv:Meta", or "https://w3id.org/fep/6158#Meta".
//
// For unmarshaling that accepts any type value, use [AnyMeta] instead.
//
// See also:
//
//	• [AnyMeta]
//	• [CoreMeta]
//	• [MetaID]
//	• [ProtoMeta]
//	• [TypeMeta]
type Meta struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Meta"`

	CoreMeta
}

func (receiver *Meta) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawMeta
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal meta")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			var typeValue string
			err := gojson.Unmarshal(bb, &typeValue)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal meta type")
				return err
			}

			switch typeValue {
			case TypeMeta, CompactTypeMeta, ExpandedTypeMeta:
				// OK
			default:
				return erorr.Errorf("jsonresume: unexpected type for meta: %q", typeValue)
			}
		}
	}

	return receiver.CoreMeta.unmarshalRawMeta(raw, &receiver.ID)
}

func (receiver Meta) ProtoNode() activitypub.AnyNode {
	const _type string = TypeMeta

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Meta) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeMeta

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Meta) ProtoObject() activitypub.AnyObject {
	const _type string = TypeMeta

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Meta) ProtoMeta() AnyMeta {
	const _type string = TypeMeta

	return AnyMeta{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreMeta: receiver.CoreMeta,
	}
}
