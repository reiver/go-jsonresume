package jsonresume

type CoreResume struct {
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
