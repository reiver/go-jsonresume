package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Reference implements an embedded item in the JSON Resume "references" fields array.
//
// In a JSON Resume document this might look like:
//
//	"references": [
//		{
//			"name": "Jane Doe",
//			"reference": "Joe Blow is an exceptionally talented professional."
//		},
//		{
//			"name": "Bob Smith",
//			"reference": "I enjoyed working with Joe Blow and give them my highest recommendation without reservation."
//		},
//	],
//
// Reference is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.References = append(cv.References, jsonresume.Reference{
//		Name:      nul.Something("Jane Doe"),
//		Reference: nul.Something("Joe Blow is an exceptionally talented professional."),
//	})
//	
//	cv.References = append(cv.References, jsonresume.Reference{
//		Name:      nul.Something("Bob Smith"),
//		Reference: nul.Something("I enjoyed working with Joe Blow and give them my highest recommendation without reservation."),
//	})
//
// Reference is for marshaling with a fixed type of "Reference".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Reference", "cv:Reference", or "https://w3id.org/fep/6158#Reference".
//
// For unmarshaling that accepts any type value, use [AnyReference] instead.
//
// See also:
//
//	• [AnyReference]
//	• [CoreReference]
//	• [ProtoReference]
//	• [ReferenceID]
//	• [TypeReference]
type Reference struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Reference"`

	CoreReference
}

func (receiver *Reference) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawReference
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal reference")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			var typeValue string
			err := gojson.Unmarshal(bb, &typeValue)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal reference type")
				return err
			}

			switch typeValue {
			case TypeReference, CompactTypeReference, ExpandedTypeReference:
				// OK
			default:
				return erorr.Errorf("jsonresume: unexpected type for reference: %q", typeValue)
			}
		}
	}

	return receiver.CoreReference.unmarshalRawReference(raw, &receiver.ID)
}

func (receiver Reference) ProtoNode() activitypub.AnyNode {
	const _type string = TypeReference

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Reference) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeReference

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver Reference) ProtoObject() activitypub.AnyObject {
	const _type string = TypeReference

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver Reference) ProtoReference() AnyReference {
	const _type string = TypeReference

	return AnyReference{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreReference: receiver.CoreReference,
	}
}
