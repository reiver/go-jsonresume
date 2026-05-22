package jsonresume

import (
	"github.com/reiver/go-jsonld"
)

type CoreResume struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Awards       []ProtoAward       `json:"awards"`
	Basics       ProtoBasics        `json:"basics"`
	Certificates []ProtoCertificate `json:"certificates"`
	Education    []ProtoEducation   `json:"education"`
	Interests    []ProtoInterest    `json:"interests"`
	Languages    []ProtoLanguage    `json:"languages"`
	Projects     []ProtoProject     `json:"projects"`
	Publications []ProtoPublication `json:"publications"`
	References   []ProtoReference   `json:"references"`
	Skills       []ProtoSkill       `json:"skills"`
	Volunteer    []ProtoExperience  `json:"volunteer"`
	Work         []ProtoExperience  `json:"work"`
}

func (receiver *CoreResume) AppendAwardID(id string) {
	if nil == receiver {
		return
	}

	receiver.Awards = append(receiver.Awards, SomeAwardID(id))
}

func (receiver *CoreResume) AppendCertificateID(id string) {
	if nil == receiver {
		return
	}

	receiver.Certificates = append(receiver.Certificates, SomeCertificateID(id))
}

func (receiver *CoreResume) AppendEducationID(id string) {
	if nil == receiver {
		return
	}

	receiver.Education = append(receiver.Education, SomeEducationID(id))
}

func (receiver *CoreResume) AppendInterestID(id string) {
	if nil == receiver {
		return
	}

	receiver.Interests = append(receiver.Interests, SomeInterestID(id))
}

func (receiver *CoreResume) AppendLanguageID(id string) {
	if nil == receiver {
		return
	}

	receiver.Languages = append(receiver.Languages, SomeLanguageID(id))
}

func (receiver *CoreResume) AppendProjectID(id string) {
	if nil == receiver {
		return
	}

	receiver.Projects = append(receiver.Projects, SomeProjectID(id))
}

func (receiver *CoreResume) AppendPublicationID(id string) {
	if nil == receiver {
		return
	}

	receiver.Publications = append(receiver.Publications, SomePublicationID(id))
}

func (receiver *CoreResume) AppendReferenceID(id string) {
	if nil == receiver {
		return
	}

	receiver.References = append(receiver.References, SomeReferenceID(id))
}

func (receiver *CoreResume) AppendSkillID(id string) {
	if nil == receiver {
		return
	}

	receiver.Skills = append(receiver.Skills, SomeSkillID(id))
}

func (receiver *CoreResume) AppendVolunteerID(id string) {
	if nil == receiver {
		return
	}

	receiver.Volunteer = append(receiver.Volunteer, SomeExperienceID(id))
}

func (receiver *CoreResume) AppendWorkID(id string) {
	if nil == receiver {
		return
	}

	receiver.Work = append(receiver.Work, SomeExperienceID(id))
}
