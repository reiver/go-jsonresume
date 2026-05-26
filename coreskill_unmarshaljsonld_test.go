package jsonresume_test

import (
	"github.com/reiver/go-jsonld"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreSkill_unmarshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"keywords":"cv:keywords"` +
		`,` +
		`"level":"cv:level"` +
	`}`

	tests := []struct {
		JSON     string
		Expected jsonresume.CoreSkill
	}{
		// 0
		{
			JSON: `{` +
				context +
				`,` +
				`"keywords":null` +
				`,` +
				`"level":null` +
				`}`,
			Expected: jsonresume.CoreSkill{
				Level: nul.Null[string](),
			},
		},

		// 1
		{
			JSON: `{` +
				context +
				`,` +
				`"keywords":["backend","concurrency"]` +
				`,` +
				`"level":"Advanced"` +
				`}`,
			Expected: jsonresume.CoreSkill{
				Keywords: activitypub.SomeStrings("backend","concurrency"),
				Level:    nul.Something("Advanced"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CoreSkill

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
