package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

var (
	_ activitypub.ProtoEntity = Award{}
	_ activitypub.ProtoEntity = AwardID("")
	_ activitypub.ProtoEntity = Interest{}
	_ activitypub.ProtoEntity = InterestID("")
	_ activitypub.ProtoEntity = Profile{}
	_ activitypub.ProtoEntity = ProfileID("")
	_ activitypub.ProtoEntity = Reference{}
	_ activitypub.ProtoEntity = ReferenceID("")
	_ activitypub.ProtoEntity = Resume{}
	_ activitypub.ProtoEntity = ResumeID("")
	_ activitypub.ProtoEntity = Skill{}
	_ activitypub.ProtoEntity = SkillID("")
)
