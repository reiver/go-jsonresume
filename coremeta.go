package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreMeta struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Canonical    nul.Nullable[string] `json:"canonical,omitempty"`
	LastModified nul.Nullable[string] `json:"lastModified,omitempty"`
	Version      nul.Nullable[string] `json:"version,omitempty"`
}

type rawMeta struct {
	ID           gojson.RawMessage `json:"id"`
	Type         gojson.RawMessage `json:"type"`
	Canonical    gojson.RawMessage `json:"canonical"`
	LastModified gojson.RawMessage `json:"lastModified"`
	Version      gojson.RawMessage `json:"version"`
}

func (receiver *CoreMeta) unmarshalRawMeta(raw rawMeta, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal meta id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Canonical)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Canonical)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal meta canonical")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.LastModified)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.LastModified)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal meta lastModified")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Version)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Version)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal meta version")
				return err
			}
		}
	}

	return nil
}
