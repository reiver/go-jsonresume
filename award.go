package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Award implements an embedded item in the JSON Resume "awards" fields array.
//
// In a JSON Resume document this might look like:
//
//	"awards": [
//		{
//			"title": "Best Employee (2024)",
//			"date": "2024-05-21",
//			"awarder": "SuperCo",
//			"summary": "He did good work."
//		},
//		{
//			"title": "Acme Excellence (2021)",
//			"date": "2021-01-17",
//			"awarder": "Acme",
//			"summary": "For Joe Blow's excellent work."
//		}
//	],
//
// Award is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Awards = append(cv.Awards, jsonresume.Award{
//		Name:     nul.Something("Photography"),
//		Keywords: activitypub.SomeStrings("Astrophotography", "Food", "Nature"),
//	})
//	
//	cv.Awards = append(cv.Awards, jsonresume.Award{
//		Name:     nul.Something("Workworking"),
//		Keywords: activitypub.SomeString("Furniture"),
//	})
type Award struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Award"`

	activitypub.CoreEntity
	activitypub.CoreObject
	CoreAward
}

func (receiver Award) ProtoNode() activitypub.AnyNode {
	const _type string = TypeAward

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Award) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeAward

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Award) ProtoObject() activitypub.AnyObject {
	const _type string = TypeAward

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

func (receiver Award) ProtoAward() AnyAward {
	const _type string = TypeAward

	var result = AnyAward{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
		CoreObject: receiver.CoreObject,
		CoreAward:  receiver.CoreAward,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}
