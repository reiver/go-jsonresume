package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestCoreBasics_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"email":"cv:email"` +
		`,` +
		`"label":"cv:label"` +
		`,` +
		`"phone":"cv:phone"` +
		`,` +
		`"profiles":"cv:profiles"` +
	`}`

	tests := []struct {
		Value    jsonresume.CoreBasics
		Expected string
	}{
		// 0
		{
			Value: jsonresume.CoreBasics{},
			Expected: `{` +
				context +
				`,` +
				`"email":null` +
				`,` +
				`"label":null` +
				`,` +
				`"phone":null` +
				`,` +
				`"profiles":[]` +
				`}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := jsonld.Marshal(test.Value)
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
