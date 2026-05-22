package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestAnyAward_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"awarder":"cv:awarder"` +
		`,` +
		`"date":"cv:date"` +
		`,` +
		`"id":"cv:id"` +
		`,` +
		`"title":"cv:title"` +
		`,` +
		`"type":"cv:type"` +
	`}`

	tests := []struct {
		Value    jsonresume.AnyAward
		Expected string
	}{
		// 0
		{
			Value: jsonresume.AnyAward{
				Type: jsonld.SomeType("Award"),
				CoreAward: jsonresume.CoreAward{
					Awarder: nul.Null[string](),
					Date:    nul.Null[string](),
					Title:   nul.Null[string](),
				},
			},
			Expected: `{` +
				context +
				`,` +
				`"type":"Award"` +
				`,` +
				`"awarder":null` +
				`,` +
				`"date":null` +
				`,` +
				`"title":null` +
				`}`,
		},
		// 1
		{
			Value: jsonresume.AnyAward{
				ID:   jsonld.SomeID("http://example.com/resume/award/best-employee-2024"),
				Type: jsonld.SomeType("Award"),
				CoreAward: jsonresume.CoreAward{
					Awarder: nul.Something("ACME Corp"),
					Date:    nul.Something("2024-01-15"),
					Title:   nul.Something("Best Employee"),
				},
			},
			Expected: `{` +
				context +
				`,` +
				`"id":"http://example.com/resume/award/best-employee-2024"` +
				`,` +
				`"type":"Award"` +
				`,` +
				`"awarder":"ACME Corp"` +
				`,` +
				`"date":"2024-01-15"` +
				`,` +
				`"title":"Best Employee"` +
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
