package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// EducationID implements a referenced item in the JSON Resume "education" fields array.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"education": [
//		"http://example.com/resume/education/sfu",
//		"http://example.com/resume/education/kpu"
//	],
//
// EducationID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Educations = append(cv.Educations, jsonresume.EducationID("http://example.com/resume/education/sfu"))
//	cv.Educations = append(cv.Educations, jsonresume.EducationID("http://example.com/resume/education/kpu"))
type EducationID string

func (receiver EducationID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeEducation

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver EducationID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeEducation

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver EducationID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeEducation

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver EducationID) ProtoEducation() AnyEducation {
	const _type string = TypeEducation

	return AnyEducation{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}
