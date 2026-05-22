package jsonresume_test

import (
	"fmt"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-jsonresume"
)

func ExampleResume() {

	var cv jsonresume.Resume

	cv.Name = opt.Something("Joe Blow")
	cv.Summary = nul.Something("CTO, Experienced Programmer")

	cv.Basics = jsonresume.Basics{
		CoreBasics: jsonresume.CoreBasics{
			EMail: activitypub.SomeString("joeblow@example.com"),
			Label: activitypub.SomeString("Programmer"),
			Phone: activitypub.SomeString("(604) 555-1234"),
			Profiles: []jsonresume.ProtoProfile{
				jsonresume.Profile{
					CoreProfile: jsonresume.CoreProfile{
						Network:  nul.Something("Mastodon"),
						UserName: nul.Something("joeblow"),
					},
				},
			},
		},
	}

	cv.Awards = append(cv.Awards, jsonresume.Award{
		CoreAward: jsonresume.CoreAward{
			Title:   nul.Something("Best Employee (2024)"),
			Date:    nul.Something("2024-05-21"),
			Awarder: nul.Something("SuperCo"),
		},
	})

	cv.Work = append(cv.Work, jsonresume.Experience{
		CoreExperience: jsonresume.CoreExperience{
			Organization: nul.Something("SuperCo"),
			Position:     activitypub.SomeString("Chief Technology Officer (CTO)"),
			StartDate:    nul.Something("2024-01-01"),
			EndDate:      nul.Null[string](),
			Highlights: activitypub.SomeStrings(
				"Hired initial team.",
				"Created architecture.",
			),
		},
	})

	cv.Skills = append(cv.Skills, jsonresume.Skill{
		CoreSkill: jsonresume.CoreSkill{
			Keywords: activitypub.SomeStrings("Golang", "PHP", "HTTP"),
			Level:    nul.Something("Senior"),
		},
	})

	cv.Languages = append(cv.Languages, jsonresume.Language{
		CoreLanguage: jsonresume.CoreLanguage{
			Language: nul.Something("English"),
			Fluency:  nul.Something("Fluent"),
		},
	})

	bytes, err := json.Marshal(cv)
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(bytes))

	// Output:
	// {"type":"Resume","name":"Joe Blow","summary":"CTO, Experienced Programmer","awards":[{"type":"Award","awarder":"SuperCo","date":"2024-05-21","title":"Best Employee (2024)"}],"basics":{"type":"Basics","email":"joeblow@example.com","label":"Programmer","phone":"(604) 555-1234","profiles":[{"type":"Profile","network":"Mastodon","username":"joeblow"}]},"certificates":[],"education":[],"interests":[],"languages":[{"type":"Language","fluency":"Fluent","language":"English"}],"projects":[],"publications":[],"references":[],"skills":[{"type":"Skill","keywords":["Golang","PHP","HTTP"],"level":"Senior"}],"volunteer":[],"work":[{"type":"Experience","endDate":null,"highlights":["Hired initial team.","Created architecture."],"organization":"SuperCo","position":"Chief Technology Officer (CTO)","startDate":"2024-01-01"}]}
}

func ExampleResume_withIDs() {

	var cv jsonresume.Resume

	cv.Basics = jsonresume.SomeBasicsID("http://example.com/resume/basics")

	cv.Awards = append(cv.Awards, jsonresume.SomeAwardID("http://example.com/resume/award/best-employee-2024"))
	cv.Awards = append(cv.Awards, jsonresume.SomeAwardID("http://example.com/resume/award/acme-excellence-2021"))

	cv.Work = append(cv.Work, jsonresume.SomeExperienceID("http://example.com/resume/experience/3"))
	cv.Work = append(cv.Work, jsonresume.SomeExperienceID("http://example.com/resume/experience/2"))

	cv.Volunteer = append(cv.Volunteer, jsonresume.SomeExperienceID("http://example.com/resume/experience/4"))

	cv.Skills = append(cv.Skills, jsonresume.SomeSkillID("http://example.com/resume/skill/backend-development"))

	cv.Languages = append(cv.Languages, jsonresume.SomeLanguageID("http://example.com/resume/language/english"))

	bytes, err := json.Marshal(cv)
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(bytes))

	// Output:
	// {"type":"Resume","awards":["http://example.com/resume/award/best-employee-2024","http://example.com/resume/award/acme-excellence-2021"],"basics":"http://example.com/resume/basics","certificates":[],"education":[],"interests":[],"languages":["http://example.com/resume/language/english"],"projects":[],"publications":[],"references":[],"skills":["http://example.com/resume/skill/backend-development"],"volunteer":["http://example.com/resume/experience/4"],"work":["http://example.com/resume/experience/3","http://example.com/resume/experience/2"]}
}
