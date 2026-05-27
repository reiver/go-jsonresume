package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
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
