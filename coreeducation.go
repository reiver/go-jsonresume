package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreEducation struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Area        activitypub.Strings     `json:"area,omitempty"`
	Courses     activitypub.Strings     `json:"courses,omitempty"`
	EndDate     nul.Nullable[string]    `json:"endDate,omitempty"`
	Institution nul.Nullable[string]    `json:"institution,omitempty"`
	Score       nul.Nullable[string]    `json:"score,omitempty"`
	StartDate   nul.Nullable[string]    `json:"startDate,omitempty"`
	StudyType   activitypub.Strings     `json:"studyType,omitempty"`
	URL         []activitypub.ProtoLink `json:"url,omitempty,jsonld.compact" jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}

func (receiver *CoreEducation) unmarshalRawEducation(raw rawEducation, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Area)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Area)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education area")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Courses)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Courses)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education courses")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.EndDate)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.EndDate)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education endDate")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Institution)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Institution)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education institution")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Score)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Score)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education score")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.StartDate)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.StartDate)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education startDate")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.StudyType)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.StudyType)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education studyType")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.URL)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoLink, activitypub.HRef, activitypub.AnyLink](bb, &receiver.URL)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education url")
				return err
			}
		}
	}

	return nil
}

func (receiver *CoreEducation) SetURL(url string) {
	if nil == receiver {
		return
	}

	receiver.URL = []activitypub.ProtoLink{activitypub.HRef(url)}
}
