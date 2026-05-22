package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestAnyLanguage_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"fluency":"cv:fluency"` +
		`,` +
		`"id":"cv:id"` +
		`,` +
		`"language":"cv:language"` +
		`,` +
		`"type":"cv:type"` +
	`}`

	tests := []struct {
		Value    jsonresume.AnyLanguage
		Expected string
	}{
		// 0
		{
			Value: jsonresume.AnyLanguage{
				ID:   jsonld.SomeID("http://example.com/resume/language/english"),
				Type: jsonld.SomeType("Language"),
				CoreLanguage: jsonresume.CoreLanguage{
					Fluency:  nul.Something("Native"),
					Language: nul.Something("English"),
				},
			},
			Expected: `{` +
				context +
				`,` +
				`"id":"http://example.com/resume/language/english"` +
				`,` +
				`"type":"Language"` +
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
