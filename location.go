package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Location implements an embedded item in the JSON Resume "location" fields array.
//
// In a JSON Resume document this might look like:
//
//	"location": {
//		"countryCode": "CA",
//		"region": "British Columbia",
//		"city": "Vancouver",
//		"address": "1234 Second Street",
//		"postalCode": "H0H 0H0"
//	},
//
// Location is an implementation of the JSON Resume location object.
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
type Location struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Location"`

	activitypub.CoreEntity
	activitypub.CoreObject
	CoreLocation
}

func (receiver Location) ProtoNode() activitypub.AnyNode {
	const _type string = TypeLocation

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Location) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeLocation

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Location) ProtoObject() activitypub.AnyObject {
	const _type string = TypeLocation

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

func (receiver Location) ProtoLocation() AnyLocation {
	const _type string = TypeLocation

	var result = AnyLocation{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
		CoreObject: receiver.CoreObject,
		CoreLocation:  receiver.CoreLocation,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}
