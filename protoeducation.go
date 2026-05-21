package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoEducation represents something that is any type 'Education', including sub-types.
//
// See also:
//
//	• [AnyEducation]
//	• [Education]
//	• [EducationID]
type ProtoEducation interface {
	activitypub.ProtoObject
	ProtoEducation() AnyEducation
}

var (
	_ ProtoEducation = Education{}
	_ ProtoEducation = EducationID("")
)
