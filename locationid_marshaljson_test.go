package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestLocationID_MarshalJSON(t *testing.T) {

	tests := []struct {
		Value    jsonresume.LocationID
		Expected string
	}{
		// 0
		{
			Expected: "null",
		},

		// 1
		{
			Value: jsonresume.SomeLocationID("http://example.com/resume/location"),
			Expected:                       `"http://example.com/resume/location"`,
		},

		// 2
		{
			Value: jsonresume.SomeLocationID("http://example.com/resume/location/vancouver"),
			Expected:                       `"http://example.com/resume/location/vancouver"`,
		},

		// 3
		{
			Value: jsonresume.SomeLocationID("http://example.com/resume/location/toronto"),
			Expected:                       `"http://example.com/resume/location/toronto"`,
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
