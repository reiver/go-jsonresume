package jsonresume

import (
	"codeberg.org/reiver/go-erorr"
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
	Meta         ProtoMeta          `json:"meta,omitempty"`
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

// unmarshalRawResume unmarshals the ID and all CoreResume fields from a rawResume.
//
// This is the shared logic used by both [Resume.UnmarshalJSON] and [AnyResume.UnmarshalJSON].
// The type field is NOT handled here — each caller handles it differently:
//
//	• [Resume.UnmarshalJSON] validates the type against a fixed set of accepted values.
//	• [AnyResume.UnmarshalJSON] stores whatever type value it receives.
func (receiver *CoreResume) unmarshalRawResume(raw rawResume, id *jsonld.ID) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, id)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Awards)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoAward, AwardID, AnyAward](bb, &receiver.Awards)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume awards")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Basics)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObject[ProtoBasics, BasicsID, AnyBasics](bb, &receiver.Basics)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume basics")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Certificates)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoCertificate, CertificateID, AnyCertificate](bb, &receiver.Certificates)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume certificates")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Education)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoEducation, EducationID, AnyEducation](bb, &receiver.Education)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume education")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Interests)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoInterest, InterestID, AnyInterest](bb, &receiver.Interests)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume interests")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Languages)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoLanguage, LanguageID, AnyLanguage](bb, &receiver.Languages)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume languages")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Meta)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObject[ProtoMeta, MetaID, AnyMeta](bb, &receiver.Meta)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume meta")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Projects)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoProject, ProjectID, AnyProject](bb, &receiver.Projects)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume projects")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Publications)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoPublication, PublicationID, AnyPublication](bb, &receiver.Publications)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume publications")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.References)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoReference, ReferenceID, AnyReference](bb, &receiver.References)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume references")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Skills)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoSkill, SkillID, AnySkill](bb, &receiver.Skills)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume skills")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Volunteer)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoExperience, ExperienceID, AnyExperience](bb, &receiver.Volunteer)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume volunteer")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Work)

		if 0 < len(bb) {
			err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[ProtoExperience, ExperienceID, AnyExperience](bb, &receiver.Work)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume work")
				return err
			}
		}
	}

	return nil
}
