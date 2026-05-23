package jsonresume_test

import (
	"fmt"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-jsonresume"
)

func ExampleResume_jsonMarshal() {

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

	bytes, err := jsonld.Marshal(cv)
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(bytes))

	// Output:
	// {"@context":{"cv":"https://w3id.org/fep/6158","as":"https://www.w3.org/ns/activitystreams","alsoKnownAs":"cv:alsoKnownAs","attachment":"cv:attachment","attributedTo":"cv:attributedTo","audience":"cv:audience","awards":"cv:awards","basics":"cv:basics","cc":"cv:cc","certificates":"cv:certificates","content":"cv:content","contentMap":"cv:contentMap","duration":"cv:duration","education":"cv:education","endTime":"cv:endTime","generator":"cv:generator","icon":"cv:icon","id":"cv:id","image":"cv:image","inReplyTo":"cv:inReplyTo","interests":"cv:interests","languages":"cv:languages","likes":"cv:likes","location":"cv:location","mediaType":"cv:mediaType","movedTo":"cv:movedTo","name":"cv:name","nameMap":"cv:nameMap","preview":"cv:preview","projects":"cv:projects","publications":"cv:publications","published":"cv:published","references":"cv:references","replies":"cv:replies","shares":"cv:shares","skills":"cv:skills","startTime":"cv:startTime","summary":"cv:summary","summaryMap":"cv:summaryMap","tag":"cv:tag","to":"cv:to","type":"cv:type","updated":"cv:updated","url":"cv:url","volunteer":"cv:volunteer","work":"cv:work","alsoKnownAs":"as:alsoKnownAs","attachment":"as:attachment","attributedTo":"as:attributedTo","audience":"as:audience","cc":"as:cc","content":"as:content","contentMap":"as:contentMap","duration":"as:duration","endTime":"as:endTime","generator":"as:generator","icon":"as:icon","image":"as:image","inReplyTo":"as:inReplyTo","likes":"as:likes","location":"as:location","mediaType":"as:mediaType","movedTo":"as:movedTo","name":"as:name","nameMap":"as:nameMap","preview":"as:preview","published":"as:published","replies":"as:replies","shares":"as:shares","startTime":"as:startTime","summary":"as:summary","summaryMap":"as:summaryMap","tag":"as:tag","to":"as:to","updated":"as:updated","url":"as:url"},"type":"Resume","name":"Joe Blow","summary":"CTO, Experienced Programmer","awards":[{"type":"Award","awarder":"SuperCo","date":"2024-05-21","title":"Best Employee (2024)"}],"basics":{"type":"Basics","email":"joeblow@example.com","label":"Programmer","phone":"(604) 555-1234","profiles":[{"type":"Profile","network":"Mastodon","username":"joeblow"}]},"certificates":[],"education":[],"interests":[],"languages":[{"type":"Language","fluency":"Fluent","language":"English"}],"projects":[],"publications":[],"references":[],"skills":[{"type":"Skill","keywords":["Golang","PHP","HTTP"],"level":"Senior"}],"volunteer":[],"work":[{"type":"Experience","endDate":null,"highlights":["Hired initial team.","Created architecture."],"organization":"SuperCo","position":"Chief Technology Officer (CTO)","startDate":"2024-01-01"}]}
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
	// {"@context":{"cv":"https://w3id.org/fep/6158","as":"https://www.w3.org/ns/activitystreams","alsoKnownAs":"cv:alsoKnownAs","attachment":"cv:attachment","attributedTo":"cv:attributedTo","audience":"cv:audience","awards":"cv:awards","basics":"cv:basics","cc":"cv:cc","certificates":"cv:certificates","content":"cv:content","contentMap":"cv:contentMap","duration":"cv:duration","education":"cv:education","endTime":"cv:endTime","generator":"cv:generator","icon":"cv:icon","id":"cv:id","image":"cv:image","inReplyTo":"cv:inReplyTo","interests":"cv:interests","languages":"cv:languages","likes":"cv:likes","location":"cv:location","mediaType":"cv:mediaType","movedTo":"cv:movedTo","name":"cv:name","nameMap":"cv:nameMap","preview":"cv:preview","projects":"cv:projects","publications":"cv:publications","published":"cv:published","references":"cv:references","replies":"cv:replies","shares":"cv:shares","skills":"cv:skills","startTime":"cv:startTime","summary":"cv:summary","summaryMap":"cv:summaryMap","tag":"cv:tag","to":"cv:to","type":"cv:type","updated":"cv:updated","url":"cv:url","volunteer":"cv:volunteer","work":"cv:work","alsoKnownAs":"as:alsoKnownAs","attachment":"as:attachment","attributedTo":"as:attributedTo","audience":"as:audience","cc":"as:cc","content":"as:content","contentMap":"as:contentMap","duration":"as:duration","endTime":"as:endTime","generator":"as:generator","icon":"as:icon","image":"as:image","inReplyTo":"as:inReplyTo","likes":"as:likes","location":"as:location","mediaType":"as:mediaType","movedTo":"as:movedTo","name":"as:name","nameMap":"as:nameMap","preview":"as:preview","published":"as:published","replies":"as:replies","shares":"as:shares","startTime":"as:startTime","summary":"as:summary","summaryMap":"as:summaryMap","tag":"as:tag","to":"as:to","updated":"as:updated","url":"as:url"},"type":"Resume","awards":["http://example.com/resume/award/best-employee-2024","http://example.com/resume/award/acme-excellence-2021"],"basics":"http://example.com/resume/basics","certificates":[],"education":[],"interests":[],"languages":["http://example.com/resume/language/english"],"projects":[],"publications":[],"references":[],"skills":["http://example.com/resume/skill/backend-development"],"volunteer":["http://example.com/resume/experience/4"],"work":["http://example.com/resume/experience/3","http://example.com/resume/experience/2"]}
}
