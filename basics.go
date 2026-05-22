package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Basics implements an embedded item in the JSON Resume "basics" field.
//
// In a JSON Resume document this might look like:
//
//	"basics": {
//		"name": "Joe Blow",
//		"label": "Programmer",
//		"image": "https://joeblow.example/photos/joeblow.jpeg",
//		"email": "joeblow@example.com",
//		"phone": "(604) 555-1234",
//		"url": "https://joeblow.example",
//		"summary": "CTO, Experienced Programmer",
//		"location": {
//			"countryCode": "CA",
//			"region": "British Columbia",
//			"city": "Vancouver",
//			"address": "1234 Second Street",
//			"postalCode": "H0H 0H0"
//		},
//		"profiles": [
//			{
//				"network": "Mastodon",
//				"username": "joeblow",
//				"url": "https://mastodon.example/@joeblow"
//			},
//			{
//				"network": "Pixelfed",
//				"username": "joeblow",
//				"url": "https://pixelfed.example/joeblow"
//			}
//		]
//	},
//
// Basics is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var basics json.Basics
//	
//	// ...
//	
//	basics.Location = append(basics.Location, jsonresume.Location{
//		CountryCode: nul.Something("CA"),
//		Region:      nul.Something("British Columbia"),
//		City:        nul.Something("Vancouver"),
//		Address:     nul.Something("1234 Second Street"),
//		PostalCode:  nul.Something("H0H 0H0"),
//	})
//	
//	// ...
//	
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Basics = basics
type Basics struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Basics"`

	activitypub.CoreEntity
	activitypub.CoreObject
	CoreBasics
}

func (receiver Basics) ProtoNode() activitypub.AnyNode {
	const _type string = TypeBasics

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Basics) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeBasics

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Basics) ProtoObject() activitypub.AnyObject {
	const _type string = TypeBasics

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

func (receiver Basics) ProtoBasics() AnyBasics {
	const _type string = TypeBasics

	var result = AnyBasics{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
		CoreObject: receiver.CoreObject,
		CoreBasics:  receiver.CoreBasics,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}
