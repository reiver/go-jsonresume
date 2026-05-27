package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestEducation_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +
		`,` +
		`"as":"http://www.w3.org/ns/activitystreams"` +

		`,` +

		`"area":"cv:area"` +
		`,` +
		`"courses":"cv:courses"` +
		`,` +
		`"endDate":"cv:endDate"` +
		`,` +
		`"id":"cv:id"` +
		`,` +
		`"institution":"cv:institution"` +
		`,` +
		`"score":"cv:score"` +
		`,` +
		`"startDate":"cv:startDate"` +
		`,` +
		`"studyType":"cv:studyType"` +
		`,` +
		`"type":"cv:type"` +
		`,` +
		`"url":"as:url"` +
	`}`

	tests := []struct {
		Value    jsonresume.Education
		Expected string
	}{
		// 0
		{
			Value: jsonresume.Education{
				ID: jsonld.SomeID("http://example.com/resume/education/mit"),
				CoreEducation: jsonresume.CoreEducation{
					Institution: nul.Something("MIT"),
					Area:        activitypub.SomeString("Computer Science"),
					StudyType:   activitypub.SomeString("Bachelor"),
					StartDate:   nul.Something("2010-09-01"),
					EndDate:     nul.Something("2014-06-15"),
					Score:       nul.Something("3.9"),
				},
			},
			Expected: `{` +
				context +
				`,` +
				`"id":"http://example.com/resume/education/mit"` +
				`,` +
				`"type":"Education"` +
				`,` +
				`"area":"Computer Science"` +
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
