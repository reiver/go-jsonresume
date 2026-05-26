package jsonresume

import (
	"errors"
	"reflect"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

func TestProtoExperienceUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON               []byte
		ExpectNil          bool
		ExpectError        bool
		ExpectedError      error
		ExpectExperienceID  bool
		ExpectAnyExperience bool
		Expected           ProtoExperience
	}{
		// 0: empty bytes
		{
			JSON:          []byte{},
			ExpectError:   true,
			ExpectedError: jsonld.ErrBytesEmpty,
		},

		// 1: null
		{
			JSON:      []byte("null"),
			ExpectNil: true,
		},

		// 2: JSON string (IRI) → ExperienceID
		{
			JSON:               []byte(`"http://example.com/resume/experience/senior-dev"`),
			ExpectExperienceID: true,
			Expected:           SomeExperienceID("http://example.com/resume/experience/senior-dev"),
		},

		// 3: JSON object → Experience (work)
		{
			JSON:             []byte(`{"type":"Experience","organization":"Acme Corp","startDate":"2020-01-15","endDate":"2024-12-31"}`),
			ExpectAnyExperience: true,
			Expected: AnyExperience{
				Type: jsonld.SomeType("Experience"),
				CoreExperience: CoreExperience{
					Organization: nul.Something("Acme Corp"),
					StartDate:    nul.Something("2020-01-15"),
					EndDate:      nul.Something("2024-12-31"),
				},
			},
		},

		// 4: JSON object → Experience with position and highlights
		{
			JSON:             []byte(`{"type":"Experience","position":["Senior Developer"],"highlights":["Led team of 5","Shipped v2.0"]}`),
			ExpectAnyExperience: true,
			Expected: AnyExperience{
				Type: jsonld.SomeType("Experience"),
				CoreExperience: CoreExperience{
					Position:   activitypub.SomeStrings("Senior Developer"),
					Highlights: activitypub.SomeStrings("Led team of 5", "Shipped v2.0"),
				},
			},
		},

		// 5: JSON object → Experience with all fields
		{
			JSON:             []byte(`{"type":"Experience","organization":"Acme Corp","position":["Senior Developer"],"startDate":"2020-01-15","endDate":"2024-12-31","highlights":["Led team of 5","Shipped v2.0"]}`),
			ExpectAnyExperience: true,
			Expected: AnyExperience{
				Type: jsonld.SomeType("Experience"),
				CoreExperience: CoreExperience{
					Organization: nul.Something("Acme Corp"),
					Position:     activitypub.SomeStrings("Senior Developer"),
					StartDate:    nul.Something("2020-01-15"),
					EndDate:      nul.Something("2024-12-31"),
					Highlights:   activitypub.SomeStrings("Led team of 5", "Shipped v2.0"),
				},
			},
		},

		// 6: unsupported type (array)
		{
			JSON:          []byte(`[1,2,3]`),
			ExpectError:   true,
			ExpectedError: jsonld.ErrJSONTypeUnsupported,
		},

		// 7: unsupported type (number)
		{
			JSON:          []byte(`42`),
			ExpectError:   true,
			ExpectedError: jsonld.ErrJSONTypeUnsupported,
		},

		// 8: another JSON string (IRI) — volunteer
		{
			JSON:               []byte(`"http://example.com/resume/experience/open-source-volunteer"`),
			ExpectExperienceID: true,
			Expected:           SomeExperienceID("http://example.com/resume/experience/open-source-volunteer"),
		},

		// 9: minimal JSON object
		{
			JSON:             []byte(`{}`),
			ExpectAnyExperience: true,
			Expected:            AnyExperience{},
		},
	}

	for testNumber, test := range tests {

		var actual ProtoExperience
		err := jsonld.UnmarshalJSONStringOrJSONObject[ProtoExperience, ExperienceID, AnyExperience](test.JSON, &actual)

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

		if test.ExpectExperienceID {
			if _, ok := actual.(ExperienceID); !ok {
				t.Errorf("For test #%d, expected type ExperienceID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectAnyExperience {
			if _, ok := actual.(AnyExperience); !ok {
				t.Errorf("For test #%d, expected type AnyExperience but actually got %T.", testNumber, actual)
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
