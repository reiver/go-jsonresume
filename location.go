package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
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
//
// Location is for marshaling with a fixed type of "Location".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Location", "cv:Location", or "https://w3id.org/fep/6158#Location".
//
// For unmarshaling that accepts any type value, use [AnyLocation] instead.
//
// See also:
//
//	• [AnyLocation]
//	• [CoreLocation]
//	• [LocationID]
//	• [ProtoLocation]
//	• [TypeLocation]
type Location struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Location"`

	CoreLocation
}

func (receiver *Location) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawLocation
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal location")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			var typeValue string
			err := gojson.Unmarshal(bb, &typeValue)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal location type")
				return err
			}

			switch typeValue {
			case TypeLocation, CompactTypeLocation, ExpandedTypeLocation:
				// OK
			default:
				return erorr.Errorf("jsonresume: unexpected type for location: %q", typeValue)
			}
		}
	}

	return receiver.CoreLocation.unmarshalRawLocation(raw, &receiver.ID)
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
	}
}

func (receiver Location) ProtoObject() activitypub.AnyObject {
	const _type string = TypeLocation

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Location) ProtoLocation() AnyLocation {
	const _type string = TypeLocation

	return AnyLocation{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreLocation: receiver.CoreLocation,
	}
}
