package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"
)

func TestAnyResume_unmarshalJSONLD(t *testing.T) {

	cmpOpts := cmp.Options{
		cmpopts.EquateComparable(
			jsonld.ID{},
			opt.Optional[string]{},
			nul.Nullable[string]{},
			activitypub.Strings{},
			jsonresume.AwardID{},
			jsonresume.ExperienceID{},
			jsonresume.SkillID{},
			jsonld.Types{},
		),
	}

	tests := []struct {
		JSON     string
		Expected jsonresume.AnyResume
	}{
		{
			JSON: `{}`,
		},

		// minimal — just type
		{
			JSON: `{"type":"Resume"}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
			},
		},

		// two  types
		{
			JSON: `{"type":["Resume","Object"]}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeTypes("Resume", "Object"),
			},
		},

		// 3 types
		{
			JSON: `{"type":["Resume","Object","Banana"]}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeTypes("Resume", "Object", "Banana"),
			},
		},

		// with name (ignored — Resume has no Name field)
		{
			JSON: `{"type":"Resume","name":"Joe Blow"}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
			},
		},

		// with name and summary (ignored — Resume has no Name or Summary field)
		{
			JSON: `{"type":"Resume","name":"Joe Blow","summary":"CTO, Experienced Programmer"}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
			},
		},

		// with name and summary null (ignored)
		{
			JSON: `{"type":"Resume","name":"Joe Blow","summary":null}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
			},
		},

		// with id, name, summary (name and summary ignored)
		{
			JSON: `{"id":"http://example.com/resume/1","type":"Resume","name":"Joe Blow","summary":"CTO"}`,
			Expected: jsonresume.AnyResume{
				ID:   jsonld.SomeID("http://example.com/resume/1"),
				Type: jsonld.SomeType("Resume"),
			},
		},

		// with @context (should be ignored)
		{
			JSON: `{` +
				`"@context":{` +
				`"cv":"https://w3id.org/fep/6158"` +
				`,` +
				`"as":"https://www.w3.org/ns/activitystreams"` +
				`}` +
				`,` +
				`"type":"Resume"` +
				`,` +
				`"name":"Joe Blow"` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
			},
		},

		// empty collections, basics null
		{
			JSON: `{` +
				`"type":"Resume"` +
				`,` +
				`"awards":[]` +
				`,` +
				`"basics":null` +
				`,` +
				`"certificates":[]` +
				`,` +
				`"education":[]` +
				`,` +
				`"interests":[]` +
				`,` +
				`"languages":[]` +
				`,` +
				`"projects":[]` +
				`,` +
				`"publications":[]` +
				`,` +
				`"references":[]` +
				`,` +
				`"skills":[]` +
				`,` +
				`"volunteer":[]` +
				`,` +
				`"work":[]` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
			},
		},

		// full with id, name, summary, empty collections, @context
		{
			JSON: `{` +
				`"@context":{` +
				`"cv":"https://w3id.org/fep/6158"` +
				`,` +
				`"as":"https://www.w3.org/ns/activitystreams"` +
				`}` +
				`,` +
				`"id":"http://example.com/resume/executive"` +
				`,` +
				`"type":"Resume"` +
				`,` +
				`"name":"Jane Doe"` +
				`,` +
				`"summary":"Senior Software Engineer"` +
				`,` +
				`"awards":[]` +
				`,` +
				`"basics":null` +
				`,` +
				`"certificates":[]` +
				`,` +
				`"education":[]` +
				`,` +
				`"interests":[]` +
				`,` +
				`"languages":[]` +
				`,` +
				`"projects":[]` +
				`,` +
				`"publications":[]` +
				`,` +
				`"references":[]` +
				`,` +
				`"skills":[]` +
				`,` +
				`"volunteer":[]` +
				`,` +
				`"work":[]` +
				`}`,
			Expected: jsonresume.AnyResume{
				ID:   jsonld.SomeID("http://example.com/resume/executive"),
				Type: jsonld.SomeType("Resume"),
			},
		},

		// with published and content (ignored — Resume has no Published or Content field)
		{
			JSON: `{` +
				`"type":"Resume"` +
				`,` +
				`"name":"Joe Blow"` +
				`,` +
				`"summary":"CTO"` +
				`,` +
				`"published":"2024-01-15"` +
				`,` +
				`"content":"Full resume content here."` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
			},
		},

		// name only, no type
		{
			JSON: `{"name":"Alice"}`,
		},

		// partial collections — only some present
		{
			JSON: `{` +
				`"type":"Resume"` +
				`,` +
				`"name":"Bob"` +
				`,` +
				`"awards":[]` +
				`,` +
				`"skills":[]` +
				`,` +
				`"work":[]` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
			},
		},

		// empty JSON object
		{
			JSON: `{}`,
		},

		// single award as JSON object
		{
			JSON: `{` +
				`"type":"Resume"` +
				`,` +
				`"awards":[{"title":"Best Employee (2024)","date":"2024-05-21","awarder":"SuperCo","summary":"He did good work."}]` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
				CoreResume: jsonresume.CoreResume{
					Awards: []jsonresume.ProtoAward{
						jsonresume.AnyAward{
							CoreAward: jsonresume.CoreAward{
								Title:   nul.Something("Best Employee (2024)"),
								Date:    nul.Something("2024-05-21"),
								Awarder: nul.Something("SuperCo"),
								Summary: nul.Something("He did good work."),
							},
						},
					},
				},
			},
		},

		// multiple awards as JSON objects
		{
			JSON: `{` +
				`"type":"Resume"` +
				`,` +
				`"awards":[` +
				`{"title":"Best Employee (2024)","awarder":"SuperCo"}` +
				`,` +
				`{"title":"Innovation Prize","awarder":"Acme"}` +
				`]` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
				CoreResume: jsonresume.CoreResume{
					Awards: []jsonresume.ProtoAward{
						jsonresume.AnyAward{
							CoreAward: jsonresume.CoreAward{
								Title:   nul.Something("Best Employee (2024)"),
								Awarder: nul.Something("SuperCo"),
							},
						},
						jsonresume.AnyAward{
							CoreAward: jsonresume.CoreAward{
								Title:   nul.Something("Innovation Prize"),
								Awarder: nul.Something("Acme"),
							},
						},
					},
				},
			},
		},

		// award as IRI string (JSON-LD style)
		{
			JSON: `{` +
				`"type":"Resume"` +
				`,` +
				`"awards":["http://example.com/resume/award/best-employee-2024"]` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
				CoreResume: jsonresume.CoreResume{
					Awards: []jsonresume.ProtoAward{
						jsonresume.SomeAwardID("http://example.com/resume/award/best-employee-2024"),
					},
				},
			},
		},

		// single work experience
		{
			JSON: `{` +
				`"type":"Resume"` +
				`,` +
				`"work":[{"name":"SuperCo","position":"CTO","startDate":"2024-01-01","summary":"Technical leadership."}]` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
				CoreResume: jsonresume.CoreResume{
					Work: []jsonresume.ProtoExperience{
						jsonresume.AnyExperience{
							CoreExperience: jsonresume.CoreExperience{
								Name:      nul.Something("SuperCo"),
								Position:  activitypub.SomeStrings("CTO"),
								StartDate: nul.Something("2024-01-01"),
								Summary:   nul.Something("Technical leadership."),
							},
						},
					},
				},
			},
		},

		// single skill
		{
			JSON: `{` +
				`"type":"Resume"` +
				`,` +
				`"skills":[{"name":"Go","level":"advanced"}]` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
				CoreResume: jsonresume.CoreResume{
					Skills: []jsonresume.ProtoSkill{
						jsonresume.AnySkill{
							CoreSkill: jsonresume.CoreSkill{
								Name:  nul.Something("Go"),
								Level: nul.Something("advanced"),
							},
						},
					},
				},
			},
		},

		// mixed — awards, skills, and work together
		{
			JSON: `{` +
				`"type":"Resume"` +
				`,` +
				`"awards":[{"title":"Best Employee"}]` +
				`,` +
				`"skills":[{"name":"Go"}]` +
				`,` +
				`"work":[{"name":"SuperCo"}]` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
				CoreResume: jsonresume.CoreResume{
					Awards: []jsonresume.ProtoAward{
						jsonresume.AnyAward{
							CoreAward: jsonresume.CoreAward{
								Title: nul.Something("Best Employee"),
							},
						},
					},
					Skills: []jsonresume.ProtoSkill{
						jsonresume.AnySkill{
							CoreSkill: jsonresume.CoreSkill{
								Name: nul.Something("Go"),
							},
						},
					},
					Work: []jsonresume.ProtoExperience{
						jsonresume.AnyExperience{
							CoreExperience: jsonresume.CoreExperience{
								Name: nul.Something("SuperCo"),
							},
						},
					},
				},
			},
		},

		// mixed awards — one IRI string, one JSON object
		{
			JSON: `{` +
				`"type":"Resume"` +
				`,` +
				`"awards":[` +
				`"http://example.com/resume/award/innovation-prize"` +
				`,` +
				`{"title":"Best Employee","awarder":"SuperCo"}` +
				`]` +
				`}`,
			Expected: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
				CoreResume: jsonresume.CoreResume{
					Awards: []jsonresume.ProtoAward{
						jsonresume.SomeAwardID("http://example.com/resume/award/innovation-prize"),
						jsonresume.AnyAward{
							CoreAward: jsonresume.CoreAward{
								Title:   nul.Something("Best Employee"),
								Awarder: nul.Something("SuperCo"),
							},
						},
					},
				},
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.AnyResume

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		if diff := cmp.Diff(test.Expected, actual, cmpOpts); diff != "" {
			t.Errorf("For test #%d, actual result is not what was expected (-expected +actual):\n%s", testNumber, diff)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
