package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

// JSONResume lets you add "resume" field to an ActivityPub actor.
//
// See also:
//
//	• [Resume]
//	• [ResumeID]
type JSONResume struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Resume []ProtoResume `json:"resume"`
}

func (receiver *JSONResume) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw struct {
		Resume gojson.RawMessage `json:"resume"`
	}

	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		return erorr.Wrap(err, "failed to json-unmarshal json-resume")
	}

	{
		var bb []byte = []byte(raw.Resume)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoResume, ResumeID, AnyResume](bb, &receiver.Resume)
			if nil != err {
				return erorr.Wrap(err, "failed to json-unmarshal json-resume resume")
			}
		}
	}

	return nil
}

func (receiver *JSONResume) AppendResumeID(id string) {
	if nil == receiver {
		return
	}

	receiver.Resume = append(receiver.Resume, SomeResumeID(id))
}

func (receiver *JSONResume) SetResumeID(id string) {
	if nil == receiver {
		return
	}

	receiver.Resume = []ProtoResume{SomeResumeID(id)}
}
