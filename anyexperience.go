package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

type AnyExperience struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID    `json:"id,omitempty"`
	Type jsonld.Types `json:"type,omitempty"`

	CoreExperience
}

func (receiver *AnyExperience) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawExperience
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal experience")
		return err
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.ID)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal experience type")
				return err
			}
		}
	}

	{
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
	}

	return nil
}

func (receiver AnyExperience) ProtoNode() activitypub.AnyNode {
	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyExperience) ProtoEntity() activitypub.AnyEntity {
	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver AnyExperience) ProtoObject() activitypub.AnyObject {
	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
		CoreObject: activitypub.CoreObject{
			Summary: receiver.Summary,
			URL:     receiver.URL,
		},
	}
}

func (receiver AnyExperience) ProtoExperience() AnyExperience {
	return receiver
}
