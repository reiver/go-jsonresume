package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreCertificate_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"date":"cv:date"` +
		`,` +
		`"issuer":"cv:issuer"` +
	`}`

	tests := []struct {
		Value    jsonresume.CoreCertificate
		Expected string
	}{
		// 0
		{
			Value: jsonresume.CoreCertificate{
				Date:   nul.Null[string](),
				Issuer: nul.Null[string](),
			},
			Expected: `{` +
				context +
				`,` +
				`"date":null` +
				`,` +
				`"issuer":null` +
				`}`,
		},

		// 1
		{
			Value: jsonresume.CoreCertificate{
				Date:   nul.Something("2023-06-01"),
				Issuer: nul.Something("Amazon"),
			},
			Expected: `{` +
				context +
				`,` +
				`"date":"2023-06-01"` +
				`,` +
				`"issuer":"Amazon"` +
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
