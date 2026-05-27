package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestMetaID_MarshalJSON(t *testing.T) {

	tests := []struct {
		Value    jsonresume.MetaID
		Expected string
	}{
		// 0
		{
			Expected: "null",
		},

		// 1
		{
			Value: jsonresume.SomeMetaID("http://example.com/resume/meta"),
			Expected:                   `"http://example.com/resume/meta"`,
		},

		// 2
		{
			Value: jsonresume.SomeMetaID("http://example.com/resume/meta/v2"),
			Expected:                   `"http://example.com/resume/meta/v2"`,
		},

		// 3
		{
			Value: jsonresume.SomeMetaID("https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json"),
			Expected:                   `"https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json"`,
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
