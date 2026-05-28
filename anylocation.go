package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

// AnyLocation is similar to [Location] except that it is less restrictive about the (ActivityPub/ActivityStreams style) JSON-LD type.
// Where [Location] hard-codes its (ActivityPub/ActivityStreams style) JSON-LD type, AnyLocation allows anything to be loaded — including multiple values, and sub-types.
type AnyLocation struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID    `json:"id,omitempty"`
	Type jsonld.Types `json:"type,omitempty"`

	CoreLocation
}

func (receiver *AnyLocation) UnmarshalJSON(bytes []byte) error {
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
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal location type")
				return err
			}
		}
	}

	return receiver.CoreLocation.unmarshalRawLocation(raw, &receiver.ID)
}

func (receiver AnyLocation) ProtoNode() activitypub.AnyNode {
	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyLocation) ProtoEntity() activitypub.AnyEntity {
	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyLocation) ProtoObject() activitypub.AnyObject {
	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyLocation) ProtoLocation() AnyLocation {
	return receiver
}
