package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

var (
	_ activitypub.ProtoEntity = Interest{}
	_ activitypub.ProtoEntity = InterestID("")
	_ activitypub.ProtoEntity = Reference{}
	_ activitypub.ProtoEntity = ReferenceID("")
	_ activitypub.ProtoEntity = Resume{}
	_ activitypub.ProtoEntity = Skill{}
	_ activitypub.ProtoEntity = SkillID("")
)
