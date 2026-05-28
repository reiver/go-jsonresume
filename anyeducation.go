package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

// AnyEducation is similar to [Education] except that it is less restrictive about the (ActivityPub/ActivityStreams style) JSON-LD type.
// Where [Education] hard-codes its (ActivityPub/ActivityStreams style) JSON-LD type, AnyEducation allows anything to be loaded — including multiple values, and sub-types.
type AnyEducation struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID    `json:"id,omitempty"`
	Type jsonld.Types `json:"type,omitempty"`

	CoreEducation
}

func (receiver *AnyEducation) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawEducation
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal education")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education type")
				return err
			}
		}
	}

	return receiver.CoreEducation.unmarshalRawEducation(raw, &receiver.ID)
}

func (receiver AnyEducation) ProtoNode() activitypub.AnyNode {
	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyEducation) ProtoEntity() activitypub.AnyEntity {
	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyEducation) ProtoObject() activitypub.AnyObject {
	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreObject: activitypub.CoreObject{
			URL: receiver.URL,
		},
	}
}

func (receiver AnyEducation) ProtoEducation() AnyEducation {
	return receiver
}
