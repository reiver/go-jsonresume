package jsonresume

import (
	gobytes "bytes"
	gojson "encoding/json"

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

func (receiver AnyResume) String() string {
	var buffer gobytes.Buffer

	bytes, err := jsonld.Marshal(receiver)
	if nil != err {
		return "{}"
	}

	err = gojson.Indent(&buffer, bytes, "", "  ")
	if nil != err {
		return "{}"
	}

	return buffer.String()
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

//@TODO: Name


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
