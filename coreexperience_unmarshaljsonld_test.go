package jsonresume_test

import (
	"encoding/json"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreExperience_unmarshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"endDate":"cv:endDate"` +
		`,` +
		`"highlights":"cv:highlights"` +
		`,` +
		`"organization":"cv:organization"` +
		`,` +
		`"position":"cv:position"` +
		`,` +
		`"startDate":"cv:startDate"` +
	`}`

	tests := []struct {
		JSON     string
		Expected jsonresume.CoreExperience
	}{
		// 0
		{
			JSON: `{` +
				context +
				`,` +
				`"endDate":null` +
				`,` +
				`"highlights":["built backend","led team"]` +
				`,` +
				`"organization":null` +
				`,` +
				`"position":"Software Engineer"` +
				`,` +
				`"startDate":"2020-01-01"` +
				`}`,
			Expected: jsonresume.CoreExperience{
				EndDate:      nul.Null[string](),
				Highlights:   activitypub.SomeStrings("built backend","led team"),
				Organization: nul.Null[string](),
				Position:     activitypub.SomeString("Software Engineer"),
				StartDate:    nul.Something("2020-01-01"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CoreExperience

		err := json.Unmarshal([]byte(test.JSON), &actual)
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
