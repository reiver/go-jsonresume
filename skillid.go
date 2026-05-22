package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// SkillID implements a referenced item in the JSON Resume "skills" fields array.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"skills": [
//		"http://example.com/resume/skill/backend-development",
//		"http://example.com/resume/skill/web-development"
//	],
//
// SkillID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Skills = append(cv.Skills, jsonresume.SkillID("http://example.com/resume/skill/backend-development"))
//	
//	cv.Skills = append(cv.Skills, jsonresume.SkillID("http://example.com/resume/skill/web-development"))
type SkillID jsonld.ID

func SomeSkillID(value string) SkillID {
	return SkillID(jsonld.SomeID(value))
}

func (receiver SkillID) MarshalJSON() ([]byte, error) {
	return jsonld.ID(receiver).MarshalJSON()
}

func (receiver SkillID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeSkill

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver SkillID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeSkill

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver SkillID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeSkill

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver SkillID) ProtoSkill() AnySkill {
	const _type string = TypeSkill

	return AnySkill{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

// String makes [SkillID] fit the [fmt.Stringer] interface.
func (receiver SkillID) String() string {
	return jsonld.ID(receiver).GetElse("")
}
