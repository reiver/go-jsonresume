package jsonresume_test

import (
	"encoding/json"
	"testing"

	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreLanguage_unmarshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"fluency":"cv:fluency"` +
		`,` +
		`"language":"cv:language"` +
	`}`

	tests := []struct {
		JSON     string
		Expected jsonresume.CoreLanguage
	}{
		// 0
		{
			JSON: `{` +
				context +
				`,` +
				`"fluency":"Native"` +
				`,` +
				`"language":"English"` +
				`}`,
			Expected: jsonresume.CoreLanguage{
				Fluency:  nul.Something("Native"),
				Language: nul.Something("English"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CoreLanguage

		err := json.Unmarshal([]byte(test.JSON), &actual)
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
