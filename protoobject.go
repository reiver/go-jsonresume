package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

var (
	_ activitypub.ProtoObject = Interest{}
	_ activitypub.ProtoObject = InterestID("")
	_ activitypub.ProtoObject = Reference{}
	_ activitypub.ProtoObject = ReferenceID("")
	_ activitypub.ProtoObject = Resume{}
	_ activitypub.ProtoObject = Skill{}
	_ activitypub.ProtoObject = SkillID("")
)
