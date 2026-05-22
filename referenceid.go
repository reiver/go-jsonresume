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
type ReferenceID jsonld.ID

func SomeReferenceID(value string) ReferenceID {
	return ReferenceID(jsonld.SomeID(value))
}

func (receiver ReferenceID) MarshalJSON() ([]byte, error) {
	return jsonld.ID(receiver).MarshalJSON()
}

func (receiver *ReferenceID) UnmarshalJSON(data []byte) error {
	return (*jsonld.ID)(receiver).UnmarshalJSON(data)
}

func (receiver ReferenceID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeReference

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ReferenceID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeReference

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ReferenceID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeReference

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ReferenceID) ProtoReference() AnyReference {
	const _type string = TypeReference

	return AnyReference{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

// String makes [ReferenceID] fit the [fmt.Stringer] interface.
func (receiver ReferenceID) String() string {
	return jsonld.ID(receiver).GetElse("")
}
