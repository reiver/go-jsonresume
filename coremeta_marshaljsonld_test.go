package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreMeta_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"canonical":"cv:canonical"` +
		`,` +
		`"lastModified":"cv:lastModified"` +
		`,` +
		`"version":"cv:version"` +
	`}`

	tests := []struct {
		Value    jsonresume.CoreMeta
		Expected string
	}{
		// 0
		{
			Value: jsonresume.CoreMeta{
				Canonical:    nul.Null[string](),
				LastModified: nul.Null[string](),
				Version:      nul.Null[string](),
			},
			Expected: `{` +
				context +
				`}`,
		},

		// 1
		{
			Value: jsonresume.CoreMeta{
				Canonical:    nul.Something("https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json"),
				LastModified: nul.Something("2017-12-24T15:53:00"),
				Version:      nul.Something("v1.0.0"),
			},
			Expected: `{` +
				context +
				`,` +
				`"canonical":"https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json"` +
				`,` +
				`"lastModified":"2017-12-24T15:53:00"` +
				`,` +
				`"version":"v1.0.0"` +
				`}`,
		},

		// 2
		{
			Value: jsonresume.CoreMeta{
				Canonical:    nul.Something("https://example.com/resume.json"),
				LastModified: nul.Null[string](),
				Version:      nul.Something("v2.0.0"),
			},
			Expected: `{` +
				context +
				`,` +
				`"canonical":"https://example.com/resume.json"` +
				`,` +
				`"version":"v2.0.0"` +
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
