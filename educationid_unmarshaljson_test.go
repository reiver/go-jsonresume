package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestEducationID_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON     []byte
		Expected jsonresume.EducationID
	}{
		// 0
		{
			JSON: []byte("null"),
		},

		// 1
		{
			JSON:                        []byte(`"http://example.com/resume/education/sfu"`),
			Expected: jsonresume.SomeEducationID("http://example.com/resume/education/sfu"),
		},

		// 2
		{
			JSON:                        []byte(`"http://example.com/resume/education/kpu"`),
			Expected: jsonresume.SomeEducationID("http://example.com/resume/education/kpu"),
		},

		// 3
		{
			JSON:                        []byte(`"http://example.com/resume/education/mit"`),
			Expected: jsonresume.SomeEducationID("http://example.com/resume/education/mit"),
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.EducationID

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
