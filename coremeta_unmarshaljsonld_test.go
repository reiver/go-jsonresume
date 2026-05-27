package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreMeta_unmarshalJSONLD(t *testing.T) {

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
		JSON     string
		Expected jsonresume.CoreMeta
	}{
		// 0
		{
			JSON: `{` +
				context +
				`,` +
				`"canonical":null` +
				`,` +
				`"lastModified":null` +
				`,` +
				`"version":null` +
				`}`,
			Expected: jsonresume.CoreMeta{
				Canonical:    nul.Null[string](),
				LastModified: nul.Null[string](),
				Version:      nul.Null[string](),
			},
		},

		// 1
		{
			JSON: `{` +
				context +
				`,` +
				`"canonical":"https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json"` +
				`,` +
				`"lastModified":"2017-12-24T15:53:00"` +
				`,` +
				`"version":"v1.0.0"` +
				`}`,
			Expected: jsonresume.CoreMeta{
				Canonical:    nul.Something("https://raw.githubusercontent.com/jsonresume/resume-schema/master/resume.json"),
				LastModified: nul.Something("2017-12-24T15:53:00"),
				Version:      nul.Something("v1.0.0"),
			},
		},

		// 2
		{
			JSON: `{` +
				context +
				`,` +
				`"canonical":"https://example.com/resume.json"` +
				`,` +
				`"version":"v2.0.0"` +
				`}`,
			Expected: jsonresume.CoreMeta{
				Canonical:    nul.Something("https://example.com/resume.json"),
				LastModified: nul.Nothing[string](),
				Version:      nul.Something("v2.0.0"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CoreMeta

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
