package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"reflect"

	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCorePublication_unmarshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"publisher":"cv:publisher"` +
		`,` +
		`"releaseDate":"cv:releaseDate"` +
	`}`

	tests := []struct {
		JSON     string
		Expected jsonresume.CorePublication
	}{
		// 0
		{
			JSON: `{` +
				context +
				`,` +
				`"publisher":"O'Reilly"` +
				`,` +
				`"releaseDate":"2022-03-15"` +
				`}`,
			Expected: jsonresume.CorePublication{
				Publisher:   nul.Something("O'Reilly"),
				ReleaseDate: nul.Something("2022-03-15"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CorePublication

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
