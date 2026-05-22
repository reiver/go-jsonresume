package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"
)

func TestResume_unmarshalJSONLD(t *testing.T) {

	tests := []struct {
		JSON string

		ExpectedID        jsonld.ID
		ExpectedName      opt.Optional[string]
		ExpectedSummary   nul.Nullable[string]
		ExpectedPublished opt.Optional[string]
		ExpectedContent   opt.Optional[string]

		ExpectedAwardsLen       int
		ExpectedCertificatesLen int
		ExpectedEducationLen    int
		ExpectedInterestsLen    int
		ExpectedLanguagesLen    int
		ExpectedProjectsLen     int
		ExpectedPublicationsLen int
		ExpectedReferencesLen   int
		ExpectedSkillsLen       int
		ExpectedVolunteerLen    int
		ExpectedWorkLen         int

		ExpectedBasicsNil bool

		ExpectedAwardsNil       bool
		ExpectedCertificatesNil bool
		ExpectedEducationNil    bool
		ExpectedInterestsNil    bool
		ExpectedLanguagesNil    bool
		ExpectedProjectsNil     bool
		ExpectedPublicationsNil bool
		ExpectedReferencesNil   bool
		ExpectedSkillsNil       bool
		ExpectedVolunteerNil    bool
		ExpectedWorkNil         bool
	}{
		// 0: minimal — just type
		{
			JSON: `{"type":"Resume"}`,

			ExpectedBasicsNil: true,

			ExpectedAwardsNil:       true,
			ExpectedCertificatesNil: true,
			ExpectedEducationNil:    true,
			ExpectedInterestsNil:    true,
			ExpectedLanguagesNil:    true,
			ExpectedProjectsNil:     true,
			ExpectedPublicationsNil: true,
			ExpectedReferencesNil:   true,
			ExpectedSkillsNil:       true,
			ExpectedVolunteerNil:    true,
			ExpectedWorkNil:         true,
		},

		// 1: with name
		{
			JSON: `{"type":"Resume","name":"Joe Blow"}`,

			ExpectedName: opt.Something("Joe Blow"),

			ExpectedBasicsNil: true,

			ExpectedAwardsNil:       true,
			ExpectedCertificatesNil: true,
			ExpectedEducationNil:    true,
			ExpectedInterestsNil:    true,
			ExpectedLanguagesNil:    true,
			ExpectedProjectsNil:     true,
			ExpectedPublicationsNil: true,
			ExpectedReferencesNil:   true,
			ExpectedSkillsNil:       true,
			ExpectedVolunteerNil:    true,
			ExpectedWorkNil:         true,
		},

		// 2: with name and summary
		{
			JSON: `{"type":"Resume","name":"Joe Blow","summary":"CTO, Experienced Programmer"}`,

			ExpectedName:    opt.Something("Joe Blow"),
			ExpectedSummary: nul.Something("CTO, Experienced Programmer"),

			ExpectedBasicsNil: true,

			ExpectedAwardsNil:       true,
			ExpectedCertificatesNil: true,
			ExpectedEducationNil:    true,
			ExpectedInterestsNil:    true,
			ExpectedLanguagesNil:    true,
			ExpectedProjectsNil:     true,
			ExpectedPublicationsNil: true,
			ExpectedReferencesNil:   true,
			ExpectedSkillsNil:       true,
			ExpectedVolunteerNil:    true,
			ExpectedWorkNil:         true,
		},

		// 3: with name and summary null
		{
			JSON: `{"type":"Resume","name":"Joe Blow","summary":null}`,

			ExpectedName:    opt.Something("Joe Blow"),
			ExpectedSummary: nul.Null[string](),

			ExpectedBasicsNil: true,

			ExpectedAwardsNil:       true,
			ExpectedCertificatesNil: true,
			ExpectedEducationNil:    true,
			ExpectedInterestsNil:    true,
			ExpectedLanguagesNil:    true,
			ExpectedProjectsNil:     true,
			ExpectedPublicationsNil: true,
			ExpectedReferencesNil:   true,
			ExpectedSkillsNil:       true,
			ExpectedVolunteerNil:    true,
			ExpectedWorkNil:         true,
		},

		// 4: with id, name, summary
		{
			JSON: `{"id":"http://example.com/resume/1","type":"Resume","name":"Joe Blow","summary":"CTO"}`,

			ExpectedID:      jsonld.SomeID("http://example.com/resume/1"),
			ExpectedName:    opt.Something("Joe Blow"),
			ExpectedSummary: nul.Something("CTO"),

			ExpectedBasicsNil: true,

			ExpectedAwardsNil:       true,
			ExpectedCertificatesNil: true,
			ExpectedEducationNil:    true,
			ExpectedInterestsNil:    true,
			ExpectedLanguagesNil:    true,
			ExpectedProjectsNil:     true,
			ExpectedPublicationsNil: true,
			ExpectedReferencesNil:   true,
			ExpectedSkillsNil:       true,
			ExpectedVolunteerNil:    true,
			ExpectedWorkNil:         true,
		},

		// 5: with @context (should be ignored)
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

			ExpectedName: opt.Something("Joe Blow"),

			ExpectedBasicsNil: true,

			ExpectedAwardsNil:       true,
			ExpectedCertificatesNil: true,
			ExpectedEducationNil:    true,
			ExpectedInterestsNil:    true,
			ExpectedLanguagesNil:    true,
			ExpectedProjectsNil:     true,
			ExpectedPublicationsNil: true,
			ExpectedReferencesNil:   true,
			ExpectedSkillsNil:       true,
			ExpectedVolunteerNil:    true,
			ExpectedWorkNil:         true,
		},

		// 6: empty collections, basics null
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

			ExpectedBasicsNil: true,
		},

		// 7: full with id, name, summary, empty collections, @context
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

			ExpectedID:      jsonld.SomeID("http://example.com/resume/executive"),
			ExpectedName:    opt.Something("Jane Doe"),
			ExpectedSummary: nul.Something("Senior Software Engineer"),

			ExpectedBasicsNil: true,
		},

		// 8: with published and content (CoreObject fields)
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

			ExpectedName:      opt.Something("Joe Blow"),
			ExpectedSummary:   nul.Something("CTO"),
			ExpectedPublished: opt.Something("2024-01-15"),
			ExpectedContent:   opt.Something("Full resume content here."),

			ExpectedBasicsNil: true,

			ExpectedAwardsNil:       true,
			ExpectedCertificatesNil: true,
			ExpectedEducationNil:    true,
			ExpectedInterestsNil:    true,
			ExpectedLanguagesNil:    true,
			ExpectedProjectsNil:     true,
			ExpectedPublicationsNil: true,
			ExpectedReferencesNil:   true,
			ExpectedSkillsNil:       true,
			ExpectedVolunteerNil:    true,
			ExpectedWorkNil:         true,
		},

		// 9: name only, no type (type is Const so always set)
		{
			JSON: `{"name":"Alice"}`,

			ExpectedName: opt.Something("Alice"),

			ExpectedBasicsNil: true,

			ExpectedAwardsNil:       true,
			ExpectedCertificatesNil: true,
			ExpectedEducationNil:    true,
			ExpectedInterestsNil:    true,
			ExpectedLanguagesNil:    true,
			ExpectedProjectsNil:     true,
			ExpectedPublicationsNil: true,
			ExpectedReferencesNil:   true,
			ExpectedSkillsNil:       true,
			ExpectedVolunteerNil:    true,
			ExpectedWorkNil:         true,
		},

		// 10: partial collections — only some present
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

			ExpectedName: opt.Something("Bob"),

			ExpectedBasicsNil: true,

			ExpectedCertificatesNil: true,
			ExpectedEducationNil:    true,
			ExpectedInterestsNil:    true,
			ExpectedLanguagesNil:    true,
			ExpectedProjectsNil:     true,
			ExpectedPublicationsNil: true,
			ExpectedReferencesNil:   true,
			ExpectedVolunteerNil:    true,
		},

		// 11: empty JSON object
		{
			JSON: `{}`,

			ExpectedBasicsNil: true,

			ExpectedAwardsNil:       true,
			ExpectedCertificatesNil: true,
			ExpectedEducationNil:    true,
			ExpectedInterestsNil:    true,
			ExpectedLanguagesNil:    true,
			ExpectedProjectsNil:     true,
			ExpectedPublicationsNil: true,
			ExpectedReferencesNil:   true,
			ExpectedSkillsNil:       true,
			ExpectedVolunteerNil:    true,
			ExpectedWorkNil:         true,
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.Resume

		err := json.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		{
			expected := test.ExpectedID
			actual := actual.ID
			if expected != actual {
				t.Errorf("For test #%d, ID is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		{
			expected := test.ExpectedName
			actual := actual.Name
			if expected != actual {
				t.Errorf("For test #%d, Name is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		{
			expected := test.ExpectedSummary
			actual := actual.Summary
			if expected != actual {
				t.Errorf("For test #%d, Summary is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		{
			expected := test.ExpectedPublished
			actual := actual.Published
			if expected != actual {
				t.Errorf("For test #%d, Published is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		{
			expected := test.ExpectedContent
			actual := actual.Content
			if expected != actual {
				t.Errorf("For test #%d, Content is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		{
			expected := test.ExpectedBasicsNil
			actual := actual.Basics == nil
			if expected != actual {
				t.Errorf("For test #%d, Basics nil is not what was expected.", testNumber)
				t.Logf("EXPECTED nil: %v", expected)
				t.Logf("ACTUAL nil:   %v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		// Check collection lengths and nil-ness.
		{
			type collectionCheck struct {
				Name        string
				ActualLen   int
				ActualNil   bool
				ExpectedLen int
				ExpectedNil bool
			}

			checks := []collectionCheck{
				{"Awards",       len(actual.Awards),       actual.Awards == nil,       test.ExpectedAwardsLen,       test.ExpectedAwardsNil},
				{"Certificates", len(actual.Certificates), actual.Certificates == nil, test.ExpectedCertificatesLen, test.ExpectedCertificatesNil},
				{"Education",    len(actual.Education),    actual.Education == nil,    test.ExpectedEducationLen,    test.ExpectedEducationNil},
				{"Interests",    len(actual.Interests),    actual.Interests == nil,    test.ExpectedInterestsLen,    test.ExpectedInterestsNil},
				{"Languages",    len(actual.Languages),    actual.Languages == nil,    test.ExpectedLanguagesLen,    test.ExpectedLanguagesNil},
				{"Projects",     len(actual.Projects),     actual.Projects == nil,     test.ExpectedProjectsLen,     test.ExpectedProjectsNil},
				{"Publications", len(actual.Publications), actual.Publications == nil, test.ExpectedPublicationsLen, test.ExpectedPublicationsNil},
				{"References",   len(actual.References),   actual.References == nil,   test.ExpectedReferencesLen,   test.ExpectedReferencesNil},
				{"Skills",       len(actual.Skills),       actual.Skills == nil,       test.ExpectedSkillsLen,       test.ExpectedSkillsNil},
				{"Volunteer",    len(actual.Volunteer),     actual.Volunteer == nil,    test.ExpectedVolunteerLen,    test.ExpectedVolunteerNil},
				{"Work",         len(actual.Work),          actual.Work == nil,         test.ExpectedWorkLen,         test.ExpectedWorkNil},
			}

			for _, check := range checks {
				if check.ExpectedNil != check.ActualNil {
					t.Errorf("For test #%d, %s nil is not what was expected.", testNumber, check.Name)
					t.Logf("EXPECTED nil: %v", check.ExpectedNil)
					t.Logf("ACTUAL nil:   %v", check.ActualNil)
					t.Logf("JSON:\n%s", test.JSON)
					continue
				}
				if check.ExpectedLen != check.ActualLen {
					t.Errorf("For test #%d, %s length is not what was expected.", testNumber, check.Name)
					t.Logf("EXPECTED len: %d", check.ExpectedLen)
					t.Logf("ACTUAL len:   %d", check.ActualLen)
					t.Logf("JSON:\n%s", test.JSON)
					continue
				}
			}
		}
	}
}
