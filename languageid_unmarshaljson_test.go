package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestLanguageID_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON     []byte
		Expected jsonresume.LanguageID
	}{
		// 0
		{
			JSON: []byte("null"),
		},

		// 1
		{
			JSON:                      []byte(`"http://example.com/resume/language/algonquian"`),
			Expected: jsonresume.SomeLanguageID("http://example.com/resume/language/algonquian"),
		},

		// 2
		{
			JSON:                       []byte(`"http://example.com/resume/language/english"`),
			Expected: jsonresume.SomeLanguageID("http://example.com/resume/language/english"),
		},

		// 3
		{
			JSON:                       []byte(`"http://example.com/resume/language/korean"`),
			Expected: jsonresume.SomeLanguageID("http://example.com/resume/language/korean"),
		},

		// 4
		{
			JSON:                       []byte(`"http://example.com/resume/language/persian"`),
			Expected: jsonresume.SomeLanguageID("http://example.com/resume/language/persian"),
		},

		// 5
		{
			JSON:                       []byte(`"http://example.com/resume/language/scots"`),
			Expected: jsonresume.SomeLanguageID("http://example.com/resume/language/scots"),
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.LanguageID

		err := actual.UnmarshalJSON(test.JSON)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
