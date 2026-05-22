package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestInterestID_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON     []byte
		Expected jsonresume.InterestID
	}{
		// 0
		{
			JSON: []byte("null"),
		},

		// 1
		{
			JSON:                       []byte(`"http://example.com/resume/interest/photography"`),
			Expected: jsonresume.SomeInterestID("http://example.com/resume/interest/photography"),
		},

		// 2
		{
			JSON:                       []byte(`"http://example.com/resume/interest/woodworking"`),
			Expected: jsonresume.SomeInterestID("http://example.com/resume/interest/woodworking"),
		},

		// 3
		{
			JSON:                       []byte(`"http://example.com/resume/interest/hiking"`),
			Expected: jsonresume.SomeInterestID("http://example.com/resume/interest/hiking"),
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.InterestID

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
