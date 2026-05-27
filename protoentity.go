package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

var (
	_ activitypub.ProtoEntity = AnyAward{}
	_ activitypub.ProtoEntity = AnyBasics{}
	_ activitypub.ProtoEntity = AnyCertificate{}
	_ activitypub.ProtoEntity = AnyEducation{}
	_ activitypub.ProtoEntity = AnyExperience{}
	_ activitypub.ProtoEntity = AnyInterest{}
	_ activitypub.ProtoEntity = AnyLanguage{}
	_ activitypub.ProtoEntity = AnyLocation{}
	_ activitypub.ProtoEntity = AnyMeta{}
	_ activitypub.ProtoEntity = AnyProfile{}
	_ activitypub.ProtoEntity = AnyProject{}
	_ activitypub.ProtoEntity = AnyPublication{}
	_ activitypub.ProtoEntity = AnyReference{}
	_ activitypub.ProtoEntity = AnyResume{}
	_ activitypub.ProtoEntity = AnySkill{}
)

var (
	_ activitypub.ProtoEntity = Award{}
	_ activitypub.ProtoEntity = AwardID{}
	_ activitypub.ProtoEntity = Interest{}
	_ activitypub.ProtoEntity = InterestID{}
	_ activitypub.ProtoEntity = Profile{}
	_ activitypub.ProtoEntity = ProfileID{}
	_ activitypub.ProtoEntity = Reference{}
	_ activitypub.ProtoEntity = ReferenceID{}
	_ activitypub.ProtoEntity = Resume{}
	_ activitypub.ProtoEntity = ResumeID{}
	_ activitypub.ProtoEntity = Skill{}
	_ activitypub.ProtoEntity = SkillID{}
)
