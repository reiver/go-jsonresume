package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-jsonld"
)

type AnyResume struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID    `json:"id,omitempty"`
	Type jsonld.Types `json:"type,omitempty"`

	CoreResume
}

func (receiver *AnyResume) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawResume
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal resume")
		return err
	}

	{
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.ID)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume id")
				return err
			}
		}
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.Type)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume type")
				return err
			}
		}
	}

	{
		{
			var bb []byte = []byte(raw.Awards)

			if 0 < len(bb) {
				protoAwardSlice, err := protoSliceUnmarshalJSON[ProtoAward, AwardID, AnyAward](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume awards")
					return err
				}
				receiver.Awards       = protoAwardSlice
			}
		}

		{
			var bb []byte = []byte(raw.Basics)

			if 0 < len(bb) {
				protoBasics, err := protoUnmarshalJSON[ProtoBasics, BasicsID, AnyBasics](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume basics")
					return err
				}
				receiver.Basics       = protoBasics
			}
		}

		{
			var bb []byte = []byte(raw.Certificates)

			if 0 < len(bb) {
				protoCertificateSlice, err := protoSliceUnmarshalJSON[ProtoCertificate, CertificateID, AnyCertificate](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume certificates")
					return err
				}
				receiver.Certificates = protoCertificateSlice
			}
		}

		{
			var bb []byte = []byte(raw.Education)

			if 0 < len(bb) {
				protoEducationSlice, err := protoSliceUnmarshalJSON[ProtoEducation, EducationID, AnyEducation](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume education")
					return err
				}
				receiver.Education    = protoEducationSlice
			}
		}

		{
			var bb []byte = []byte(raw.Interests)

			if 0 < len(bb) {
				protoInterestSlice, err := protoSliceUnmarshalJSON[ProtoInterest, InterestID, AnyInterest](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume interests")
					return err
				}
				receiver.Interests    = protoInterestSlice
			}
		}

		{
			var bb []byte = []byte(raw.Languages)

			if 0 < len(bb) {
				protoLanguageSlice, err := protoSliceUnmarshalJSON[ProtoLanguage, LanguageID, AnyLanguage](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume languages")
					return err
				}
				receiver.Languages    = protoLanguageSlice
			}
		}


//@TODO: Name


		{
			var bb []byte = []byte(raw.Projects)

			if 0 < len(bb) {
				protoProjectSlice, err := protoSliceUnmarshalJSON[ProtoProject, ProjectID, AnyProject](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume projects")
					return err
				}
				receiver.Projects     = protoProjectSlice
			}
		}

		{
			var bb []byte = []byte(raw.Publications)

			if 0 < len(bb) {
				protoPublicationSlice, err := protoSliceUnmarshalJSON[ProtoPublication, PublicationID, AnyPublication](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume publications")
					return err
				}
				receiver.Publications = protoPublicationSlice
			}
		}

		{
			var bb []byte = []byte(raw.References)

			if 0 < len(bb) {
				protoReferenceSlice, err := protoSliceUnmarshalJSON[ProtoReference, ReferenceID, AnyReference](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume references")
					return err
				}
				receiver.References   = protoReferenceSlice
			}
		}

		{
			var bb []byte = []byte(raw.Skills)

			if 0 < len(bb) {
				protoSkillSlice, err := protoSliceUnmarshalJSON[ProtoSkill, SkillID, AnySkill](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume skills")
					return err
				}
				receiver.Skills       = protoSkillSlice
			}
		}

		{
			var bb []byte = []byte(raw.Volunteer)

			if 0 < len(bb) {
				protoVolunteerSlice, err := protoSliceUnmarshalJSON[ProtoExperience, ExperienceID, AnyExperience](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume volunteer")
					return err
				}
				receiver.Volunteer    = protoVolunteerSlice
			}
		}

		{
			var bb []byte = []byte(raw.Work)

			if 0 < len(bb) {
				protoWorkSlice, err := protoSliceUnmarshalJSON[ProtoExperience, ExperienceID, AnyExperience](bb)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal resume work")
					return err
				}
				receiver.Work         = protoWorkSlice
			}
		}
	}

	return nil
}

func (receiver AnyResume) ProtoNode() activitypub.AnyNode {
	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyResume) ProtoEntity() activitypub.AnyEntity {
	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyResume) ProtoObject() activitypub.AnyObject {
	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: receiver.Type,
	}
}

func (receiver AnyResume) ProtoResume() AnyResume {
	return receiver
}
