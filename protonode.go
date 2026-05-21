package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

var (
	_ activitypub.ProtoNode = Interest{}
	_ activitypub.ProtoNode = InterestID("")
	_ activitypub.ProtoNode = Reference{}
	_ activitypub.ProtoNode = ReferenceID("")
	_ activitypub.ProtoNode = Resume{}
	_ activitypub.ProtoNode = Skill{}
	_ activitypub.ProtoNode = SkillID("")
)
