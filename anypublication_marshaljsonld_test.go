package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestAnyPublication_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"as":"https://www.w3.org/ns/activitystreams"` +
		`,` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"name":"as:name"` +

		`,` +

		`"id":"cv:id"` +
		`,` +
		`"publisher":"cv:publisher"` +
		`,` +
		`"releaseDate":"cv:releaseDate"` +
		`,` +
		`"type":"cv:type"` +
	`}`

	tests := []struct {
		Value    jsonresume.AnyPublication
		Name     activitypub.Name
		Expected string
	}{
		// 0
		{
			Value: jsonresume.AnyPublication{
				ID:   jsonld.SomeID("http://example.com/resume/publication/distributed-systems"),
				Type: jsonld.SomeType("Publication"),
				CorePublication: jsonresume.CorePublication{
					Publisher:   nul.Something("O'Reilly"),
					ReleaseDate: nul.Something("2022-03-15"),
				},
			},
			Name: activitypub.SomeName("Distributed Systems"),
			Expected: `{` +
				context +
				`,` +
				`"name":"Distributed Systems"` +
				`,` +
				`"id":"http://example.com/resume/publication/distributed-systems"` +
				`,` +
				`"type":"Publication"` +
				`,` +
				`"publisher":"O'Reilly"` +
				`,` +
				`"releaseDate":"2022-03-15"` +
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
