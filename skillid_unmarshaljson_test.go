package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestSkillID_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON     []byte
		Expected jsonresume.SkillID
	}{
		// 0
		{
			JSON: []byte("null"),
		},

		// 1
		{
			JSON:                    []byte(`"http://example.com/resume/skill/backend-development"`),
			Expected: jsonresume.SomeSkillID("http://example.com/resume/skill/backend-development"),
		},

		// 2
		{
			JSON:                    []byte(`"http://example.com/resume/skill/web-development"`),
			Expected: jsonresume.SomeSkillID("http://example.com/resume/skill/web-development"),
		},

		// 3
		{
			JSON:                    []byte(`"http://example.com/resume/skill/go"`),
			Expected: jsonresume.SomeSkillID("http://example.com/resume/skill/go"),
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.SkillID

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
