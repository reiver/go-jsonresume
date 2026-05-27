package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestLocation_marshalJSONLD(t *testing.T) {

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
		Value    jsonresume.Location
		Expected string
	}{
		// 0
		{
			Value: jsonresume.Location{
				ID: jsonld.SomeID("http://example.com/resume/location/vancouver"),
				CoreLocation: jsonresume.CoreLocation{
					Address:     nul.Something("1234 Main Street"),
					City:        nul.Something("Vancouver"),
					CountryCode: nul.Something("CA"),
					PostalCode:  nul.Something("H0H 0H0"),
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
				`"address":"1234 Main Street"` +
				`,` +
				`"city":"Vancouver"` +
				`,` +
				`"countryCode":"CA"` +
				`,` +
				`"postalCode":"H0H 0H0"` +
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
