package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

type CorePublication struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Name        nul.Nullable[string]    `json:"name,omitempty"        jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	Publisher   nul.Nullable[string]    `json:"publisher,omitempty"`
	ReleaseDate nul.Nullable[string]    `json:"releaseDate,omitempty"`
	Summary     nul.Nullable[string]    `json:"summary,omitempty"     jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
	URL         []activitypub.ProtoLink `json:"url"         jsonld.namespace:"http://www.w3.org/ns/activitystreams" jsonld.prefix:"as"`
}
