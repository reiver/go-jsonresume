package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestProject_marshalJSONLD(t *testing.T) {

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
		`"entity":"cv:entity"` +
		`,` +
		`"highlights":"cv:highlights"` +
		`,` +
		`"id":"cv:id"` +
		`,` +
		`"keywords":"cv:keywords"` +
		`,` +
		`"roles":"cv:roles"` +
		`,` +
		`"startDate":"cv:startDate"` +
		`,` +
		`"type":"cv:type"` +
		`,` +
		`"name":"as:name"` +
		`,` +
		`"url":"as:url"` +
	`}`

	tests := []struct {
		Value    jsonresume.Project
		Name     activitypub.Name
		Expected string
	}{
		// 0
		{
			Value: jsonresume.Project{
				ID: jsonld.SomeID("http://example.com/resume/project/microdon"),
				CoreProject: jsonresume.CoreProject{
					Description: nul.Something("An ActivityPub server"),
					StartDate:   nul.Something("2023-01-01"),
					EndDate:     nul.Null[string](),
					Entity:      nul.Null[string](),
					Highlights:  activitypub.SomeStrings("federation support"),
					Keywords:    activitypub.SomeStrings("activitypub", "go"),
					Roles:       activitypub.SomeStrings("lead developer"),
				},
			},
			Name: activitypub.SomeName("Microdon"),
			Expected: `{` +
				context +
				`,` +
				`"name":"Microdon"` +
				`,` +
				`"id":"http://example.com/resume/project/microdon"` +
				`,` +
				`"type":"Project"` +
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
