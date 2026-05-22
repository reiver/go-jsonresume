package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestAnyLocation_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"address":"cv:address"` +
		`,` +
		`"city":"cv:city"` +
		`,` +
		`"countryCode":"cv:countryCode"` +
		`,` +
		`"id":"cv:id"` +
		`,` +
		`"postalCode":"cv:postalCode"` +
		`,` +
		`"region":"cv:region"` +
		`,` +
		`"type":"cv:type"` +
	`}`

	tests := []struct {
		Value    jsonresume.AnyLocation
		Expected string
	}{
		// 0
		{
			Value: jsonresume.AnyLocation{
				ID:   jsonld.SomeID("http://example.com/resume/location/vancouver"),
				Type: jsonld.SomeType("Location"),
				CoreLocation: jsonresume.CoreLocation{
					Address:     nul.Something("123 Main St"),
					City:        nul.Something("Vancouver"),
					CountryCode: nul.Something("CA"),
					PostalCode:  nul.Something("V5K 0A1"),
					Region:      nul.Something("British Columbia"),
				},
			},
			Expected: `{` +
				context +
				`,` +
				`"id":"http://example.com/resume/location/vancouver"` +
				`,` +
				`"type":"Location"` +
				`,` +
				`"address":"123 Main St"` +
				`,` +
				`"city":"Vancouver"` +
				`,` +
				`"countryCode":"CA"` +
				`,` +
				`"postalCode":"V5K 0A1"` +
				`,` +
				`"region":"British Columbia"` +
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
