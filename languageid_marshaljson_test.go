package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestLanguageID_MarshalJSON(t *testing.T) {

	tests := []struct {
		Value    jsonresume.LanguageID
		Expected string
	}{
		// 0
		{
			Expected: "null",
		},

		// 1
		{
			Value: jsonresume.SomeLanguageID("http://example.com/resume/language/algonquian"),
			Expected:                       `"http://example.com/resume/language/algonquian"`,
		},

		// 2
		{
			Value: jsonresume.SomeLanguageID("http://example.com/resume/language/english"),
			Expected:                       `"http://example.com/resume/language/english"`,
		},

		// 3
		{
			Value: jsonresume.SomeLanguageID("http://example.com/resume/language/korean"),
			Expected:                       `"http://example.com/resume/language/korean"`,
		},

		// 4
		{
			Value: jsonresume.SomeLanguageID("http://example.com/resume/language/persian"),
			Expected:                       `"http://example.com/resume/language/persian"`,
		},

		// 5
		{
			Value: jsonresume.SomeLanguageID("http://example.com/resume/language/scots"),
			Expected:                       `"http://example.com/resume/language/scots"`,
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
