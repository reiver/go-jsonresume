package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreLocation struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Address     nul.Nullable[string] `json:"address,omitempty"`
	City        nul.Nullable[string] `json:"city,omitempty"`
	CountryCode nul.Nullable[string] `json:"countryCode,omitempty"`
	PostalCode  nul.Nullable[string] `json:"postalCode,omitempty"`
	Region      nul.Nullable[string] `json:"region,omitempty"`
}

type rawLocation struct {
	ID          gojson.RawMessage `json:"id"`
	Type        gojson.RawMessage `json:"type"`
	Address     gojson.RawMessage `json:"address"`
	City        gojson.RawMessage `json:"city"`
	CountryCode gojson.RawMessage `json:"countryCode"`
	PostalCode  gojson.RawMessage `json:"postalCode"`
	Region      gojson.RawMessage `json:"region"`
}

func (receiver *CoreLocation) unmarshalRawLocation(raw rawLocation, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal location id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Address)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Address)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal location address")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.City)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.City)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal location city")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.CountryCode)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.CountryCode)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal location countryCode")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.PostalCode)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.PostalCode)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal location postalCode")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Region)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Region)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal location region")
				return err
			}
		}
	}

	return nil
}
