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
	_ activitypub.ProtoObject = Basics{}
	_ activitypub.ProtoObject = BasicsID{}
	_ activitypub.ProtoObject = Certificate{}
	_ activitypub.ProtoObject = CertificateID{}
	_ activitypub.ProtoObject = Education{}
	_ activitypub.ProtoObject = EducationID{}
	_ activitypub.ProtoObject = Experience{}
	_ activitypub.ProtoObject = ExperienceID{}
	_ activitypub.ProtoObject = Interest{}
	_ activitypub.ProtoObject = InterestID{}
	_ activitypub.ProtoObject = Language{}
	_ activitypub.ProtoObject = LanguageID{}
	_ activitypub.ProtoObject = Location{}
	_ activitypub.ProtoObject = LocationID{}
	_ activitypub.ProtoObject = Meta{}
	_ activitypub.ProtoObject = MetaID{}
	_ activitypub.ProtoObject = Profile{}
	_ activitypub.ProtoObject = ProfileID{}
	_ activitypub.ProtoObject = Project{}
	_ activitypub.ProtoObject = ProjectID{}
	_ activitypub.ProtoObject = Publication{}
	_ activitypub.ProtoObject = PublicationID{}
	_ activitypub.ProtoObject = Reference{}
	_ activitypub.ProtoObject = ReferenceID{}
	_ activitypub.ProtoObject = Resume{}
	_ activitypub.ProtoObject = ResumeID{}
	_ activitypub.ProtoObject = Skill{}
	_ activitypub.ProtoObject = SkillID{}
)
