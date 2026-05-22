package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestAnyProfile_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"id":"cv:id"` +
		`,` +
		`"network":"cv:network"` +
		`,` +
		`"type":"cv:type"` +
		`,` +
		`"username":"cv:username"` +
	`}`

	tests := []struct {
		Value    jsonresume.AnyProfile
		Expected string
	}{
		// 0
		{
			Value: jsonresume.AnyProfile{
				ID:   jsonld.SomeID("http://example.com/resume/profile/mastodon"),
				Type: jsonld.SomeType("Profile"),
				CoreProfile: jsonresume.CoreProfile{
					Network:  nul.Something("Mastodon"),
					UserName: nul.Something("@joe@example.com"),
				},
			},
			Expected: `{` +
				context +
				`,` +
				`"id":"http://example.com/resume/profile/mastodon"` +
				`,` +
				`"type":"Profile"` +
				`,` +
				`"network":"Mastodon"` +
				`,` +
				`"username":"@joe@example.com"` +
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
