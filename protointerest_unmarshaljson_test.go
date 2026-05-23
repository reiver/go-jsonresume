package jsonresume_test

import (
	"errors"
	"reflect"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonresume"
)

func TestProtoInterestUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON             []byte
		ExpectNil        bool
		ExpectError      bool
		ExpectedError    error
		ExpectInterestID bool
		ExpectInterest   bool
		Expected         jsonresume.ProtoInterest
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

		// 2: JSON string (IRI) → InterestID
		{
			JSON:             []byte(`"http://example.com/resume/interest/hiking"`),
			ExpectInterestID: true,
			Expected:         jsonresume.SomeInterestID("http://example.com/resume/interest/hiking"),
		},

		// 3: JSON object → Interest
		{
			JSON:           []byte(`{"type":"Interest","keywords":["hiking","camping"]}`),
			ExpectInterest: true,
			Expected: jsonresume.Interest{
				CoreInterest: jsonresume.CoreInterest{
					Keywords: activitypub.SomeStrings("hiking", "camping"),
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
			JSON:             []byte(`"http://example.com/resume/interest/photography"`),
			ExpectInterestID: true,
			Expected:         jsonresume.SomeInterestID("http://example.com/resume/interest/photography"),
		},

		// 7: minimal JSON object
		{
			JSON:           []byte(`{}`),
			ExpectInterest: true,
			Expected:       jsonresume.Interest{},
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonresume.ProtoInterestUnmarshalJSON(test.JSON)

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

		if test.ExpectInterestID {
			if _, ok := actual.(jsonresume.InterestID); !ok {
				t.Errorf("For test #%d, expected type jsonresume.InterestID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectInterest {
			if _, ok := actual.(jsonresume.Interest); !ok {
				t.Errorf("For test #%d, expected type jsonresume.Interest but actually got %T.", testNumber, actual)
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
