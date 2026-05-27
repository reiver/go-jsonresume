package jsonresume_test

import (
	"fmt"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

func ExampleResume_jsonUnmarshal() {
	const jsonString string =
`{
"references": [{
    "name": "Jane Doe",
    "reference": "Reference…"
  }]
}
`
	var resume jsonresume.Resume

	err := json.Unmarshal([]byte(jsonString), &resume)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Println(resume)

	// Output:
	// {
	//   "@context": {
	//     "cv": "https://w3id.org/fep/6158",
	//     "awards": "cv:awards",
	//     "basics": "cv:basics",
	//     "certificates": "cv:certificates",
	//     "education": "cv:education",
	//     "id": "cv:id",
	//     "interests": "cv:interests",
	//     "languages": "cv:languages",
	//     "meta": "cv:meta",
	//     "projects": "cv:projects",
	//     "publications": "cv:publications",
	//     "references": "cv:references",
	//     "skills": "cv:skills",
	//     "type": "cv:type",
	//     "volunteer": "cv:volunteer",
	//     "work": "cv:work"
	//   },
	//   "type": "Resume",
	//   "references": [
	//     {
	//       "name": "Jane Doe",
	//       "reference": "Reference…"
	//     }
	//   ]
	// }
}

func ExampleResume_jsonldMarshal() {

	var cv jsonresume.Resume

	cv.Basics = jsonresume.Basics{
		CoreBasics: jsonresume.CoreBasics{
			EMail: activitypub.SomeString("joeblow@example.com"),
			Label: activitypub.SomeString("Programmer"),
			Name:  nul.Something("Joe Blow"),
			Phone: activitypub.SomeString("(604) 555-1234"),
			Profiles: []jsonresume.ProtoProfile{
				jsonresume.Profile{
					CoreProfile: jsonresume.CoreProfile{
						Network:  nul.Something("Mastodon"),
						UserName: nul.Something("joeblow"),
					},
				},
			},
			Summary: nul.Something("CTO, Experienced Programmer"),
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

	bytes, err := jsonld.Marshal(cv)
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(bytes))

	// Output:
	// {"@context":{"cv":"https://w3id.org/fep/6158","as":"http://www.w3.org/ns/activitystreams","awards":"cv:awards","basics":"cv:basics","certificates":"cv:certificates","education":"cv:education","email":"cv:email","id":"cv:id","interests":"cv:interests","label":"cv:label","languages":"cv:languages","location":"cv:location","meta":"cv:meta","phone":"cv:phone","profiles":"cv:profiles","projects":"cv:projects","publications":"cv:publications","references":"cv:references","skills":"cv:skills","type":"cv:type","volunteer":"cv:volunteer","work":"cv:work","image":"as:image","name":"as:name","summary":"as:summary","url":"as:url"},"type":"Resume","awards":[{"type":"Award","awarder":"SuperCo","date":"2024-05-21","title":"Best Employee (2024)"}],"basics":{"type":"Basics","email":"joeblow@example.com","label":"Programmer","name":"Joe Blow","phone":"(604) 555-1234","profiles":[{"type":"Profile","network":"Mastodon","username":"joeblow"}],"summary":"CTO, Experienced Programmer"},"languages":[{"type":"Language","fluency":"Fluent","language":"English"}],"skills":[{"type":"Skill","keywords":["Golang","PHP","HTTP"],"level":"Senior"}],"work":[{"type":"Experience","highlights":["Hired initial team.","Created architecture."],"organization":"SuperCo","position":"Chief Technology Officer (CTO)","startDate":"2024-01-01"}]}
}

func ExampleResume_withIDs() {

	var cv jsonresume.Resume

	cv.Basics = jsonresume.SomeBasicsID("http://example.com/resume/basics")

	cv.AppendAwardID("http://example.com/resume/award/best-employee-2024")
	cv.AppendAwardID("http://example.com/resume/award/acme-excellence-2021")

	cv.AppendWorkID("http://example.com/resume/experience/3")
	cv.AppendWorkID("http://example.com/resume/experience/2")

	cv.AppendVolunteerID("http://example.com/resume/experience/4")

	cv.AppendSkillID("http://example.com/resume/skill/backend-development")

	cv.AppendLanguageID("http://example.com/resume/language/english")

	bytes, err := jsonld.Marshal(cv)
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(bytes))

	// Output:
	// {"@context":{"cv":"https://w3id.org/fep/6158","awards":"cv:awards","basics":"cv:basics","certificates":"cv:certificates","education":"cv:education","id":"cv:id","interests":"cv:interests","languages":"cv:languages","meta":"cv:meta","projects":"cv:projects","publications":"cv:publications","references":"cv:references","skills":"cv:skills","type":"cv:type","volunteer":"cv:volunteer","work":"cv:work"},"type":"Resume","awards":["http://example.com/resume/award/best-employee-2024","http://example.com/resume/award/acme-excellence-2021"],"basics":"http://example.com/resume/basics","languages":["http://example.com/resume/language/english"],"skills":["http://example.com/resume/skill/backend-development"],"volunteer":["http://example.com/resume/experience/4"],"work":["http://example.com/resume/experience/3","http://example.com/resume/experience/2"]}
}
