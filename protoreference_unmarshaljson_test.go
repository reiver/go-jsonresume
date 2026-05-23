package jsonresume_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestProtoReferenceUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON              []byte
		ExpectNil         bool
		ExpectError       bool
		ExpectedError     error
		ExpectReferenceID bool
		ExpectReference   bool
		Expected          jsonresume.ProtoReference
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

		// 2: JSON string (IRI) → ReferenceID
		{
			JSON:              []byte(`"http://example.com/resume/reference/jane-doe"`),
			ExpectReferenceID: true,
			Expected:          jsonresume.SomeReferenceID("http://example.com/resume/reference/jane-doe"),
		},

		// 3: JSON object → Reference
		{
			JSON:            []byte(`{"type":"Reference","reference":"Joe is a great worker."}`),
			ExpectReference: true,
			Expected: jsonresume.Reference{
				CoreReference: jsonresume.CoreReference{
					Reference: nul.Something("Joe is a great worker."),
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
			JSON:              []byte(`"http://example.com/resume/reference/john-smith"`),
			ExpectReferenceID: true,
			Expected:          jsonresume.SomeReferenceID("http://example.com/resume/reference/john-smith"),
		},

		// 7: minimal JSON object
		{
			JSON:            []byte(`{}`),
			ExpectReference: true,
			Expected:        jsonresume.Reference{},
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonresume.ProtoReferenceUnmarshalJSON(test.JSON)

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

		if test.ExpectReferenceID {
			if _, ok := actual.(jsonresume.ReferenceID); !ok {
				t.Errorf("For test #%d, expected type jsonresume.ReferenceID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectReference {
			if _, ok := actual.(jsonresume.Reference); !ok {
				t.Errorf("For test #%d, expected type jsonresume.Reference but actually got %T.", testNumber, actual)
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
