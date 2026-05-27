package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

var (
	_ activitypub.ProtoNode = AnyAward{}
	_ activitypub.ProtoNode = AnyBasics{}
	_ activitypub.ProtoNode = AnyCertificate{}
	_ activitypub.ProtoNode = AnyEducation{}
	_ activitypub.ProtoNode = AnyExperience{}
	_ activitypub.ProtoNode = AnyInterest{}
	_ activitypub.ProtoNode = AnyLanguage{}
	_ activitypub.ProtoNode = AnyLocation{}
	_ activitypub.ProtoNode = AnyMeta{}
	_ activitypub.ProtoNode = AnyProfile{}
	_ activitypub.ProtoNode = AnyProject{}
	_ activitypub.ProtoNode = AnyPublication{}
	_ activitypub.ProtoNode = AnyReference{}
	_ activitypub.ProtoNode = AnyResume{}
	_ activitypub.ProtoNode = AnySkill{}
)

var (
	_ activitypub.ProtoNode = Award{}
	_ activitypub.ProtoNode = AwardID{}
	_ activitypub.ProtoNode = Interest{}
	_ activitypub.ProtoNode = InterestID{}
	_ activitypub.ProtoNode = Profile{}
	_ activitypub.ProtoNode = ProfileID{}
	_ activitypub.ProtoNode = Reference{}
	_ activitypub.ProtoNode = ReferenceID{}
	_ activitypub.ProtoNode = Resume{}
	_ activitypub.ProtoNode = ResumeID{}
	_ activitypub.ProtoNode = Skill{}
	_ activitypub.ProtoNode = SkillID{}
)
