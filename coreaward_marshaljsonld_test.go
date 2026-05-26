package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreAward_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +
		`,` +
		`"as":"http://www.w3.org/ns/activitystreams"` +

		`,` +

		`"awarder":"cv:awarder"` +
		`,` +
		`"date":"cv:date"` +
		`,` +
		`"title":"cv:title"` +
		`,` +
		`"summary":"as:summary"` +
	`}`

	tests := []struct {
		Value    jsonresume.CoreAward
		Expected string
	}{
		// 0
		{
			Value: jsonresume.CoreAward{
				Awarder: nul.Null[string](),
				Date:    nul.Null[string](),
				Title:   nul.Null[string](),
			},
			Expected: `{` +
				context +
				`}`,
		},

		// 1
		{
			Value: jsonresume.CoreAward{
				Awarder: nul.Something("ACME Corp"),
				Date:    nul.Something("2024-01-15"),
				Title:   nul.Something("Best Employee"),
			},
			Expected: `{` +
				context +
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
