package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// PublicationID implements a referenced item in the JSON Resume "publications" fields array.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"publications": [
//		"http://example.com/resume/publication/distributed-systems",
//		"http://example.com/resume/publication/video-compression"
//	],
//
// PublicationID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Publications = append(cv.Publications, jsonresume.PublicationID("http://example.com/resume/publication/distributed-systems"))
//	cv.Publications = append(cv.Publications, jsonresume.PublicationID("http://example.com/resume/publication/video-compression"))
type PublicationID string

func (receiver PublicationID) ProtoNode() activitypub.AnyNode {
	const _type string = TypePublication

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver PublicationID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypePublication

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver PublicationID) ProtoObject() activitypub.AnyObject {
	const _type string = TypePublication

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver PublicationID) ProtoPublication() AnyPublication {
	const _type string = TypePublication

	return AnyPublication{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

// String makes [PublicationID] fit the [fmt.Stringer] interface.
func (receiver PublicationID) String() string {
	return string(receiver)
}
