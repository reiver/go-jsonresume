package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoSkill represents something that is any type 'Skill', including sub-types.
//
// See also:
//
//	• [AnySkill]
//	• [Skill]
//	• [SkillID]
type ProtoSkill interface {
	activitypub.ProtoObject
	ProtoSkill() AnySkill
}

var (
	_ ProtoSkill = Skill{}
	_ ProtoSkill = SkillID{}
)
