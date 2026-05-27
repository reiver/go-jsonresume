package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestExperience_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"as":"https://www.w3.org/ns/activitystreams"` +
		`,` +
		`"cv":"https://w3id.org/fep/6158"` +
		`,` +
		`"as":"http://www.w3.org/ns/activitystreams"` +

		`,` +

		`"name":"as:name"` +

		`,` +

		`"description":"cv:description"` +
		`,` +
		`"endDate":"cv:endDate"` +
		`,` +
		`"highlights":"cv:highlights"` +
		`,` +
		`"id":"cv:id"` +
		`,` +
		`"location":"cv:location"` +
		`,` +
		`"organization":"cv:organization"` +
		`,` +
		`"position":"cv:position"` +
		`,` +
		`"startDate":"cv:startDate"` +
		`,` +
		`"type":"cv:type"` +
		`,` +
		`"name":"as:name"` +
		`,` +
		`"summary":"as:summary"` +
		`,` +
		`"url":"as:url"` +
	`}`

	tests := []struct {
		Value    jsonresume.Experience
		Name     activitypub.Name
		Expected string
	}{
		// 0
		{
			Value: jsonresume.Experience{
				ID: jsonld.SomeID("http://example.com/resume/experience/3"),
				CoreExperience: jsonresume.CoreExperience{
					Position:     activitypub.SomeString("Software Engineer"),
					StartDate:    nul.Something("2020-01-01"),
					EndDate:      nul.Null[string](),
					Organization: nul.Null[string](),
					Highlights:   activitypub.SomeStrings("built backend", "led team"),
				},
			},
			Name: activitypub.SomeName("ACME Corp"),
			Expected: `{` +
				context +
				`,` +
				`"name":"ACME Corp"` +
				`,` +
				`"id":"http://example.com/resume/experience/3"` +
				`,` +
				`"type":"Experience"` +
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

		actualBytes, err := jsonld.Marshal(test.Name, test.Value)
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
