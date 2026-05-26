package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestAnyCertificate_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"as":"https://www.w3.org/ns/activitystreams"` +
		`,` +
		`"cv":"https://w3id.org/fep/6158"` +
		`,` +
		`"as":"http://www.w3.org/ns/activitystreams"` +

		`,` +

		`"name":"as:name"` +

		`,` +

		`"date":"cv:date"` +
		`,` +
		`"id":"cv:id"` +
		`,` +
		`"issuer":"cv:issuer"` +
		`,` +
		`"type":"cv:type"` +
		`,` +
		`"name":"as:name"` +
		`,` +
		`"url":"as:url"` +
	`}`

	tests := []struct {
		Value    jsonresume.AnyCertificate
		Name     activitypub.Name
		Expected string
	}{
		// 0
		{
			Value: jsonresume.AnyCertificate{
				ID:   jsonld.SomeID("http://example.com/resume/certificate/aws"),
				Type: jsonld.SomeType("Certificate"),
				CoreCertificate: jsonresume.CoreCertificate{
					Date:   nul.Something("2023-06-01"),
					Issuer: nul.Something("Amazon"),
				},
			},
			Name: activitypub.SomeName("AWS Solutions Architect"),
			Expected: `{` +
				context +
				`,` +
				`"name":"AWS Solutions Architect"` +
				`,` +
				`"id":"http://example.com/resume/certificate/aws"` +
				`,` +
				`"type":"Certificate"` +
				`,` +
				`"date":"2023-06-01"` +
				`,` +
				`"issuer":"Amazon"` +
				`,` +
				`"url ":[]` +
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
