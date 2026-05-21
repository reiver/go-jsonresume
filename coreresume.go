package jsonresume

type CoreResume struct {
	Awards       []ProtoAward      `json:"awards"`
//	Basics       ??? `json:"basics"`
//	Certificates ??? `json:"certificates"`
//	Education    ??? `json:"education"`
//	Interests    ??? `json:"interests"`
//	Languages    ??? `json:"languages"`
//	Projects     ??? `json:"projects"`
//	References   ??? `json:"references"`
	Skills       []ProtoSkill      `json:"skills"`
	Volunteer    []ProtoExperience `json:"volunteer"`
	Work         []ProtoExperience `json:"work"`
}
