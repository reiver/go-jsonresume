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
	_ activitypub.ProtoNode = Basics{}
	_ activitypub.ProtoNode = BasicsID{}
	_ activitypub.ProtoNode = Certificate{}
	_ activitypub.ProtoNode = CertificateID{}
	_ activitypub.ProtoNode = Education{}
	_ activitypub.ProtoNode = EducationID{}
	_ activitypub.ProtoNode = Experience{}
	_ activitypub.ProtoNode = ExperienceID{}
	_ activitypub.ProtoNode = Interest{}
	_ activitypub.ProtoNode = InterestID{}
	_ activitypub.ProtoNode = Language{}
	_ activitypub.ProtoNode = LanguageID{}
	_ activitypub.ProtoNode = Location{}
	_ activitypub.ProtoNode = LocationID{}
	_ activitypub.ProtoNode = Meta{}
	_ activitypub.ProtoNode = MetaID{}
	_ activitypub.ProtoNode = Profile{}
	_ activitypub.ProtoNode = ProfileID{}
	_ activitypub.ProtoNode = Project{}
	_ activitypub.ProtoNode = ProjectID{}
	_ activitypub.ProtoNode = Publication{}
	_ activitypub.ProtoNode = PublicationID{}
	_ activitypub.ProtoNode = Reference{}
	_ activitypub.ProtoNode = ReferenceID{}
	_ activitypub.ProtoNode = Resume{}
	_ activitypub.ProtoNode = ResumeID{}
	_ activitypub.ProtoNode = Skill{}
	_ activitypub.ProtoNode = SkillID{}
)
