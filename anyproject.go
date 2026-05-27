package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

type AnyProject struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID      `json:"id,omitempty"`
	Type jsonld.Types `json:"@type,omitempty"`

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
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project @type")
				return err
			}
		}
	}

	return receiver.CoreProject.unmarshalRawProject(raw, &receiver.ID)
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
