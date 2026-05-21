package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// AwardID implements a referenced item in the JSON Resume "awards" fields array.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"awards": [
//		"http://example.com/resume/award/best-employee-2024",
//		"http://example.com/resume/award/acme-excellence-2021"
//	],
//
// AwardID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.References = append(cv.References, jsonresume.AwardID("http://example.com/resume/award/best-employee-2024"))
//	
//	cv.References = append(cv.References, jsonresume.AwardID("http://example.com/resume/award/acme-excellence-2021"))
type AwardID string

func (receiver AwardID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeAward

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver AwardID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeAward

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver AwardID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeAward

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver AwardID) ProtoAward() AnyAward {
	const _type string = TypeAward

	return AnyAward{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

