package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CorePublication struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Name        nul.Nullable[string]    `json:"name,omitempty"               jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Publisher   nul.Nullable[string]    `json:"publisher,omitempty"`
	ReleaseDate nul.Nullable[string]    `json:"releaseDate,omitempty"`
	Summary     nul.Nullable[string]    `json:"summary,omitempty"            jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	URL         []activitypub.ProtoLink `json:"url,omitempty,jsonld.compact" jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}

func (receiver *CorePublication) unmarshalRawPublication(raw rawPublication, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal publication id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Name)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Name)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal publication name")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Publisher)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Publisher)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal publication publisher")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.ReleaseDate)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.ReleaseDate)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal publication releaseDate")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Summary)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Summary)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal publication summary")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.URL)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoLink, activitypub.HRef, activitypub.AnyLink](bb, &receiver.URL)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal publication url")
				return err
			}
		}
	}

	return nil
}
