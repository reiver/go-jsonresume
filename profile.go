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
//			"url": "https://mastodon.example/@joeblow"
//		},
//		{
//			"network": "Pixelfed",
//			"username": "joeblow",
//			"url": "https://pixelfed.example/joeblow"
//		},
//	],
//
// Profile is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var basics jsonresume.Basics
//	
//	// ...
//	
//	basics.Profiles = append(basics.Profiles, jsonresume.Profile{
//		Network:   nul.Something("Mastodon"),
//		UserName:  nul.Something("joeblow"),
//		URL:       activitypub.HRef("https://mastodon.example/@joeblow"),
//	})
//	
//	basics.Profiles = append(basics.Profiles, jsonresume.Profile{
//		Network:   nul.Something("Pixelfed"),
//		UserName:  nul.Something("joeblow"),
//		URL:       activitypub.HRef("https://pixelfed.example/joeblow"),
//	})
//
// Note that you should use Profile for marshaling but not unmarshaling.
// For unmarshaling instead use [AnyProfile].
//
// See also:
//
//	• [AnyProfile]
//	• [CoreProfile]
//	• [ProfileID]
//	• [ProtoProfile]
//	• [TypeProfile]
type Profile struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Profile"`

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
	}
}

func (receiver Profile) ProtoObject() activitypub.AnyObject {
	const _type string = TypeProfile

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreObject: activitypub.CoreObject{
			URL: receiver.URL,
		},
	}
}

func (receiver Profile) ProtoProfile() AnyProfile {
	const _type string = TypeProfile

	return AnyProfile{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreProfile: receiver.CoreProfile,
	}
}
