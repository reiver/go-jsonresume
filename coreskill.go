package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreSkill struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Keywords activitypub.Strings  `json:"keywords,omitempty"`
	Level    nul.Nullable[string] `json:"level,omitempty"`
	Name     nul.Nullable[string] `json:"name,omitempty"     jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}

type rawSkill struct {
	ID       gojson.RawMessage `json:"id"`
	Type     gojson.RawMessage `json:"type"`
	Keywords gojson.RawMessage `json:"keywords"`
	Level    gojson.RawMessage `json:"level"`
	Name     gojson.RawMessage `json:"name"`
}

func (receiver *CoreSkill) unmarshalRawSkill(raw rawSkill, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal skill id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Keywords)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Keywords)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal skill keywords")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Level)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Level)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal skill level")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Name)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Name)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal skill name")
				return err
			}
		}
	}

	return nil
}
