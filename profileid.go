package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// ProfileID implements a referenced item in the JSON Resume "profiles" fields array.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"profiles": [
//		"http://example.com/resume/profile/mastodon",
//		"http://example.com/resume/profile/pixelfed"
//	],
//
// ProfileID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.References = append(cv.References, jsonresume.ProfileID("http://example.com/resume/profile/mastodon"))
//	
//	cv.References = append(cv.References, jsonresume.ProfileID("http://example.com/resume/profile/pixelfed"))
type ProfileID string

func (receiver ProfileID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeProfile

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ProfileID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeProfile

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ProfileID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeProfile

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}
