package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
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
//			"name": "Woodworking",
//			"keywords": ["Furniture"],
//		}
//	],
//
// Skill is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Skills = append(cv.Skills, jsonresume.Skill{
//		Name:     nul.Something("Photography"),
//		Keywords: activitypub.SomeStrings("Astrophotography", "Food", "Nature"),
//	})
//	
//	cv.Skills = append(cv.Skills, jsonresume.Skill{
//		Name:     nul.Something("Woodworking"),
//		Keywords: activitypub.SomeString("Furniture"),
//	})
//
// Note that you should use Skill for marshaling but not unmarshaling.
// For unmarshaling instead use [AnySkill].
//
// See also:
//
//	• [AnySkill]
//	• [CoreSkill]
//	• [ProtoSkill]
//	• [SkillID]
//	• [TypeSkill]
type Skill struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Skill"`

	CoreSkill
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
	}
}

func (receiver Skill) ProtoObject() activitypub.AnyObject {
	const _type string = TypeSkill

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Skill) ProtoSkill() AnySkill {
	const _type string = TypeSkill

	return AnySkill{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreSkill: receiver.CoreSkill,
	}
}
