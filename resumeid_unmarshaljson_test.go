package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestResumeID_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON     []byte
		Expected jsonresume.ResumeID
	}{
		// 0
		{
			JSON: []byte("null"),
		},

		// 1
		{
			JSON:                     []byte(`"http://example.com/resume/executive"`),
			Expected: jsonresume.SomeResumeID("http://example.com/resume/executive"),
		},

		// 2
		{
			JSON:                     []byte(`"http://example.com/resume/programmer"`),
			Expected: jsonresume.SomeResumeID("http://example.com/resume/programmer"),
		},

		// 3
		{
			JSON:                     []byte(`"http://example.com/resume/designer"`),
			Expected: jsonresume.SomeResumeID("http://example.com/resume/designer"),
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.ResumeID

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
