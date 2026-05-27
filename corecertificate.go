package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreCertificate struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Date   nul.Nullable[string]    `json:"date,omitempty"`
	Name   nul.Nullable[string]    `json:"name,omitempty"               jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Issuer nul.Nullable[string]    `json:"issuer,omitempty"`
	URL    []activitypub.ProtoLink `json:"url,omitempty,jsonld.compact" jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}

func (receiver *CoreCertificate) unmarshalRawCertificate(raw rawCertificate, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal certificate id")
				return err
			}
		}
	}

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

	return nil
}

func (receiver *CoreCertificate) SetURL(url string) {
	if nil == receiver {
		return
	}

	receiver.URL = []activitypub.ProtoLink{activitypub.HRef(url)}
}
