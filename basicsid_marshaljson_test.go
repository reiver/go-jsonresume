package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestBasicsID_MarshalJSON(t *testing.T) {

	tests := []struct {
		Value    jsonresume.BasicsID
		Expected string
	}{
		// 0
		{
			Expected: "null",
		},

		// 1
		{
			Value: jsonresume.SomeBasicsID("http://example.com/resume/basics"),
			Expected:                     `"http://example.com/resume/basics"`,
		},

		// 2
		{
			Value: jsonresume.SomeBasicsID("http://example.com/resume/basics/joe-blow"),
			Expected:                     `"http://example.com/resume/basics/joe-blow"`,
		},

		// 3
		{
			Value: jsonresume.SomeBasicsID("http://example.com/resume/basics/jane-doe"),
			Expected:                     `"http://example.com/resume/basics/jane-doe"`,
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
