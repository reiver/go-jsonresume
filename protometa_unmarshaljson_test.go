package jsonresume

import (
	"errors"
	"reflect"
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

func TestProtoMetaUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON           []byte
		ExpectNil      bool
		ExpectError    bool
		ExpectedError  error
		ExpectMetaID   bool
		ExpectAnyMeta  bool
		Expected       ProtoMeta
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

		// 2: JSON string (IRI) → MetaID
		{
			JSON:         []byte(`"http://example.com/resume/meta"`),
			ExpectMetaID: true,
			Expected:     SomeMetaID("http://example.com/resume/meta"),
		},

		// 3: JSON object → AnyMeta
		{
			JSON:        []byte(`{"type":"Meta","canonical":"https://example.com/resume.json","version":"v1.0.0","lastModified":"2017-12-24T15:53:00"}`),
			ExpectAnyMeta: true,
			Expected: AnyMeta{
				Type: jsonld.SomeType("Meta"),
				CoreMeta: CoreMeta{
					Canonical:    nul.Something("https://example.com/resume.json"),
					Version:      nul.Something("v1.0.0"),
					LastModified: nul.Something("2017-12-24T15:53:00"),
				},
			},
		},

		// 4: unsupported type (array)
		{
			JSON:          []byte(`[1,2,3]`),
			ExpectError:   true,
			ExpectedError: jsonld.ErrJSONTypeUnsupported,
		},

		// 5: unsupported type (number)
		{
			JSON:          []byte(`42`),
			ExpectError:   true,
			ExpectedError: jsonld.ErrJSONTypeUnsupported,
		},

		// 6: another JSON string (IRI)
		{
			JSON:         []byte(`"https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json"`),
			ExpectMetaID: true,
			Expected:     SomeMetaID("https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json"),
		},

		// 7: minimal JSON object
		{
			JSON:          []byte(`{}`),
			ExpectAnyMeta: true,
			Expected:      AnyMeta{},
		},
	}

	for testNumber, test := range tests {

		var actual ProtoMeta
		err := jsonld.UnmarshalJSONStringOrJSONObject[ProtoMeta, MetaID, AnyMeta](test.JSON, &actual)

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

		if test.ExpectMetaID {
			if _, ok := actual.(MetaID); !ok {
				t.Errorf("For test #%d, expected type MetaID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectAnyMeta {
			if _, ok := actual.(AnyMeta); !ok {
				t.Errorf("For test #%d, expected type AnyMeta but actually got %T.", testNumber, actual)
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
