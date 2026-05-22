package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestAnyResume_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"as":"https://www.w3.org/ns/activitystreams"` +
		`,` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"name":"as:name"` +

		`,` +

		`"awards":"cv:awards"` +
		`,` +
		`"basics":"cv:basics"` +
		`,` +
		`"certificates":"cv:certificates"` +
		`,` +
		`"education":"cv:education"` +
		`,` +
		`"id":"cv:id"` +
		`,` +
		`"interests":"cv:interests"` +
		`,` +
		`"languages":"cv:languages"` +
		`,` +
		`"projects":"cv:projects"` +
		`,` +
		`"publications":"cv:publications"` +
		`,` +
		`"references":"cv:references"` +
		`,` +
		`"skills":"cv:skills"` +
		`,` +
		`"type":"cv:type"` +
		`,` +
		`"volunteer":"cv:volunteer"` +
		`,` +
		`"work":"cv:work"` +
	`}`

	tests := []struct {
		Value    jsonresume.AnyResume
		Name     activitypub.Name
		Expected string
	}{
		// 0
		{
			Value: jsonresume.AnyResume{
				Type: jsonld.SomeType("Resume"),
			},
			Name: activitypub.SomeName("Joe Blow"),
			Expected: `{` +
				context +
				`,` +
				`"name":"Joe Blow"` +
				`,` +
				`"type":"Resume"` +
				`,` +
				`"awards":[]` +
				`,` +
				`"basics":null` +
				`,` +
				`"certificates":[]` +
				`,` +
				`"education":[]` +
				`,` +
				`"interests":[]` +
				`,` +
				`"languages":[]` +
				`,` +
				`"projects":[]` +
				`,` +
				`"publications":[]` +
				`,` +
				`"references":[]` +
				`,` +
				`"skills":[]` +
				`,` +
				`"volunteer":[]` +
				`,` +
				`"work":[]` +
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
