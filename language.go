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
//			"language": "English",
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
//
// Language is for marshaling with a fixed type of "Language".
// For unmarshaling use [AnyLanguage] instead.
//
// See also:
//
//	• [AnyLanguage]
//	• [CoreLanguage]
//	• [LanguageID]
//	• [ProtoLanguage]
//	• [TypeLanguage]
type Language struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Language"`

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
	}
}

func (receiver Language) ProtoObject() activitypub.AnyObject {
	const _type string = TypeLanguage

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Language) ProtoLanguage() AnyLanguage {
	const _type string = TypeLanguage

	return AnyLanguage{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreLanguage: receiver.CoreLanguage,
	}
}
