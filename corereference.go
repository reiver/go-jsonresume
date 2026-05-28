package jsonresume

import (
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreReference struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Name      nul.Nullable[string] `json:"name,omitempty"      jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Reference nul.Nullable[string] `json:"reference,omitempty"`
}

func (receiver *CoreReference) unmarshalRawReference(raw rawReference, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	if nil != id {
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal reference id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Name)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Name)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal reference name")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Reference)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Reference)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal reference reference")
				return err
			}
		}
	}

	return nil
}
