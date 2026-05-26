package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreProject_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +
		`,` +
		`"as":"http://www.w3.org/ns/activitystreams"` +

		`,` +

		`"description":"cv:description"` +
		`,` +
		`"endDate":"cv:endDate"` +
		`,` +
		`"entity":"cv:entity"` +
		`,` +
		`"highlights":"cv:highlights"` +
		`,` +
		`"keywords":"cv:keywords"` +
		`,` +
		`"roles":"cv:roles"` +
		`,` +
		`"startDate":"cv:startDate"` +
		`,` +
		`"name":"as:name"` +
		`,` +
		`"url":"as:url"` +
	`}`

	tests := []struct {
		Value    jsonresume.CoreProject
		Expected string
	}{
		// 0
		{
			Value: jsonresume.CoreProject{
				Description: nul.Something("An ActivityPub server"),
				EndDate:     nul.Null[string](),
				Entity:      nul.Null[string](),
				Highlights:  activitypub.SomeStrings("federation support"),
				Keywords:    activitypub.SomeStrings("activitypub", "go"),
				Roles:       activitypub.SomeStrings("lead developer"),
				StartDate:   nul.Something("2023-01-01"),
			},
			Expected: `{` +
				context +
				`,` +
				`"description":"An ActivityPub server"` +
				`,` +
				`"highlights":"federation support"` +
				`,` +
				`"keywords":["activitypub","go"]` +
				`,` +
				`"roles":"lead developer"` +
				`,` +
				`"startDate":"2023-01-01"` +
				`,` +
				`"url":[]` +
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
