package jsonresume

type CoreResume struct {
	Awards       []ProtoAward      `json:"awards"`
//	Basics       ??? `json:"basics"`
//	Certificates ??? `json:"certificates"`
	Education    []ProtoEducation  `json:"education"`
	Interests    []ProtoInterest   `json:"interests"`
	Languages    []ProtoLanguage   `json:"languages"`
	Projects     []ProtoProject    `json:"projects"`
	References   []ProtoReference  `json:"references"`
	Skills       []ProtoSkill      `json:"skills"`
	Volunteer    []ProtoExperience `json:"volunteer"`
	Work         []ProtoExperience `json:"work"`
}
