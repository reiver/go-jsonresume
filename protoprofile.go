package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoProfile represents something that is any type 'Profile', including sub-types.
//
// See also:
//
//	• [AnyProfile]
//	• [Profile]
//	• [ProfileID]
type ProtoProfile interface {
	activitypub.ProtoObject
	ProtoProfile() AnyProfile
}

var (
	_ ProtoProfile = AnyProfile{}
	_ ProtoProfile = Profile{}
	_ ProtoProfile = ProfileID{}
)
