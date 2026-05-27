package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestReference_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"as":"https://www.w3.org/ns/activitystreams"` +
		`,` +
		`"cv":"https://w3id.org/fep/6158"` +
		`,` +
		`"as":"http://www.w3.org/ns/activitystreams"` +

		`,` +

		`"name":"as:name"` +

		`,` +

		`"id":"cv:id"` +
		`,` +
		`"reference":"cv:reference"` +
		`,` +
		`"type":"cv:type"` +
		`,` +
		`"name":"as:name"` +
	`}`

	tests := []struct {
		Value    jsonresume.Reference
		Name     activitypub.Name
		Expected string
	}{
		// 0
		{
			Value: jsonresume.Reference{
				ID: jsonld.SomeID("http://example.com/resume/reference/jane-doe"),
				CoreReference: jsonresume.CoreReference{
					Reference: nul.Something("Joe was great to work with."),
				},
			},
			Name: activitypub.SomeName("Jane Doe"),
			Expected: `{` +
				context +
				`,` +
				`"name":"Jane Doe"` +
				`,` +
				`"id":"http://example.com/resume/reference/jane-doe"` +
				`,` +
				`"type":"Reference"` +
				`,` +
				`"reference":"Joe was great to work with."` +
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
