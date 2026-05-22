package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestResumeID_MarshalJSON(t *testing.T) {

	tests := []struct {
		Value    jsonresume.ResumeID
		Expected string
	}{
		// 0
		{
			Expected: "null",
		},

		// 1
		{
			Value: jsonresume.SomeResumeID("http://example.com/resume/executive"),
			Expected:                     `"http://example.com/resume/executive"`,
		},

		// 2
		{
			Value: jsonresume.SomeResumeID("http://example.com/resume/programmer"),
			Expected:                     `"http://example.com/resume/programmer"`,
		},

		// 3
		{
			Value: jsonresume.SomeResumeID("http://example.com/resume/designer"),
			Expected:                     `"http://example.com/resume/designer"`,
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
