package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoResume represents something that is any type 'Resume', including sub-types.
//
// See also:
//
//	• [AnyResume]
//	• [Resume]
//	• [ResumeID]
type ProtoResume interface {
	activitypub.ProtoObject
	ProtoResume() AnyResume
}

var (
	_ ProtoResume = AnyResume{}
	_ ProtoResume = Resume{}
	_ ProtoResume = ResumeID{}
)
