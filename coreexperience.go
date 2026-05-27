package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreExperience struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Description  nul.Nullable[string]    `json:"description,omitempty"`
	EndDate      nul.Nullable[string]    `json:"endDate,omitempty"`
	Highlights   activitypub.Strings     `json:"highlights,omitempty"`
	Location     nul.Nullable[string]    `json:"location,omitempty"`
	Name         nul.Nullable[string]    `json:"name,omitempty"                    jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Organization nul.Nullable[string]    `json:"organization,omitempty"`
	Position     activitypub.Strings     `json:"position,omitempty,jsonld.compact"`
	StartDate    nul.Nullable[string]    `json:"startDate,omitempty"`
	Summary      nul.Nullable[string]    `json:"summary,omitempty"                 jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	URL          []activitypub.ProtoLink `json:"url,omitempty,jsonld.compact"      jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}

func (receiver *CoreExperience) unmarshalRawExperience(raw rawExperience, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Description)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Description)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience description")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.EndDate)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.EndDate)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience endDate")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Highlights)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Highlights)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience highlights")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Location)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Location)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience location")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Name)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Name)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience name")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Organization)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Organization)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience organization")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Position)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Position)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience position")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.StartDate)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.StartDate)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience startDate")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Summary)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Summary)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience summary")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.URL)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoLink, activitypub.HRef, activitypub.AnyLink](bb, &receiver.URL)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience url")
				return err
			}
		}
	}

	return nil
}

func (receiver *CoreExperience) SetURL(url string) {
	if nil == receiver {
		return
	}

	receiver.URL = []activitypub.ProtoLink{activitypub.HRef(url)}
}
