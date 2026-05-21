package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-nul"
)

type CoreEducation struct {
	Area        activitypub.Strings  `json:"area"`
	Courses     activitypub.Strings  `json:"courses"`
	EndDate     nul.Nullable[string] `json:"endDate"`
	Institution nul.Nullable[string] `json:"institution"`
	Score       nul.Nullable[string] `json:"score"`
	StartDate   nul.Nullable[string] `json:"startDate"`
	StudyType   activitypub.Strings  `json:"studyType"`
}
