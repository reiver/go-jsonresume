package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

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
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.ID)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal certificate id")
				return err
			}
		}
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

	{
		{
			var bb []byte = []byte(raw.Date)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Date)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal certificate date")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Name)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Name)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal certificate name")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Issuer)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Issuer)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal certificate issuer")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.URL)

			if 0 < len(bb) {
				err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoLink, activitypub.HRef, activitypub.AnyLink](bb, &receiver.URL)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal certificate url")
					return err
				}
			}
		}
	}

	return nil
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
