package jsonresume

import (
	"github.com/reiver/go-jsonld"
)

type CoreResume struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	Awards       []ProtoAward       `json:"awards,omitempty"`
	Basics       ProtoBasics        `json:"basics,omitempty"`
	Certificates []ProtoCertificate `json:"certificates,omitempty"`
	Education    []ProtoEducation   `json:"education,omitempty"`
	Interests    []ProtoInterest    `json:"interests,omitempty"`
	Languages    []ProtoLanguage    `json:"languages,omitempty"`
	Projects     []ProtoProject     `json:"projects,omitempty"`
	Publications []ProtoPublication `json:"publications,omitempty"`
	References   []ProtoReference   `json:"references,omitempty"`
	Skills       []ProtoSkill       `json:"skills,omitempty"`
	Volunteer    []ProtoExperience  `json:"volunteer,omitempty"`
	Work         []ProtoExperience  `json:"work,omitempty"`
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

func (receiver *CoreResume) SetAwardID(id string) {
	if nil == receiver {
		return
	}

	receiver.Awards = []ProtoAward{SomeAwardID(id)}
}

func (receiver *CoreResume) SetCertificateID(id string) {
	if nil == receiver {
		return
	}

	receiver.Certificates = []ProtoCertificate{SomeCertificateID(id)}
}

func (receiver *CoreResume) SetEducationID(id string) {
	if nil == receiver {
		return
	}

	receiver.Education = []ProtoEducation{SomeEducationID(id)}
}

func (receiver *CoreResume) SetInterestID(id string) {
	if nil == receiver {
		return
	}

	receiver.Interests = []ProtoInterest{SomeInterestID(id)}
}

func (receiver *CoreResume) SetLanguageID(id string) {
	if nil == receiver {
		return
	}

	receiver.Languages = []ProtoLanguage{SomeLanguageID(id)}
}

func (receiver *CoreResume) SetProjectID(id string) {
	if nil == receiver {
		return
	}

	receiver.Projects = []ProtoProject{SomeProjectID(id)}
}

func (receiver *CoreResume) SetPublicationID(id string) {
	if nil == receiver {
		return
	}

	receiver.Publications = []ProtoPublication{SomePublicationID(id)}
}

func (receiver *CoreResume) SetReferenceID(id string) {
	if nil == receiver {
		return
	}

	receiver.References = []ProtoReference{SomeReferenceID(id)}
}

func (receiver *CoreResume) SetSkillID(id string) {
	if nil == receiver {
		return
	}

	receiver.Skills = []ProtoSkill{SomeSkillID(id)}
}

func (receiver *CoreResume) SetVolunteerID(id string) {
	if nil == receiver {
		return
	}

	receiver.Volunteer = []ProtoExperience{SomeExperienceID(id)}
}

func (receiver *CoreResume) SetWorkID(id string) {
	if nil == receiver {
		return
	}

	receiver.Work = []ProtoExperience{SomeExperienceID(id)}
}
