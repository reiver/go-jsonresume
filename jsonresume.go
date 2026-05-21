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
