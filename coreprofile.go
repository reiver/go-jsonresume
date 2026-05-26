package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CoreProfile struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Network  nul.Nullable[string]    `json:"network,omitempty"`
	UserName nul.Nullable[string]    `json:"username,omitempty"`
	URL      []activitypub.ProtoLink `json:"url"      jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}
