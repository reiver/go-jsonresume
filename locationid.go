package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// LocationID implements a referenced item in the JSON Resume "location" field.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"location": "http://example.com/resume/location",
//
// LocationID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Locations = jsonresume.LocationID("http://example.com/resume/location")
type LocationID string

func (receiver LocationID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeLocation

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver LocationID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeLocation

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver LocationID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeLocation

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver LocationID) ProtoLocation() AnyLocation {
	const _type string = TypeLocation

	return AnyLocation{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}
