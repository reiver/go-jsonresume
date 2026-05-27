package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoMeta represents something that is any type 'Meta', including sub-types.
//
// See also:
//
//	• [AnyMeta]
//	• [Meta]
//	• [MetaID]
type ProtoMeta interface {
	activitypub.ProtoObject
	ProtoMeta() AnyMeta
}

var (
	_ ProtoMeta = AnyMeta{}
	_ ProtoMeta = Meta{}
	_ ProtoMeta = MetaID{}
)
