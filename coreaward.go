package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreAward struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Awarder nul.Nullable[string] `json:"awarder,omitempty"`
	Date    nul.Nullable[string] `json:"date,omitempty"`
	Summary nul.Nullable[string] `json:"summary,omitempty" jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Title   nul.Nullable[string] `json:"title,omitempty"`
}

type rawAward struct {
	ID      gojson.RawMessage `json:"id"`
	Type    gojson.RawMessage `json:"type"`
	Awarder gojson.RawMessage `json:"awarder"`
	Date    gojson.RawMessage `json:"date"`
	Summary gojson.RawMessage `json:"summary"`
	Title   gojson.RawMessage `json:"title"`
}

func (receiver *CoreAward) unmarshalRawAward(raw rawAward, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	if nil != id {
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal award id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Awarder)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Awarder)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal award awarder")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Date)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Date)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal award date")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Summary)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Summary)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal award summary")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Title)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Title)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal award title")
				return err
			}
		}
	}

	return nil
}
