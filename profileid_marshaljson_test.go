package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestProfileID_MarshalJSON(t *testing.T) {

	tests := []struct {
		Value    jsonresume.ProfileID
		Expected string
	}{
		// 0
		{
			Expected: "null",
		},

		// 1
		{
			Value: jsonresume.SomeProfileID("http://example.com/resume/profile/mastodon"),
			Expected:                      `"http://example.com/resume/profile/mastodon"`,
		},

		// 2
		{
			Value: jsonresume.SomeProfileID("http://example.com/resume/profile/pixelfed"),
			Expected:                      `"http://example.com/resume/profile/pixelfed"`,
		},

		// 3
		{
			Value: jsonresume.SomeProfileID("http://example.com/resume/profile/github"),
			Expected:                      `"http://example.com/resume/profile/github"`,
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
