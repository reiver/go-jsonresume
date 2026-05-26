package jsonresume_test

import (
	"github.com/reiver/go-jsonld"
	"testing"

	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreAward_unmarshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"awarder":"cv:awarder"` +
		`,` +
		`"date":"cv:date"` +
		`,` +
		`"title":"cv:title"` +
	`}`

	tests := []struct {
		JSON     string
		Expected jsonresume.CoreAward
	}{
		// 0
		{
			JSON: `{` +
				context +
				`,` +
				`"awarder":null` +
				`,` +
				`"date":null` +
				`,` +
				`"title":null` +
				`}`,
			Expected: jsonresume.CoreAward{
				Awarder: nul.Null[string](),
				Date:    nul.Null[string](),
				Title:   nul.Null[string](),
			},
		},

		// 1
		{
			JSON: `{` +
				context +
				`,` +
				`"awarder":"ACME Corp"` +
				`,` +
				`"date":"2024-01-15"` +
				`,` +
				`"title":"Best Employee"` +
				`}`,
			Expected: jsonresume.CoreAward{
				Awarder: nul.Something("ACME Corp"),
				Date:    nul.Something("2024-01-15"),
				Title:   nul.Something("Best Employee"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CoreAward

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
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
