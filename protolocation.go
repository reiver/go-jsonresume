package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoLocation represents something that is any type 'Location', including sub-types.
//
// See also:
//
//	• [AnyLocation]
//	• [Location]
//	• [LocationID]
type ProtoLocation interface {
	activitypub.ProtoObject
	ProtoLocation() AnyLocation
}

var (
	_ ProtoLocation = Location{}
	_ ProtoLocation = LocationID{}
)
