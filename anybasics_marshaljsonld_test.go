package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestAnyBasics_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"as":"https://www.w3.org/ns/activitystreams"` +
		`,` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"name":"as:name"` +

		`,` +

		`"email":"cv:email"` +
		`,` +
		`"id":"cv:id"` +
		`,` +
		`"label":"cv:label"` +
		`,` +
		`"phone":"cv:phone"` +
		`,` +
		`"profiles":"cv:profiles"` +
		`,` +
		`"type":"cv:type"` +
	`}`

	tests := []struct {
		Value    jsonresume.AnyBasics
		Name     activitypub.Name
		Expected string
	}{
		// 0
		{
			Value: jsonresume.AnyBasics{
				Type: jsonld.SomeType("Basics"),
			},
			Name: activitypub.SomeName("Joe Blow"),
			Expected: `{` +
				context +
				`,` +
				`"name":"Joe Blow"` +
				`,` +
				`"type":"Basics"` +
				`,` +
				`"email":null` +
				`,` +
				`"label":null` +
				`,` +
				`"phone":null` +
				`,` +
				`"profiles":[]` +
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
