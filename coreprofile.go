package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreProfile struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Network  nul.Nullable[string]    `json:"network,omitempty"`
	UserName nul.Nullable[string]    `json:"username,omitempty"`
	URL      []activitypub.ProtoLink `json:"url,omitempty,jsonld.compact" jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}

func (receiver *CoreProfile) unmarshalRawProfile(raw rawProfile, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal profile id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Network)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Network)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal profile network")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.UserName)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.UserName)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal profile username")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.URL)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoLink, activitypub.HRef, activitypub.AnyLink](bb, &receiver.URL)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal profile url")
				return err
			}
		}
	}

	return nil
}
