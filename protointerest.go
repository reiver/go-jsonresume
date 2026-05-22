package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoInterest represents something that is any type 'Interest', including sub-types.
//
// See also:
//
//	• [AnyInterest]
//	• [Interest]
//	• [InterestID]
type ProtoInterest interface {
	activitypub.ProtoObject
	ProtoInterest() AnyInterest
}

var (
	_ ProtoInterest = Interest{}
	_ ProtoInterest = InterestID{}
)
