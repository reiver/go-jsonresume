package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreLanguage_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"fluency":"cv:fluency"` +
		`,` +
		`"language":"cv:language"` +
	`}`

	tests := []struct {
		Value    jsonresume.CoreLanguage
		Expected string
	}{
		// 0
		{
			Value: jsonresume.CoreLanguage{
				Fluency:  nul.Something("Native"),
				Language: nul.Something("English"),
			},
			Expected: `{` +
				context +
				`,` +
				`"fluency":"Native"` +
				`,` +
				`"language":"English"` +
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
