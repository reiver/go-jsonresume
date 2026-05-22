package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestCertificateID_MarshalJSON(t *testing.T) {

	tests := []struct {
		Value    jsonresume.CertificateID
		Expected string
	}{
		// 0
		{
			Expected: "null",
		},

		// 1
		{
			Value: jsonresume.SomeCertificateID("http://example.com/resume/certificate/8"),
			Expected:                          `"http://example.com/resume/certificate/8"`,
		},

		// 2
		{
			Value: jsonresume.SomeCertificateID("http://example.com/resume/certificate/7"),
			Expected:                          `"http://example.com/resume/certificate/7"`,
		},

		// 3
		{
			Value: jsonresume.SomeCertificateID("http://example.com/resume/certificate/aws-solutions-architect"),
			Expected:                          `"http://example.com/resume/certificate/aws-solutions-architect"`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Value.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		actual := string(actualBytes)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
