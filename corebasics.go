package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreBasics struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	EMail    activitypub.Strings                 `json:"email"`
	Image    []activitypub.ProtoImageOrProtoLink `json:"image"    jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Label    activitypub.Strings                 `json:"label"`
	Location []ProtoLocation                     `json:"location"`
	Name     nul.Nullable[string]                `json:"name,omitempty"     jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Phone    activitypub.Strings                 `json:"phone"`
	Profiles []ProtoProfile                      `json:"profiles"`
	Summary  nul.Nullable[string]                `json:"summary,omitempty"  jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	URL      []activitypub.ProtoLink             `json:"url"      jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}
