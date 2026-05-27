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
//
// Interest is for marshaling with a fixed type of "Interest".
// For unmarshaling use [AnyInterest] instead.
//
// See also:
//
//	• [AnyInterest]
//	• [CoreInterest]
//	• [InterestID]
//	• [ProtoInterest]
//	• [TypeInterest]
type Interest struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Interest"`

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

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver Interest) ProtoObject() activitypub.AnyObject {
	const _type string = TypeInterest

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver Interest) ProtoInterest() AnyInterest {
	const _type string = TypeInterest

	return AnyInterest{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreInterest: receiver.CoreInterest,
	}
}
