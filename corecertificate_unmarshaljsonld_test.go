package jsonresume_test

import (
	"github.com/reiver/go-jsonld"
	"reflect"
	"testing"

	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreCertificate_unmarshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"date":"cv:date"` +
		`,` +
		`"issuer":"cv:issuer"` +
	`}`

	tests := []struct {
		JSON     string
		Expected jsonresume.CoreCertificate
	}{
		// 0
		{
			JSON: `{` +
				context +
				`,` +
				`"date":null` +
				`,` +
				`"issuer":null` +
				`}`,
			Expected: jsonresume.CoreCertificate{
				Date:   nul.Null[string](),
				Issuer: nul.Null[string](),
			},
		},

		// 1
		{
			JSON: `{` +
				context +
				`,` +
				`"date":"2023-06-01"` +
				`,` +
				`"issuer":"Amazon"` +
				`}`,
			Expected: jsonresume.CoreCertificate{
				Date:   nul.Something("2023-06-01"),
				Issuer: nul.Something("Amazon"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CoreCertificate

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
