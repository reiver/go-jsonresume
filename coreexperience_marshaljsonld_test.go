package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreExperience_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +
		`,` +
		`"as":"http://www.w3.org/ns/activitystreams"` +

		`,` +

		`"description":"cv:description"` +
		`,` +
		`"endDate":"cv:endDate"` +
		`,` +
		`"highlights":"cv:highlights"` +
		`,` +
		`"location":"cv:location"` +
		`,` +
		`"organization":"cv:organization"` +
		`,` +
		`"position":"cv:position"` +
		`,` +
		`"startDate":"cv:startDate"` +
		`,` +
		`"name":"as:name"` +
		`,` +
		`"summary":"as:summary"` +
		`,` +
		`"url":"as:url"` +
	`}`

	tests := []struct {
		Value    jsonresume.CoreExperience
		Expected string
	}{
		// 0
		{
			Value: jsonresume.CoreExperience{
				EndDate:      nul.Null[string](),
				Highlights:   activitypub.SomeStrings("built backend","led team"),
				Organization: nul.Null[string](),
				Position:     activitypub.SomeString("Software Engineer"),
				StartDate:    nul.Something("2020-01-01"),
			},
			Expected: `{` +
				context +
				`,` +
				`"highlights":["built backend","led team"]` +
				`,` +
				`"position":"Software Engineer"` +
				`,` +
				`"startDate":"2020-01-01"` +
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
