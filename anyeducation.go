package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

type AnyEducation struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID    `json:"id,omitempty"`
	Type jsonld.Types `json:"type,omitempty"`

	CoreEducation
}

func (receiver *AnyEducation) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawEducation
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal education")
		return err
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.ID)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education type")
				return err
			}
		}
	}

	{
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
	}

	return nil
}

func (receiver AnyEducation) ProtoNode() activitypub.AnyNode {
	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyEducation) ProtoEntity() activitypub.AnyEntity {
	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyEducation) ProtoObject() activitypub.AnyObject {
	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreObject: activitypub.CoreObject{
			URL: receiver.URL,
		},
	}
}

func (receiver AnyEducation) ProtoEducation() AnyEducation {
	return receiver
}
