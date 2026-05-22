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

func (receiver *CoreResume) AppendAwardIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Awards = append(receiver.Awards, SomeAwardID(iri))
}

func (receiver *CoreResume) AppendCertificateIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Certificates = append(receiver.Certificates, SomeCertificateID(iri))
}

func (receiver *CoreResume) AppendEducationIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Education = append(receiver.Education, SomeEducationID(iri))
}

func (receiver *CoreResume) AppendInterestIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Interests = append(receiver.Interests, SomeInterestID(iri))
}

func (receiver *CoreResume) AppendLanguageIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Languages = append(receiver.Languages, SomeLanguageID(iri))
}

func (receiver *CoreResume) AppendProjectIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Projects = append(receiver.Projects, SomeProjectID(iri))
}

func (receiver *CoreResume) AppendPublicationIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Publications = append(receiver.Publications, SomePublicationID(iri))
}

func (receiver *CoreResume) AppendReferenceIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.References = append(receiver.References, SomeReferenceID(iri))
}

func (receiver *CoreResume) AppendSkillIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Skills = append(receiver.Skills, SomeSkillID(iri))
}

func (receiver *CoreResume) AppendVolunteerIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Volunteer = append(receiver.Volunteer, SomeExperienceID(iri))
}

func (receiver *CoreResume) AppendWorkIRI(iri string) {
	if nil == receiver {
		return
	}

	receiver.Work = append(receiver.Work, SomeExperienceID(iri))
}
