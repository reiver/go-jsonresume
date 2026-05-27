package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

type AnyPublication struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID    `json:"id,omitempty"`
	Type jsonld.Types `json:"type,omitempty"`

	CorePublication
}

func (receiver *AnyPublication) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawPublication
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal publication")
		return err
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.ID)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal publication id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal publication type")
				return err
			}
		}
	}

	{
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
	}

	return nil
}

func (receiver AnyPublication) ProtoNode() activitypub.AnyNode {
	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyPublication) ProtoEntity() activitypub.AnyEntity {
	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver AnyPublication) ProtoObject() activitypub.AnyObject {
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

func (receiver AnyPublication) ProtoPublication() AnyPublication {
	return receiver
}
