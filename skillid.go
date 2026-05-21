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
//	var cv jsonresume.JSONResume
//	
//	// ...
//	
//	cv.Skills = append(cv.Skills, jsonresume.SkillID("http://example.com/resume/skill/backend-development"))
//	
//	cv.Skills = append(cv.Skills, jsonresume.SkillID("http://example.com/resume/skill/web-development"))
type SkillID string

func (receiver SkillID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeSkill

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver SkillID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeSkill

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver SkillID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeSkill

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}
