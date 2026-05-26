package jsonresume

import (
	"errors"
	"reflect"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

func TestProtoProjectUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON            []byte
		ExpectNil       bool
		ExpectError     bool
		ExpectedError   error
		ExpectProjectID  bool
		ExpectAnyProject bool
		Expected        ProtoProject
	}{
		// 0: empty bytes
		{
			JSON:          []byte{},
			ExpectError:   true,
			ExpectedError: ErrBytesEmpty,
		},

		// 1: null
		{
			JSON:      []byte("null"),
			ExpectNil: true,
		},

		// 2: JSON string (IRI) → ProjectID
		{
			JSON:            []byte(`"http://example.com/resume/project/cool-app"`),
			ExpectProjectID: true,
			Expected:        SomeProjectID("http://example.com/resume/project/cool-app"),
		},

		// 3: JSON object → Project
		{
			JSON:          []byte(`{"type":"Project","entity":"Acme Corp","startDate":"2023-06-01","description":"A cool project."}`),
			ExpectAnyProject: true,
			Expected: AnyProject{
				Type: jsonld.SomeType("Project"),
				CoreProject: CoreProject{
					Entity:      nul.Something("Acme Corp"),
					StartDate:   nul.Something("2023-06-01"),
					Description: nul.Something("A cool project."),
				},
			},
		},

		// 4: JSON object → Project with highlights, keywords, roles
		{
			JSON:          []byte(`{"type":"Project","highlights":["Built API","Wrote docs"],"keywords":["Go","REST"],"roles":["Lead","Developer"]}`),
			ExpectAnyProject: true,
			Expected: AnyProject{
				Type: jsonld.SomeType("Project"),
				CoreProject: CoreProject{
					Highlights: activitypub.SomeStrings("Built API", "Wrote docs"),
					Keywords:   activitypub.SomeStrings("Go", "REST"),
					Roles:      activitypub.SomeStrings("Lead", "Developer"),
				},
			},
		},

		// 5: JSON object → Project with all fields
		{
			JSON:          []byte(`{"type":"Project","description":"A cool project.","entity":"Acme Corp","startDate":"2023-06-01","endDate":"2024-01-31","highlights":["Built API"],"keywords":["Go"],"roles":["Lead"]}`),
			ExpectAnyProject: true,
			Expected: AnyProject{
				Type: jsonld.SomeType("Project"),
				CoreProject: CoreProject{
					Description: nul.Something("A cool project."),
					Entity:      nul.Something("Acme Corp"),
					StartDate:   nul.Something("2023-06-01"),
					EndDate:     nul.Something("2024-01-31"),
					Highlights:  activitypub.SomeStrings("Built API"),
					Keywords:    activitypub.SomeStrings("Go"),
					Roles:       activitypub.SomeStrings("Lead"),
				},
			},
		},

		// 6: unsupported type (array)
		{
			JSON:          []byte(`[1,2,3]`),
			ExpectError:   true,
			ExpectedError: ErrTypeUnsupported,
		},

		// 7: unsupported type (number)
		{
			JSON:          []byte(`42`),
			ExpectError:   true,
			ExpectedError: ErrTypeUnsupported,
		},

		// 8: another JSON string (IRI)
		{
			JSON:            []byte(`"http://example.com/resume/project/side-project"`),
			ExpectProjectID: true,
			Expected:        SomeProjectID("http://example.com/resume/project/side-project"),
		},

		// 9: minimal JSON object
		{
			JSON:          []byte(`{}`),
			ExpectAnyProject: true,
			Expected:         AnyProject{},
		},
	}

	for testNumber, test := range tests {

		actual, err := protoUnmarshalJSON[ProtoProject, ProjectID, AnyProject](test.JSON)

		if test.ExpectError {
			if nil == err {
				t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}

			if nil != test.ExpectedError {
				if !errors.Is(err, test.ExpectedError) {
					t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
					t.Logf("EXPECTED-ERROR: %s", test.ExpectedError)
					t.Logf("ACTUAL-ERROR:   %s", err)
					t.Logf("JSON:\n%s", test.JSON)
					continue
				}
			}

			continue
		}

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		if test.ExpectNil {
			if nil != actual {
				t.Errorf("For test #%d, expected nil but actually got non-nil.", testNumber)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
			}
			continue
		}

		if nil == actual {
			t.Errorf("For test #%d, did not expect nil but actually got nil.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		if test.ExpectProjectID {
			if _, ok := actual.(ProjectID); !ok {
				t.Errorf("For test #%d, expected type ProjectID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectAnyProject {
			if _, ok := actual.(AnyProject); !ok {
				t.Errorf("For test #%d, expected type AnyProject but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
