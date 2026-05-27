package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestMetaID_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON     []byte
		Expected jsonresume.MetaID
	}{
		// 0
		{
			JSON: []byte("null"),
		},

		// 1
		{
			JSON:                   []byte(`"http://example.com/resume/meta"`),
			Expected: jsonresume.SomeMetaID("http://example.com/resume/meta"),
		},

		// 2
		{
			JSON:                   []byte(`"http://example.com/resume/meta/v2"`),
			Expected: jsonresume.SomeMetaID("http://example.com/resume/meta/v2"),
		},

		// 3
		{
			JSON:                   []byte(`"https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json"`),
			Expected: jsonresume.SomeMetaID("https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json"),
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.MetaID

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
