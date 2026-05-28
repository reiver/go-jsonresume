package jsonresume

import (
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreLanguage struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Fluency  nul.Nullable[string] `json:"fluency,omitempty"`
	Language nul.Nullable[string] `json:"language,omitempty"`
}

func (receiver *CoreLanguage) unmarshalRawLanguage(raw rawLanguage, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	if nil != id {
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal language id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Fluency)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Fluency)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal language fluency")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Language)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Language)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal language language")
				return err
			}
		}
	}

	return nil
}
