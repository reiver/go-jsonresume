package jsonresume_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestProtoPublicationUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON                []byte
		ExpectNil           bool
		ExpectError         bool
		ExpectedError       error
		ExpectPublicationID bool
		ExpectPublication   bool
		Expected            jsonresume.ProtoPublication
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

		// 2: JSON string (IRI) → PublicationID
		{
			JSON:                []byte(`"http://example.com/resume/publication/my-paper"`),
			ExpectPublicationID: true,
			Expected:            jsonresume.SomePublicationID("http://example.com/resume/publication/my-paper"),
		},

		// 3: JSON object → Publication
		{
			JSON:              []byte(`{"type":"Publication","publisher":"IEEE","releaseDate":"2023-01-15"}`),
			ExpectPublication: true,
			Expected: jsonresume.Publication{
				CorePublication: jsonresume.CorePublication{
					Publisher:   nul.Something("IEEE"),
					ReleaseDate: nul.Something("2023-01-15"),
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
			JSON:                []byte(`"http://example.com/resume/publication/research-2024"`),
			ExpectPublicationID: true,
			Expected:            jsonresume.SomePublicationID("http://example.com/resume/publication/research-2024"),
		},

		// 7: minimal JSON object
		{
			JSON:              []byte(`{}`),
			ExpectPublication: true,
			Expected:          jsonresume.Publication{},
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonresume.ProtoPublicationUnmarshalJSON(test.JSON)

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

		if test.ExpectPublicationID {
			if _, ok := actual.(jsonresume.PublicationID); !ok {
				t.Errorf("For test #%d, expected type jsonresume.PublicationID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectPublication {
			if _, ok := actual.(jsonresume.Publication); !ok {
				t.Errorf("For test #%d, expected type jsonresume.Publication but actually got %T.", testNumber, actual)
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
