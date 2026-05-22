package jsonresume_test

import (
	"encoding/json"
	"testing"

	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreLocation_unmarshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"address":"cv:address"` +
		`,` +
		`"city":"cv:city"` +
		`,` +
		`"countryCode":"cv:countryCode"` +
		`,` +
		`"postalCode":"cv:postalCode"` +
		`,` +
		`"region":"cv:region"` +
	`}`

	tests := []struct {
		JSON     string
		Expected jsonresume.CoreLocation
	}{
		// 0
		{
			JSON: `{` +
				context +
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
			Expected: jsonresume.CoreLocation{
				Address:     nul.Something("123 Main St"),
				City:        nul.Something("Vancouver"),
				CountryCode: nul.Something("CA"),
				PostalCode:  nul.Something("V5K 0A1"),
				Region:      nul.Something("British Columbia"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CoreLocation

		err := json.Unmarshal([]byte(test.JSON), &actual)
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
