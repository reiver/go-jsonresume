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
// BasicsID is an implementation of a JSON-LD reference for the basics object.
//
// Example usage:
//
//	var cv jsonresume.Resume
//
//	// ...
//
//	cv.Basics = jsonresume.BasicsID("http://example.com/resume/basics")
type BasicsID jsonld.ID

func SomeBasicsID(value string) BasicsID {
	return BasicsID(jsonld.SomeID(value))
}

func (receiver BasicsID) MarshalJSON() ([]byte, error) {
	return jsonld.ID(receiver).MarshalJSON()
}

func (receiver BasicsID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeBasics

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver BasicsID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeBasics

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver BasicsID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeBasics

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver BasicsID) ProtoBasics() AnyBasics {
	const _type string = TypeBasics

	return AnyBasics{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

// String makes [BasicsID] fit the [fmt.Stringer] interface.
func (receiver BasicsID) String() string {
	return jsonld.ID(receiver).GetElse("")
}
