package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

// CoreProject holds the shared fields for [Project] and [AnyProject].
//
// The ProjectType field holds the JSON Resume project category (e.g., "application", "library", "tool") — NOT the JSON-LD type.
// The JSON-LD type is stored in [Project].Type or [AnyProject].Type as "@type".
// This separation exists because JSON Resume's "type" field on projects collides with ActivityPub/ActivityStreams JSON-LD's "type".
type CoreProject struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Description nul.Nullable[string]    `json:"description,omitempty"`
	EndDate     nul.Nullable[string]    `json:"endDate,omitempty"`
	Entity      nul.Nullable[string]    `json:"entity,omitempty"`
	Highlights  activitypub.Strings     `json:"highlights,omitempty"`
	Keywords    activitypub.Strings     `json:"keywords,omitempty"`
	Name        nul.Nullable[string]    `json:"name,omitempty"               jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Roles       activitypub.Strings     `json:"roles,omitempty"`
	StartDate   nul.Nullable[string]    `json:"startDate,omitempty"`
	ProjectType nul.Nullable[string]    `json:"type,omitempty"`              // project category (e.g., "application"), not the JSON-LD type
	URL         []activitypub.ProtoLink `json:"url,omitempty,jsonld.compact" jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}

func (receiver *CoreProject) unmarshalRawProject(raw rawProject, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Description)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Description)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project description")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.EndDate)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.EndDate)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project endDate")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Entity)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Entity)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project entity")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Highlights)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Highlights)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project highlights")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Keywords)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Keywords)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project keywords")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Name)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Name)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project name")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Roles)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Roles)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project roles")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.StartDate)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.StartDate)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project startDate")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.ProjectType)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.ProjectType)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project type")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.URL)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoLink, activitypub.HRef, activitypub.AnyLink](bb, &receiver.URL)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project url")
				return err
			}
		}
	}

	return nil
}

func (receiver *CoreProject) SetURL(url string) {
	if nil == receiver {
		return
	}

	receiver.URL = []activitypub.ProtoLink{activitypub.HRef(url)}
}
