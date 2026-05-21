package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Language implements an embedded item in the JSON Resume "languages" fields array.
//
// In a JSON Resume document this might look like:
//
//	"languages": [
//		{
//			"language": "Enlish",
//			"fluency": "Fluent"
//		},
//		{
//			"language": "Persian",
//			"fluency": "Beginner"
//		}
//	],
//
// Language is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Languages = append(cv.Languages, jsonresume.Language{
//		Language: nul.Something("English"),
//		Fluency:  nul.Something("Fluent"),
//	})
//	
//	cv.Languages = append(cv.Languages, jsonresume.Language{
//		Language: nul.Something("Persian"),
//		Fluency:  nul.Something("Beginner"),
//	})
type Language struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Language"`

	activitypub.CoreEntity
	activitypub.CoreObject
	CoreLanguage
}

func (receiver Language) ProtoNode() activitypub.AnyNode {
	const _type string = TypeLanguage

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Language) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeLanguage

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Language) ProtoObject() activitypub.AnyObject {
	const _type string = TypeLanguage

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

func (receiver Language) ProtoLanguage() AnyLanguage {
	const _type string = TypeLanguage

	var result = AnyLanguage{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
		CoreObject: receiver.CoreObject,
		CoreLanguage:  receiver.CoreLanguage,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}
