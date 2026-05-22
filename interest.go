package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Interest implements an embedded item in the JSON Resume "interests" fields array.
//
// In a JSON Resume document this might look like:
//
//	"interests": [
//		{
//			"name": "Back-End Development",
//			"keywords": ["Golang", "PHP", "HTTP"],
//		},
//		{
//			"name": "Web Development",
//			"keywords": ["CSS", "HTML", "JavaScript"],
//		}
//	],
//
// Interest is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Interests = append(cv.Interests, jsonresume.Interest{
//		Name:     nul.Something("Back-End Development"),
//		Keywords: activitypub.SomeStrings("Golang", "PHP", "HTTP"),
//	})
//	
//	cv.Interests = append(cv.Interests, jsonresume.Interest{
//		Name:     nul.Something("Web Development"),
//		Keywords: activitypub.SomeStrings("CSS", "HTML", "JavaScript"),
//	})
type Interest struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Interest"`

	activitypub.CoreEntity
	activitypub.CoreObject
	CoreInterest
}

func (receiver Interest) ProtoNode() activitypub.AnyNode {
	const _type string = TypeInterest

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Interest) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeInterest

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Interest) ProtoObject() activitypub.AnyObject {
	const _type string = TypeInterest

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

func (receiver Interest) ProtoInterest() AnyInterest {
	const _type string = TypeInterest

	return AnyInterest{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreInterest: receiver.CoreInterest,
	}
}
