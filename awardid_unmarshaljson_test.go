package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestAwardID_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON     []byte
		Expected jsonresume.AwardID
	}{
		// 0
		{
			JSON: []byte("null"),
		},

		// 1
		{
			JSON:                    []byte(`"http://example.com/resume/award/best-employee-2024"`),
			Expected: jsonresume.SomeAwardID("http://example.com/resume/award/best-employee-2024"),
		},

		// 2
		{
			JSON:                    []byte(`"http://example.com/resume/award/acme-excellence-2021"`),
			Expected: jsonresume.SomeAwardID("http://example.com/resume/award/acme-excellence-2021"),
		},

		// 3
		{
			JSON:                    []byte(`"http://example.com/resume/award/innovation-prize"`),
			Expected: jsonresume.SomeAwardID("http://example.com/resume/award/innovation-prize"),
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.AwardID

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
