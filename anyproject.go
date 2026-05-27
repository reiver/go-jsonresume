package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

type AnyProject struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID    `json:"id,omitempty"`
	Type jsonld.Types `json:"type,omitempty"`

	CoreProject
}

func (receiver *AnyProject) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawProject
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal project")
		return err
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.ID)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project type")
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
			var bb []byte = []byte(raw.URL)

			if 0 < len(bb) {
				err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoLink, activitypub.HRef, activitypub.AnyLink](bb, &receiver.URL)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal project url")
					return err
				}
			}
		}
	}

	return nil
}

func (receiver AnyProject) ProtoNode() activitypub.AnyNode {
	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyProject) ProtoEntity() activitypub.AnyEntity {
	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver AnyProject) ProtoObject() activitypub.AnyObject {
	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
		CoreObject: activitypub.CoreObject{
			URL: receiver.URL,
		},
	}
}

func (receiver AnyProject) ProtoProject() AnyProject {
	return receiver
}
