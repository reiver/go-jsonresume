package jsonresume

import (
	gobytes "bytes"
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Note that you should use Resume for marshaling but not unmarshaling.
// For unmarshaling instead use [AnyResume].
//
// See also:
//
//	• [AnyResume]
//	• [CoreResume]
//	• [ProtoResume]
//	• [ResumeID]
//	• [TypeResume]
type Resume struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Resume"`

	CoreResume
}

func (receiver Resume) ProtoNode() activitypub.AnyNode {
	const _type string = TypeResume

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Resume) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeResume

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Resume) ProtoObject() activitypub.AnyObject {
	const _type string = TypeResume

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Resume) ProtoResume() AnyResume {
	const _type string = TypeResume

	return AnyResume{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreResume: receiver.CoreResume,
	}
}

func (receiver Resume) String() string {
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

//@TODO: the user should NOT unmarshal into this, but should instead unmarshal into [AnyResume]
func (receiver *Resume) UnmarshalJSON(bytes []byte) error {
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
			var typeValue string
			err := gojson.Unmarshal(bb, &typeValue)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal resume type")
				return err
			}

			switch typeValue {
			case TypeResume, CompactTypeResume, ExpandedTypeResume:
				// OK
			default:
				return erorr.Errorf("jsonresume: unexpected type for resume: %q", typeValue)
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
