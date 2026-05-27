package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// MetaID implements a referenced item in the JSON Resume "meta" field.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"meta": "http://example.com/resume/meta",
//
// MetaID is an implementation of a JSON-LD reference for the meta object.
//
// Example usage:
//
//	var cv jsonresume.Resume
//
//	// ...
//
//	cv.Meta = jsonresume.MetaID("http://example.com/resume/meta")
type MetaID jsonld.ID

func SomeMetaID(value string) MetaID {
	return MetaID(jsonld.SomeID(value))
}

func (receiver MetaID) MarshalJSON() ([]byte, error) {
	return jsonld.ID(receiver).MarshalJSON()
}

func (receiver *MetaID) UnmarshalJSON(data []byte) error {
	return (*jsonld.ID)(receiver).UnmarshalJSON(data)
}

func (receiver MetaID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeMeta

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver MetaID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeMeta

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver MetaID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeMeta

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver MetaID) ProtoMeta() AnyMeta {
	const _type string = TypeMeta

	return AnyMeta{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

// String makes [MetaID] fit the [fmt.Stringer] interface.
func (receiver MetaID) String() string {
	return jsonld.ID(receiver).GetElse("")
}
