package jsonresume

import (
	"errors"
	"reflect"
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

func TestProtoPublicationUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON                []byte
		ExpectNil           bool
		ExpectError         bool
		ExpectedError       error
		ExpectPublicationID  bool
		ExpectAnyPublication bool
		Expected            ProtoPublication
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

		// 2: JSON string (IRI) → PublicationID
		{
			JSON:                []byte(`"http://example.com/resume/publication/my-paper"`),
			ExpectPublicationID: true,
			Expected:            SomePublicationID("http://example.com/resume/publication/my-paper"),
		},

		// 3: JSON object → Publication
		{
			JSON:              []byte(`{"type":"Publication","publisher":"IEEE","releaseDate":"2023-01-15"}`),
			ExpectAnyPublication: true,
			Expected: AnyPublication{
				Type: jsonld.SomeType("Publication"),
				CorePublication: CorePublication{
					Publisher:   nul.Something("IEEE"),
					ReleaseDate: nul.Something("2023-01-15"),
				},
			},
		},

		// 4: unsupported type (array)
		{
			JSON:          []byte(`[1,2,3]`),
			ExpectError:   true,
			ExpectedError: ErrTypeUnsupported,
		},

		// 5: unsupported type (number)
		{
			JSON:          []byte(`42`),
			ExpectError:   true,
			ExpectedError: ErrTypeUnsupported,
		},

		// 6: another JSON string (IRI)
		{
			JSON:                []byte(`"http://example.com/resume/publication/research-2024"`),
			ExpectPublicationID: true,
			Expected:            SomePublicationID("http://example.com/resume/publication/research-2024"),
		},

		// 7: minimal JSON object
		{
			JSON:              []byte(`{}`),
			ExpectAnyPublication: true,
			Expected:             AnyPublication{},
		},
	}

	for testNumber, test := range tests {

		actual, err := protoUnmarshalJSON[ProtoPublication, PublicationID, AnyPublication](test.JSON)

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
			if _, ok := actual.(PublicationID); !ok {
				t.Errorf("For test #%d, expected type PublicationID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectAnyPublication {
			if _, ok := actual.(AnyPublication); !ok {
				t.Errorf("For test #%d, expected type AnyPublication but actually got %T.", testNumber, actual)
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
