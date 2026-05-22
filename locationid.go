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
// LocationID is an implementation of a JSON-LD reference for the location object.
//
// Example usage:
//
//	var cv jsonresume.Resume
//
//	// ...
//
//	cv.Locations = jsonresume.LocationID("http://example.com/resume/location")
type LocationID jsonld.ID

func SomeLocationID(value string) LocationID {
	return LocationID(jsonld.SomeID(value))
}

func (receiver LocationID) MarshalJSON() ([]byte, error) {
	return jsonld.ID(receiver).MarshalJSON()
}

func (receiver LocationID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeLocation

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver LocationID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeLocation

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver LocationID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeLocation

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver LocationID) ProtoLocation() AnyLocation {
	const _type string = TypeLocation

	return AnyLocation{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

// String makes [LocationID] fit the [fmt.Stringer] interface.
func (receiver LocationID) String() string {
	return jsonld.ID(receiver).GetElse("")
}
