package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// ReferenceID implements a referenced item in the JSON Resume "references" fields array.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"references": [
//		"http://example.com/resume/reference/jane-doe",
//		"http://example.com/resume/reference/bob-smith"
//	],
//
// ReferenceID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.References = append(cv.References, jsonresume.ReferenceID("http://example.com/resume/reference/jane-doe"))
//	
//	cv.References = append(cv.References, jsonresume.ReferenceID("http://example.com/resume/reference/bob-smith"))
type ReferenceID string

func (receiver ReferenceID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeReference

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ReferenceID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeReference

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ReferenceID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeReference

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}
