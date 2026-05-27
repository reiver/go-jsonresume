package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoBasics represents something that is any type 'Basics', including sub-types.
//
// See also:
//
//	• [AnyBasics]
//	• [Basics]
//	• [BasicsID]
type ProtoBasics interface {
	activitypub.ProtoObject
	ProtoBasics() AnyBasics
}

var (
	_ ProtoBasics = AnyBasics{}
	_ ProtoBasics = Basics{}
	_ ProtoBasics = BasicsID{}
)
