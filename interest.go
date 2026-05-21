package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

// Skill implements an embedded item in the JSON Resume "skills" fields array.
//
// In a JSON Resume document this might look like:
//
//	"skills": [
//		{
//			"name": "Photography",
//			"keywords": ["Astrophotography", "Food", "Nature"],
//		},
//		{
//			"name": "Workworking",
//			"keywords": ["Furniture"],
//		}
//	],
//
// Skill is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.JSONResume
//	
//	// ...
//	
//	cv.Skills = append(cv.Skills, jsonresume.Skill{
//		Name:     nul.Nullable("Photography"),
//		Keywords: activitypub.SomeStrings("Astrophotography", "Food", "Nature"),
//	})
//	
//	cv.Skills = append(cv.Skills, jsonresume.Skill{
//		Name:     nul.Nullable("Workworking"),
//		Keywords: activitypub.SomeString("Furniture"),
//	})
type Skill struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Skill"`

	activitypub.CoreEntity
	activitypub.CoreObject

	Level    nul.Nullable[string] `json:"level"`
	Keywords activitypub.Strings  `json:"keywords"`
}

func (receiver Skill) ProtoNode() activitypub.AnyNode {
	const _type string = TypeSkill

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Skill) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeSkill

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Skill) ProtoObject() activitypub.AnyObject {
	const _type string = TypeSkill

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
