package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
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
// Basics is an implementation of the JSON Resume basics object.
//
// Example usage:
//
//	var basics jsonresume.Basics
//
//	// ...
//
//	basics.Location = jsonresume.Location{
//		CountryCode: nul.Something("CA"),
//		Region:      nul.Something("British Columbia"),
//		City:        nul.Something("Vancouver"),
//		Address:     nul.Something("1234 Second Street"),
//		PostalCode:  nul.Something("H0H 0H0"),
//	}
//	
//	// ...
//	
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Basics = basics
//
// Basics is for marshaling with a fixed type of "Basics".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Basics", "cv:Basics", or "https://w3id.org/fep/6158#Basics".
//
// For unmarshaling that accepts any type value, use [AnyBasics] instead.
//
// See also:
//
//	• [AnyBasics]
//	• [BasicsID]
//	• [CoreBasics]
//	• [ProtoBasics]
//	• [TypeBasics]
type Basics struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Basics"`

	CoreBasics
}

func (receiver *Basics) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawBasics
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal basics")
		return err
	}

	if err := validateTypeForUnmarshalJSON([]byte(raw.Type), "basics", TypeBasics, CompactTypeBasics, ExpandedTypeBasics); nil != err {
		return err
	}

	return receiver.CoreBasics.unmarshalRawBasics(raw, &receiver.ID)
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

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver Basics) ProtoObject() activitypub.AnyObject {
	const _type string = TypeBasics

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
		CoreObject: activitypub.CoreObject{
			Image:   receiver.Image,
			Summary: receiver.Summary,
			URL:     receiver.URL,
		},
	}
}

func (receiver Basics) ProtoBasics() AnyBasics {
	const _type string = TypeBasics

	return AnyBasics{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreBasics: receiver.CoreBasics,
	}
}
