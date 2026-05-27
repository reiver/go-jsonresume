package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreBasics struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	EMail    activitypub.Strings                 `json:"email,omitempty,jsonld.compact"`
	Image    []activitypub.ProtoImageOrProtoLink `json:"image,omitempty,jsonld.compact"    jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Label    activitypub.Strings                 `json:"label,omitempty,jsonld.compact"`
	Location []ProtoLocation                     `json:"location,omitempty,jsonld.compact"`
	Name     nul.Nullable[string]                `json:"name,omitempty"                    jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Phone    activitypub.Strings                 `json:"phone,omitempty,jsonld.compact"`
	Profiles []ProtoProfile                      `json:"profiles,omitempty"`
	Summary  nul.Nullable[string]                `json:"summary,omitempty"                 jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	URL      []activitypub.ProtoLink             `json:"url,omitempty,jsonld.compact"      jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}

func (receiver *CoreBasics) unmarshalRawBasics(raw rawBasics, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal basics id")
				return err
			}
		}
	}

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

	return nil
}
