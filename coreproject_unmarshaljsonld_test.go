package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"reflect"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreProject_unmarshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

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
	`}`

	tests := []struct {
		JSON     string
		Expected jsonresume.CoreProject
	}{
		// 0
		{
			JSON: `{` +
				context +
				`,` +
				`"description":"An ActivityPub server"` +
				`,` +
				`"endDate":null` +
				`,` +
				`"entity":null` +
				`,` +
				`"highlights":"federation support"` +
				`,` +
				`"keywords":["activitypub","go"]` +
				`,` +
				`"roles":"lead developer"` +
				`,` +
				`"startDate":"2023-01-01"` +
				`}`,
			Expected: jsonresume.CoreProject{
				Description: nul.Something("An ActivityPub server"),
				EndDate:     nul.Null[string](),
				Entity:      nul.Null[string](),
				Highlights:  activitypub.SomeStrings("federation support"),
				Keywords:    activitypub.SomeStrings("activitypub","go"),
				Roles:       activitypub.SomeStrings("lead developer"),
				StartDate:   nul.Something("2023-01-01"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CoreProject

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
