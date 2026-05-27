package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestMeta_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"canonical":"cv:canonical"` +
		`,` +
		`"id":"cv:id"` +
		`,` +
		`"lastModified":"cv:lastModified"` +
		`,` +
		`"type":"cv:type"` +
		`,` +
		`"version":"cv:version"` +
	`}`

	tests := []struct {
		Value    jsonresume.Meta
		Expected string
	}{
		// 0: minimal
		{
			Value: jsonresume.Meta{},
			Expected: `{` +
				context +
				`,` +
				`"type":"Meta"` +
				`}`,
		},

		// 1: with ID and fields
		{
			Value: jsonresume.Meta{
				ID: jsonld.SomeID("http://example.com/resume/meta"),
				CoreMeta: jsonresume.CoreMeta{
					Canonical:    nul.Something("https://example.com/resume.json"),
					LastModified: nul.Something("2017-12-24T15:53:00"),
					Version:      nul.Something("v1.0.0"),
				},
			},
			Expected: `{` +
				context +
				`,` +
				`"id":"http://example.com/resume/meta"` +
				`,` +
				`"type":"Meta"` +
				`,` +
				`"canonical":"https://example.com/resume.json"` +
				`,` +
				`"lastModified":"2017-12-24T15:53:00"` +
				`,` +
				`"version":"v1.0.0"` +
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
