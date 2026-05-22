package jsonresume

import (
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

func (receiver *JSONResume) AppendResumeIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Resume = append(receiver.Resume, SomeResumeID(iri))
}

func (receiver *JSONResume) SetResumeIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Resume = []ProtoResume{SomeResumeID(iri)}
}
