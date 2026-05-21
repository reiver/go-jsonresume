package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

var (
	_ activitypub.ProtoObject = Award{}
	_ activitypub.ProtoObject = AwardID("")
	_ activitypub.ProtoObject = Interest{}
	_ activitypub.ProtoObject = InterestID("")
	_ activitypub.ProtoObject = Profile{}
	_ activitypub.ProtoObject = ProfileID("")
	_ activitypub.ProtoObject = Reference{}
	_ activitypub.ProtoObject = ReferenceID("")
	_ activitypub.ProtoObject = Resume{}
	_ activitypub.ProtoObject = ResumeID("")
	_ activitypub.ProtoObject = Skill{}
	_ activitypub.ProtoObject = SkillID("")
)
