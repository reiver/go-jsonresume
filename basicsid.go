package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// BasicsID implements a referenced item in the JSON Resume "basics" field.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"basics": "http://example.com/resume/basics",
//
// BasicsID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Basics = activitypub.HRef("http://example.com/resume/reference/jane-doe")
type BasicsID string

func (receiver BasicsID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeBasics

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver BasicsID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeBasics

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver BasicsID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeBasics

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver BasicsID) ProtoBasics() AnyBasics {
	const _type string = TypeBasics

	return AnyBasics{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}
