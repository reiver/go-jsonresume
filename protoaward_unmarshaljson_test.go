package jsonresume_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestProtoAwardUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON          []byte
		ExpectNil     bool
		ExpectError   bool
		ExpectedError error
		ExpectAwardID bool
		ExpectAward   bool
		Expected      jsonresume.ProtoAward
	}{
		// 0: empty bytes
		{
			JSON:          []byte{},
			ExpectError:   true,
			ExpectedError: jsonresume.ErrBytesEmpty,
		},

		// 1: null
		{
			JSON:      []byte("null"),
			ExpectNil: true,
		},

		// 2: JSON string (IRI) → AwardID
		{
			JSON:          []byte(`"http://example.com/resume/award/best-employee-2024"`),
			ExpectAwardID: true,
			Expected:      jsonresume.SomeAwardID("http://example.com/resume/award/best-employee-2024"),
		},

		// 3: JSON object → Award
		{
			JSON:        []byte(`{"type":"Award","awarder":"SuperCo","date":"2024-05-21","title":"Best Employee (2024)"}`),
			ExpectAward: true,
			Expected: jsonresume.Award{
				CoreAward: jsonresume.CoreAward{
					Awarder: nul.Something("SuperCo"),
					Date:    nul.Something("2024-05-21"),
					Title:   nul.Something("Best Employee (2024)"),
				},
			},
		},

		// 4: unsupported type (array)
		{
			JSON:          []byte(`[1,2,3]`),
			ExpectError:   true,
			ExpectedError: jsonresume.ErrTypeUnsupported,
		},

		// 5: unsupported type (number)
		{
			JSON:          []byte(`42`),
			ExpectError:   true,
			ExpectedError: jsonresume.ErrTypeUnsupported,
		},

		// 6: another JSON string (IRI)
		{
			JSON:          []byte(`"http://example.com/resume/award/innovation-prize"`),
			ExpectAwardID: true,
			Expected:      jsonresume.SomeAwardID("http://example.com/resume/award/innovation-prize"),
		},

		// 7: minimal JSON object
		{
			JSON:        []byte(`{}`),
			ExpectAward: true,
			Expected:    jsonresume.Award{},
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonresume.ProtoAwardUnmarshalJSON(test.JSON)

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

		if test.ExpectAwardID {
			if _, ok := actual.(jsonresume.AwardID); !ok {
				t.Errorf("For test #%d, expected type jsonresume.AwardID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectAward {
			if _, ok := actual.(jsonresume.Award); !ok {
				t.Errorf("For test #%d, expected type jsonresume.Award but actually got %T.", testNumber, actual)
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
