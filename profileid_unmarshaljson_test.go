package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestProfileID_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON     []byte
		Expected jsonresume.ProfileID
	}{
		// 0
		{
			JSON: []byte("null"),
		},

		// 1
		{
			JSON:                      []byte(`"http://example.com/resume/profile/mastodon"`),
			Expected: jsonresume.SomeProfileID("http://example.com/resume/profile/mastodon"),
		},

		// 2
		{
			JSON:                      []byte(`"http://example.com/resume/profile/pixelfed"`),
			Expected: jsonresume.SomeProfileID("http://example.com/resume/profile/pixelfed"),
		},

		// 3
		{
			JSON:                      []byte(`"http://example.com/resume/profile/codeberg"`),
			Expected: jsonresume.SomeProfileID("http://example.com/resume/profile/codeberg"),
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.ProfileID

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
