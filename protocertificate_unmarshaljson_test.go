package jsonresume

import (
	"errors"
	"reflect"
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

func TestProtoCertificateUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON                []byte
		ExpectNil           bool
		ExpectError         bool
		ExpectedError       error
		ExpectCertificateID  bool
		ExpectAnyCertificate bool
		Expected            ProtoCertificate
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

		// 2: JSON string (IRI) → CertificateID
		{
			JSON:                []byte(`"http://example.com/resume/certificate/aws-solutions-architect"`),
			ExpectCertificateID: true,
			Expected:            SomeCertificateID("http://example.com/resume/certificate/aws-solutions-architect"),
		},

		// 3: JSON object → AnyCertificate
		{
			JSON:              []byte(`{"type":"Certificate","date":"2023-03-15","issuer":"Amazon Web Services"}`),
			ExpectAnyCertificate: true,
			Expected: AnyCertificate{
				Type: jsonld.SomeType("Certificate"),
				CoreCertificate: CoreCertificate{
					Date:   nul.Something("2023-03-15"),
					Issuer: nul.Something("Amazon Web Services"),
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
			JSON:                []byte(`"http://example.com/resume/certificate/cka"`),
			ExpectCertificateID: true,
			Expected:            SomeCertificateID("http://example.com/resume/certificate/cka"),
		},

		// 7: minimal JSON object
		{
			JSON:              []byte(`{}`),
			ExpectAnyCertificate: true,
			Expected:             AnyCertificate{},
		},
	}

	for testNumber, test := range tests {

		actual, err := protoUnmarshalJSON[ProtoCertificate, CertificateID, AnyCertificate](test.JSON)

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

		if test.ExpectCertificateID {
			if _, ok := actual.(CertificateID); !ok {
				t.Errorf("For test #%d, expected type CertificateID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectAnyCertificate {
			if _, ok := actual.(AnyCertificate); !ok {
				t.Errorf("For test #%d, expected type AnyCertificate but actually got %T.", testNumber, actual)
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
