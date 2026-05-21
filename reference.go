package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
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
type Reference struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Reference"`

	activitypub.CoreEntity
	activitypub.CoreObject
	CoreReference
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

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Reference) ProtoObject() activitypub.AnyObject {
	const _type string = TypeReference

	var result = activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
		CoreObject: receiver.CoreObject,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}

func (receiver Reference) ProtoReference() AnyReference {
	const _type string = TypeReference

	var result = AnyReference{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity:    receiver.CoreEntity,
		CoreObject:    receiver.CoreObject,
		CoreReference: receiver.CoreReference,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}
