package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

type Resume struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Resume"`

	activitypub.CoreEntity
	activitypub.CoreObject

}

func (receiver Resume) ProtoNode() activitypub.AnyNode {
	const _type string = TypeResume

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Resume) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeResume

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Resume) ProtoObject() activitypub.AnyObject {
	const _type string = TypeResume

	var result = activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
		CoreObject: receiver.CoreObject,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}
