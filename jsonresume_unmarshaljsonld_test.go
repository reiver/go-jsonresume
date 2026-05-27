package jsonresume_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"
)

func TestJSONResume_unmarshalJSONLD(t *testing.T) {

	cmpOpts := cmp.Options{
		cmpopts.EquateComparable(
			jsonld.ID{},
			jsonld.Types{},
			nul.Nullable[string]{},
			opt.Optional[string]{},
			jsonresume.ResumeID{},
		),
	}

	tests := []struct {
		JSON     string
		Expected jsonresume.JSONResume
	}{
		// 0: empty JSON object
		{
			JSON: `{}`,
		},

		// 1: resume field omitted
		{
			JSON: `{` +
				`"@context":{` +
				`"cv":"https://w3id.org/fep/6158"` +
				`}` +
				`}`,
		},

		// 2: resume null
		{
			JSON: `{"resume":null}`,
		},

		// 3: resume empty array
		{
			JSON: `{"resume":[]}`,
		},

		// 4: resume empty array with @context
		{
			JSON: `{` +
				`"@context":{` +
				`"cv":"https://w3id.org/fep/6158"` +

				`,` +

				`"resume":"cv:resume"` +
				`}` +
				`,` +
				`"resume":[]` +
				`}`,
		},

		// 5: resume null with @context
		{
			JSON: `{` +
				`"@context":{` +
				`"cv":"https://w3id.org/fep/6158"` +

				`,` +

				`"resume":"cv:resume"` +
				`}` +
				`,` +
				`"resume":null` +
				`}`,
		},

		// 6: with both AS and CV namespaces in @context (as produced by jsonld.Marshal with activitypub.Name)
		{
			JSON: `{` +
				`"@context":{` +
				`"as":"https://www.w3.org/ns/activitystreams"` +
				`,` +
				`"cv":"https://w3id.org/fep/6158"` +

				`,` +

				`"name":"as:name"` +

				`,` +

				`"resume":"cv:resume"` +
				`}` +
				`,` +
				`"name":"Joe Blow"` +
				`,` +
				`"resume":[]` +
				`}`,
		},

		// 7: only @context, no resume field
		{
			JSON: `{` +
				`"@context":{` +
				`"as":"https://www.w3.org/ns/activitystreams"` +
				`,` +
				`"cv":"https://w3id.org/fep/6158"` +

				`,` +

				`"name":"as:name"` +

				`,` +

				`"resume":"cv:resume"` +
				`}` +
				`,` +
				`"name":"Joe Blow"` +
				`}`,
		},

		// 8: extra unknown fields (should be ignored)
		{
			JSON: `{` +
				`"resume":[]` +
				`,` +
				`"unknownField":"some value"` +
				`,` +
				`"anotherUnknown":42` +
				`}`,
		},

		// 9: resume is a single IRI string (not in an array)
		{
			JSON: `{"resume":"https://example.com/resume"}`,
			Expected: jsonresume.JSONResume{
				Resume: []jsonresume.ProtoResume{
					jsonresume.SomeResumeID("https://example.com/resume"),
				},
			},
		},

		// 10: resume is an array with a single IRI string
		{
			JSON: `{"resume":["https://example.com/resume"]}`,
			Expected: jsonresume.JSONResume{
				Resume: []jsonresume.ProtoResume{
					jsonresume.SomeResumeID("https://example.com/resume"),
				},
			},
		},

		// 11: resume is an array with multiple IRI strings
		{
			JSON: `{` +
				`"resume":[` +
				`"https://example.com/resume/executive"` +
				`,` +
				`"https://example.com/resume/programmer"` +
				`]` +
				`}`,
			Expected: jsonresume.JSONResume{
				Resume: []jsonresume.ProtoResume{
					jsonresume.SomeResumeID("https://example.com/resume/executive"),
					jsonresume.SomeResumeID("https://example.com/resume/programmer"),
				},
			},
		},

		// 12: resume is an array with IRI strings and @context
		{
			JSON: `{` +
				`"@context":{` +
				`"cv":"https://w3id.org/fep/6158"` +

				`,` +

				`"resume":"cv:resume"` +
				`}` +
				`,` +
				`"resume":[` +
				`"https://example.com/resume/executive"` +
				`]` +
				`}`,
			Expected: jsonresume.JSONResume{
				Resume: []jsonresume.ProtoResume{
					jsonresume.SomeResumeID("https://example.com/resume/executive"),
				},
			},
		},

		// 13: resume is an array mixing an IRI string and an object
		{
			JSON: `{` +
				`"resume":[` +
				`"https://example.com/resume/executive"` +
				`,` +
				`{"type":"Resume"}` +
				`]` +
				`}`,
			Expected: jsonresume.JSONResume{
				Resume: []jsonresume.ProtoResume{
					jsonresume.SomeResumeID("https://example.com/resume/executive"),
					jsonresume.AnyResume{
						Type: jsonld.SomeType("Resume"),
					},
				},
			},
		},
	}

	for testNumber, test := range tests {

		t.Run(fmt.Sprintf("%d", testNumber), func(t *testing.T) {

			defer func() {
				if r := recover(); nil != r {
					t.Errorf("For test #%d, unmarshal panicked.", testNumber)
					t.Logf("PANIC: %v", r)
					t.Logf("JSON:\n%s", test.JSON)
				}
			}()

			var actual jsonresume.JSONResume

			err := jsonld.Unmarshal([]byte(test.JSON), &actual)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:\n%s", test.JSON)
				return
			}

			if diff := cmp.Diff(test.Expected, actual, cmpOpts); diff != "" {
				t.Errorf("For test #%d, actual result is not what was expected (-expected +actual):\n%s", testNumber, diff)
				t.Logf("JSON:\n%s", test.JSON)
				return
			}
		})
	}
}
