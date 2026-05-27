package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

var (
	_ activitypub.ProtoObject = AnyAward{}
	_ activitypub.ProtoObject = AnyBasics{}
	_ activitypub.ProtoObject = AnyCertificate{}
	_ activitypub.ProtoObject = AnyEducation{}
	_ activitypub.ProtoObject = AnyExperience{}
	_ activitypub.ProtoObject = AnyInterest{}
	_ activitypub.ProtoObject = AnyLanguage{}
	_ activitypub.ProtoObject = AnyLocation{}
	_ activitypub.ProtoObject = AnyMeta{}
	_ activitypub.ProtoObject = AnyProfile{}
	_ activitypub.ProtoObject = AnyProject{}
	_ activitypub.ProtoObject = AnyPublication{}
	_ activitypub.ProtoObject = AnyReference{}
	_ activitypub.ProtoObject = AnyResume{}
	_ activitypub.ProtoObject = AnySkill{}
)

var (
	_ activitypub.ProtoObject = Award{}
	_ activitypub.ProtoObject = AwardID{}
	_ activitypub.ProtoObject = Interest{}
	_ activitypub.ProtoObject = InterestID{}
	_ activitypub.ProtoObject = Profile{}
	_ activitypub.ProtoObject = ProfileID{}
	_ activitypub.ProtoObject = Reference{}
	_ activitypub.ProtoObject = ReferenceID{}
	_ activitypub.ProtoObject = Resume{}
	_ activitypub.ProtoObject = ResumeID{}
	_ activitypub.ProtoObject = Skill{}
	_ activitypub.ProtoObject = SkillID{}
)
