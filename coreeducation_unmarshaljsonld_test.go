package jsonresume_test

import (
	"encoding/json"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreEducation_unmarshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"area":"cv:area"` +
		`,` +
		`"courses":"cv:courses"` +
		`,` +
		`"endDate":"cv:endDate"` +
		`,` +
		`"institution":"cv:institution"` +
		`,` +
		`"score":"cv:score"` +
		`,` +
		`"startDate":"cv:startDate"` +
		`,` +
		`"studyType":"cv:studyType"` +
	`}`

	tests := []struct {
		JSON     string
		Expected jsonresume.CoreEducation
	}{
		// 0
		{
			JSON: `{` +
				context +
				`,` +
				`"area":"Computer Science"` +
				`,` +
				`"courses":null` +
				`,` +
				`"endDate":"2014-06-15"` +
				`,` +
				`"institution":"MIT"` +
				`,` +
				`"score":"3.9"` +
				`,` +
				`"startDate":"2010-09-01"` +
				`,` +
				`"studyType":"Bachelor"` +
				`}`,
			Expected: jsonresume.CoreEducation{
				Area:        activitypub.SomeString("Computer Science"),
				EndDate:     nul.Something("2014-06-15"),
				Institution: nul.Something("MIT"),
				Score:       nul.Something("3.9"),
				StartDate:   nul.Something("2010-09-01"),
				StudyType:   activitypub.SomeString("Bachelor"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CoreEducation

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
