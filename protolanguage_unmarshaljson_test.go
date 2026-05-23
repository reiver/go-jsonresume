package jsonresume_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestProtoLanguageUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON             []byte
		ExpectNil        bool
		ExpectError      bool
		ExpectedError    error
		ExpectLanguageID bool
		ExpectLanguage   bool
		Expected         jsonresume.ProtoLanguage
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

		// 2: JSON string (IRI) → LanguageID
		{
			JSON:             []byte(`"http://example.com/resume/language/english"`),
			ExpectLanguageID: true,
			Expected:         jsonresume.SomeLanguageID("http://example.com/resume/language/english"),
		},

		// 3: JSON object → Language
		{
			JSON:           []byte(`{"type":"Language","fluency":"native","language":"English"}`),
			ExpectLanguage: true,
			Expected: jsonresume.Language{
				CoreLanguage: jsonresume.CoreLanguage{
					Fluency:  nul.Something("native"),
					Language: nul.Something("English"),
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
			JSON:             []byte(`"http://example.com/resume/language/french"`),
			ExpectLanguageID: true,
			Expected:         jsonresume.SomeLanguageID("http://example.com/resume/language/french"),
		},

		// 7: minimal JSON object
		{
			JSON:           []byte(`{}`),
			ExpectLanguage: true,
			Expected:       jsonresume.Language{},
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonresume.ProtoLanguageUnmarshalJSON(test.JSON)

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

		if test.ExpectLanguageID {
			if _, ok := actual.(jsonresume.LanguageID); !ok {
				t.Errorf("For test #%d, expected type jsonresume.LanguageID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectLanguage {
			if _, ok := actual.(jsonresume.Language); !ok {
				t.Errorf("For test #%d, expected type jsonresume.Language but actually got %T.", testNumber, actual)
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
