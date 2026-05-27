package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
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
// Skill is for marshaling with a fixed type of "Skill".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Skill", "cv:Skill", or "https://w3id.org/fep/6158#Skill".
//
// For unmarshaling that accepts any type value, use [AnySkill] instead.
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

func (receiver *Skill) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawSkill
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal skill")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			var typeValue string
			err := gojson.Unmarshal(bb, &typeValue)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal skill type")
				return err
			}

			switch typeValue {
			case TypeSkill, CompactTypeSkill, ExpandedTypeSkill:
				// OK
			default:
				return erorr.Errorf("jsonresume: unexpected type for skill: %q", typeValue)
			}
		}
	}

	return receiver.CoreSkill.unmarshalRawSkill(raw, &receiver.ID)
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
