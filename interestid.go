package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// InterestID implements a referenced item in the JSON Resume "interests" fields array.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"interests": [
//		"http://example.com/resume/interest/photography",
//		"http://example.com/resume/interest/woodworking"
//	],
//
// InterestID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Interests = append(cv.Interests, jsonresume.InterestID("http://example.com/resume/interest/photography"))
//	
//	cv.Interests = append(cv.Interests, jsonresume.InterestID("http://example.com/resume/interest/woodworking"))
type InterestID string

func (receiver InterestID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeInterest

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver InterestID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeInterest

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver InterestID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeInterest

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}
