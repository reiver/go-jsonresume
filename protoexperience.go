package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoExperience represents something that is any type 'Experience', including sub-types.
//
// See also:
//
//	• [AnyExperience]
//	• [Experience]
//	• [ExperienceID]
type ProtoExperience interface {
	activitypub.ProtoObject
	ProtoExperience() AnyExperience
}

var (
	_ ProtoExperience = AnyExperience{}
	_ ProtoExperience = Experience{}
	_ ProtoExperience = ExperienceID{}
)
