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
