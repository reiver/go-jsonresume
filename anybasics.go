package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

type AnyBasics struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID    `json:"id,omitempty"`
	Type jsonld.Types `json:"type,omitempty"`

	CoreBasics
}

func (receiver *AnyBasics) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawBasics
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal basics")
		return err
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.ID)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal basics id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal basics type")
				return err
			}
		}
	}

	{
		{
			var bb []byte = []byte(raw.EMail)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.EMail)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal basics email")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Image)

			if 0 < len(bb) {
				err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoImageOrProtoLink, activitypub.HRef, activitypub.AnyImage](bb, &receiver.Image)
				if nil != err {
					err = jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoImageOrProtoLink, activitypub.HRef, activitypub.AnyLink](bb, &receiver.Image)
					if nil != err {
						err = erorr.Wrap(err, "failed to json-unmarshal basics image")
						return err
					}
				}
			}
		}

		{
			var bb []byte = []byte(raw.Label)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Label)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal basics label")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Location)

			if 0 < len(bb) {
				err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoLocation, LocationID, AnyLocation](bb, &receiver.Location)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal basics location")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Name)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Name)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal basics name")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Phone)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Phone)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal basics phone")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Profiles)

			if 0 < len(bb) {
				err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoProfile, ProfileID, AnyProfile](bb, &receiver.Profiles)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal basics profiles")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Summary)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Summary)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal basics summary")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.URL)

			if 0 < len(bb) {
				err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoLink, activitypub.HRef, activitypub.AnyLink](bb, &receiver.URL)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal basics url")
					return err
				}
			}
		}
	}

	return nil
}

func (receiver AnyBasics) ProtoNode() activitypub.AnyNode {
	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyBasics) ProtoEntity() activitypub.AnyEntity {
	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver AnyBasics) ProtoObject() activitypub.AnyObject {
	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: receiver.Type,

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
		CoreObject: activitypub.CoreObject{
			Image:   receiver.Image,
			Summary: receiver.Summary,
			URL:     receiver.URL,
		},
	}
}

func (receiver AnyBasics) ProtoBasics() AnyBasics {
	return receiver
}
