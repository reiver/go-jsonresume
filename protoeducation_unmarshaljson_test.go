package jsonresume

import (
	"errors"
	"reflect"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

func TestProtoEducationUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON              []byte
		ExpectNil         bool
		ExpectError       bool
		ExpectedError     error
		ExpectEducationID  bool
		ExpectAnyEducation bool
		Expected          ProtoEducation
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

		// 2: JSON string (IRI) → EducationID
		{
			JSON:              []byte(`"http://example.com/resume/education/bs-compsci"`),
			ExpectEducationID: true,
			Expected:          SomeEducationID("http://example.com/resume/education/bs-compsci"),
		},

		// 3: JSON object → Education
		{
			JSON:            []byte(`{"type":"Education","institution":"MIT","score":"3.8","startDate":"2010-09-01","endDate":"2014-06-15"}`),
			ExpectAnyEducation: true,
			Expected: AnyEducation{
				Type: jsonld.SomeType("Education"),
				CoreEducation: CoreEducation{
					Institution: nul.Something("MIT"),
					Score:       nul.Something("3.8"),
					StartDate:   nul.Something("2010-09-01"),
					EndDate:     nul.Something("2014-06-15"),
				},
			},
		},

		// 4: JSON object → Education with area, studyType, courses
		{
			JSON:            []byte(`{"type":"Education","area":["Computer Science"],"studyType":["Bachelor"],"courses":["CS101","CS201"]}`),
			ExpectAnyEducation: true,
			Expected: AnyEducation{
				Type: jsonld.SomeType("Education"),
				CoreEducation: CoreEducation{
					Area:      activitypub.SomeStrings("Computer Science"),
					StudyType: activitypub.SomeStrings("Bachelor"),
					Courses:   activitypub.SomeStrings("CS101", "CS201"),
				},
			},
		},

		// 5: JSON object → Education with all fields
		{
			JSON:            []byte(`{"type":"Education","institution":"MIT","area":["Computer Science"],"studyType":["Bachelor"],"startDate":"2010-09-01","endDate":"2014-06-15","score":"3.8","courses":["CS101","CS201","CS301"]}`),
			ExpectAnyEducation: true,
			Expected: AnyEducation{
				Type: jsonld.SomeType("Education"),
				CoreEducation: CoreEducation{
					Institution: nul.Something("MIT"),
					Area:        activitypub.SomeStrings("Computer Science"),
					StudyType:   activitypub.SomeStrings("Bachelor"),
					StartDate:   nul.Something("2010-09-01"),
					EndDate:     nul.Something("2014-06-15"),
					Score:       nul.Something("3.8"),
					Courses:     activitypub.SomeStrings("CS101", "CS201", "CS301"),
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
			JSON:              []byte(`"http://example.com/resume/education/ms-physics"`),
			ExpectEducationID: true,
			Expected:          SomeEducationID("http://example.com/resume/education/ms-physics"),
		},

		// 9: minimal JSON object
		{
			JSON:            []byte(`{}`),
			ExpectAnyEducation: true,
			Expected:           AnyEducation{},
		},
	}

	for testNumber, test := range tests {

		actual, err := protoUnmarshalJSON[ProtoEducation, EducationID, AnyEducation](test.JSON)

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

		if test.ExpectEducationID {
			if _, ok := actual.(EducationID); !ok {
				t.Errorf("For test #%d, expected type EducationID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectAnyEducation {
			if _, ok := actual.(AnyEducation); !ok {
				t.Errorf("For test #%d, expected type AnyEducation but actually got %T.", testNumber, actual)
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
