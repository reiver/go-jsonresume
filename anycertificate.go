package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

// AnyCertificate is similar to [Certificate] except that it is less restrictive about the (ActivityPub/ActivityStreams style) JSON-LD type.
// Where [Certificate] hard-codes its (ActivityPub/ActivityStreams style) JSON-LD type, AnyCertificate allows anything to be loaded — including multiple values, and sub-types.
type AnyCertificate struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID    `json:"id,omitempty"`
	Type jsonld.Types `json:"type,omitempty"`

	CoreCertificate
}

func (receiver *AnyCertificate) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawCertificate
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal certificate")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal certificate type")
				return err
			}
		}
	}

	return receiver.CoreCertificate.unmarshalRawCertificate(raw, &receiver.ID)
}

func (receiver AnyCertificate) ProtoNode() activitypub.AnyNode {
	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyCertificate) ProtoEntity() activitypub.AnyEntity {
	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver AnyCertificate) ProtoObject() activitypub.AnyObject {
	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
		CoreObject: activitypub.CoreObject{
			URL: receiver.URL,
		},
	}
}

func (receiver AnyCertificate) ProtoCertificate() AnyCertificate {
	return receiver
}
