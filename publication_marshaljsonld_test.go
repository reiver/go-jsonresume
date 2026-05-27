package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestPublication_marshalJSONLD(t *testing.T) {

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
		`"publisher":"cv:publisher"` +
		`,` +
		`"releaseDate":"cv:releaseDate"` +
		`,` +
		`"type":"cv:type"` +
		`,` +
		`"name":"as:name"` +
		`,` +
		`"summary":"as:summary"` +
		`,` +
		`"url":"as:url"` +
	`}`

	tests := []struct {
		Value    jsonresume.Publication
		Name     activitypub.Name
		Expected string
	}{
		// 0
		{
			Value: jsonresume.Publication{
				ID: jsonld.SomeID("http://example.com/resume/publication/distributed-systems"),
				CorePublication: jsonresume.CorePublication{
					Publisher:   nul.Something("BookCo"),
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
				`"publisher":"BookCo"` +
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
