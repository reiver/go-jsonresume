package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoPublication represents something that is any type 'Publication', including sub-types.
//
// See also:
//
//	• [AnyPublication]
//	• [Publication]
//	• [PublicationID]
type ProtoPublication interface {
	activitypub.ProtoObject
	ProtoPublication() AnyPublication
}

var (
	_ ProtoPublication = Publication{}
	_ ProtoPublication = PublicationID("")
)
