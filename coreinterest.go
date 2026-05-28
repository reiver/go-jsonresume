package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreInterest struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Name     nul.Nullable[string] `json:"name,omitempty"     jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Keywords activitypub.Strings  `json:"keywords,omitempty"`
}

type rawInterest struct {
	ID       gojson.RawMessage `json:"id"`
	Type     gojson.RawMessage `json:"type"`
	Name     gojson.RawMessage `json:"name"`
	Keywords gojson.RawMessage `json:"keywords"`
}

func (receiver *CoreInterest) unmarshalRawInterest(raw rawInterest, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	if nil != id {
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal interest id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Name)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Name)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal interest name")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Keywords)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Keywords)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal interest keywords")
				return err
			}
		}
	}

	return nil
}
