package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoAward represents something that is any type 'Award', including sub-types.
//
// See also:
//
//	• [AnyAward]
//	• [Award]
//	• [AwardID]
type ProtoAward interface {
	activitypub.ProtoObject
	ProtoAward() AnyAward
}

var (
	_ ProtoAward = Award{}
	_ ProtoAward = AwardID("")
)
