package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

type CoreBasics struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	EMail    activitypub.Strings `json:"email"`
	Label    activitypub.Strings `json:"label"`
	Phone    activitypub.Strings `json:"phone"`
	Profiles []ProtoProfile      `json:"profiles"`
}
