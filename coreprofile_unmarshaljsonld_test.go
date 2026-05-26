package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"reflect"

	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestCoreProfile_unmarshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"network":"cv:network"` +
		`,` +
		`"username":"cv:username"` +
	`}`

	tests := []struct {
		JSON     string
		Expected jsonresume.CoreProfile
	}{
		// 0
		{
			JSON: `{` +
				context +
				`,` +
				`"network":"Mastodon"` +
				`,` +
				`"username":"@joe@example.com"` +
				`}`,
			Expected: jsonresume.CoreProfile{
				Network:  nul.Something("Mastodon"),
				UserName: nul.Something("@joe@example.com"),
			},
		},
	}

	for testNumber, test := range tests {

		var actual jsonresume.CoreProfile

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
