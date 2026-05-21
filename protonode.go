package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

var (
	_ activitypub.ProtoNode = Award{}
	_ activitypub.ProtoNode = AwardID("")
	_ activitypub.ProtoNode = Interest{}
	_ activitypub.ProtoNode = InterestID("")
	_ activitypub.ProtoNode = Profile{}
	_ activitypub.ProtoNode = ProfileID("")
	_ activitypub.ProtoNode = Reference{}
	_ activitypub.ProtoNode = ReferenceID("")
	_ activitypub.ProtoNode = Resume{}
	_ activitypub.ProtoNode = ResumeID("")
	_ activitypub.ProtoNode = Skill{}
	_ activitypub.ProtoNode = SkillID("")
)
