package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

// AnyBasics is similar to [Basics] except that it is less restrictive about the (ActivityPub/ActivityStreams style) JSON-LD type.
// Where [Basics] hard-codes its (ActivityPub/ActivityStreams style) JSON-LD type, AnyBasics allows anything to be loaded — including multiple values, and sub-types.
type AnyBasics struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID    `json:"id,omitempty"`
	Type jsonld.Types `json:"type,omitempty"`

	CoreBasics
}

func (receiver *AnyBasics) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawBasics
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal basics")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal basics type")
				return err
			}
		}
	}

	return receiver.CoreBasics.unmarshalRawBasics(raw, &receiver.ID)
}

func (receiver AnyBasics) ProtoNode() activitypub.AnyNode {
	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyBasics) ProtoEntity() activitypub.AnyEntity {
	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver AnyBasics) ProtoObject() activitypub.AnyObject {
	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: receiver.Type,

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

func (receiver AnyBasics) ProtoBasics() AnyBasics {
	return receiver
}
