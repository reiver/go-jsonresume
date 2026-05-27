package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreSkill_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +
		`,` +
		`"as":"http://www.w3.org/ns/activitystreams"` +

		`,` +

		`"keywords":"cv:keywords"` +
		`,` +
		`"level":"cv:level"` +
		`,` +
		`"name":"as:name"` +
	`}`

	tests := []struct {
		Value    jsonresume.CoreSkill
		Expected string
	}{
		// 0
		{
			Value: jsonresume.CoreSkill{
				Level: nul.Null[string](),
			},
			Expected: `{` +
				context +
				`}`,
		},

		// 1
		{
			Value: jsonresume.CoreSkill{
				Keywords: activitypub.SomeStrings("backend","concurrency"),
				Level:    nul.Something("Advanced"),
			},
			Expected: `{` +
				context +
				`,` +
				`"keywords":["backend","concurrency"]` +
				`,` +
				`"level":"Advanced"` +
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
