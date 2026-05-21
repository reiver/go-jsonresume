package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoReference represents something that is any type 'Reference', including sub-types.
//
// See also:
//
//	• [AnyReference]
//	• [Reference]
//	• [ReferenceID]
type ProtoReference interface {
	activitypub.ProtoObject
	ProtoReference() AnyReference
}

var (
	_ ProtoReference = Reference{}
	_ ProtoReference = ReferenceID("")
)
