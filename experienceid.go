package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// ExperienceID implements a referenced item in the JSON Resume "work" and "volunteer" fields arrays.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"work": [
//		"http://example.com/resume/experience/3",
//		"http://example.com/resume/experience/2"
//	],
//	"volunteer": [
//		"http://example.com/resume/experience/4",
//		"http://example.com/resume/experience/1"
//	],
//
// ExperienceID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//
//	// ...
//
//	cv.Work = append(cv.Work, jsonresume.ExperienceID("http://example.com/resume/experience/3"))
//
//	cv.Work = append(cv.Work, jsonresume.ExperienceID("http://example.com/resume/experience/2"))
//
//	// ...
//
//	cv.Volunteer = append(cv.Volunteer, jsonresume.ExperienceID("http://example.com/resume/experience/4"))
//
//	cv.Volunteer = append(cv.Volunteer, jsonresume.ExperienceID("http://example.com/resume/experience/1"))
type ExperienceID jsonld.ID

func SomeExperienceID(value string) ExperienceID {
	return ExperienceID(jsonld.SomeID(value))
}

func (receiver ExperienceID) MarshalJSON() ([]byte, error) {
	return jsonld.ID(receiver).MarshalJSON()
}

func (receiver ExperienceID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeExperience

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ExperienceID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeExperience

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ExperienceID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeExperience

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ExperienceID) ProtoExperience() AnyExperience {
	const _type string = TypeExperience

	return AnyExperience{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

// String makes [ExperienceID] fit the [fmt.Stringer] interface.
func (receiver ExperienceID) String() string {
	return jsonld.ID(receiver).GetElse("")
}
