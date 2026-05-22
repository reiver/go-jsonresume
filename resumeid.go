package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// ResumeID implements a resume by reference.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In ActivityPub this might look like:
//
//
//	{
//		"@context": [
//			"https://www.w3.org/ns/activitystreams",
//			"https://w3id.org/fep/6158",
//		],
//
//		...
//
//		"resume": [
//			"http://example.com/resume/executive",
//			"http://example.com/resume/programmer"
//		],
//
//		...
//	}
//
// ResumeID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var obj jsonresume.JSONResume
//	
//	// ...
//	
//	obj.Resume = append(cv.Resume, jsonresume.ResumeID("http://example.com/resume/executive"))
//	
//	obj.Resume = append(cv.Resume, jsonresume.ResumeID("http://example.com/resume/programmer"))
type ResumeID string

func (receiver ResumeID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeResume

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ResumeID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeResume

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ResumeID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeResume

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ResumeID) ProtoResume() AnyResume {
	const _type string = TypeResume

	return AnyResume{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

// String makes [ResumeID] fit the [fmt.Stringer] interface.
func (receiver ResumeID) String() string {
	return string(receiver)
}
