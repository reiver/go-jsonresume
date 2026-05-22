package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Profile implements an embedded item in the JSON Resume "profiles" fields array.
//
// In a JSON Resume document this might look like:
//
//	"profiles": [
//		{
//			"network": "Mastodon",
//			"username": "joeblow",
//			"reference": "https://mastodon.example/@joeblow"
//		},
//		{
//			"network": "Pixelfed",
//			"username": "joeblow",
//			"reference": "https://pixelfed.example/joeblow"
//		},
//	],
//
// Profile is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Profiles = append(cv.Profiles, jsonresume.Profile{
//		Network:   nul.Something("Mastodon"),
//		UserName:  nul.Something("joeblow"),
//		Reference: activitypub.HRef("https://mastodon.example/@joeblow"),
//	})
//	
//	cv.Profiles = append(cv.Profiles, jsonresume.Profile{
//		Name:      nul.Something("Pixelfed"),
//		UserName:  nul.Something("joeblow"),
//		Reference: activitypub.HRef("https://pixelfed.example/joeblow"),
//	})
type Profile struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Profile"`

	activitypub.CoreEntity
	activitypub.CoreObject
	CoreProfile
}

func (receiver Profile) ProtoNode() activitypub.AnyNode {
	const _type string = TypeProfile

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Profile) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeProfile

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Profile) ProtoObject() activitypub.AnyObject {
	const _type string = TypeProfile

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

func (receiver Profile) ProtoProfile() AnyProfile {
	const _type string = TypeProfile

	var result = AnyProfile{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity:  receiver.CoreEntity,
		CoreObject:  receiver.CoreObject,
		CoreProfile: receiver.CoreProfile,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}
