package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestExperienceID_MarshalJSON(t *testing.T) {

	tests := []struct {
		Value    jsonresume.ExperienceID
		Expected string
	}{
		// 0
		{
			Expected: "null",
		},

		// 1
		{
			Value: jsonresume.SomeExperienceID("http://example.com/resume/experience/3"),
			Expected:                         `"http://example.com/resume/experience/3"`,
		},

		// 2
		{
			Value: jsonresume.SomeExperienceID("http://example.com/resume/experience/2"),
			Expected:                         `"http://example.com/resume/experience/2"`,
		},

		// 3
		{
			Value: jsonresume.SomeExperienceID("http://example.com/resume/experience/1"),
			Expected:                         `"http://example.com/resume/experience/1"`,
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
