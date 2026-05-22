package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestAnyInterest_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"as":"https://www.w3.org/ns/activitystreams"` +
		`,` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"name":"as:name"` +

		`,` +

		`"id":"cv:id"` +
		`,` +
		`"keywords":"cv:keywords"` +
		`,` +
		`"type":"cv:type"` +
	`}`

	tests := []struct {
		Value      jsonresume.AnyInterest
		Name     activitypub.Name
		Expected string
	}{
		// 0
		{
			Value: jsonresume.AnyInterest{
				ID:   jsonld.SomeID("http://example.com/resume/interest/photography"),
				Type: jsonld.SomeType("Interest"),
				CoreInterest: jsonresume.CoreInterest{
					Keywords: activitypub.SomeStrings("nature", "landscape"),
				},
			},
			Name: activitypub.SomeName("Photography"),
			Expected: `{` +
				context +
				`,` +
				`"name":"Photography"` +
				`,` +
				`"id":"http://example.com/resume/interest/photography"` +
				`,` +
				`"type":"Interest"` +
				`,` +
				`"keywords":["nature","landscape"]` +
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
