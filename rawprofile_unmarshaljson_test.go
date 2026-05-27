package jsonresume

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestRawProfile_unmarshalJSON(t *testing.T) {
	tests := []struct{
		JSON []byte
		Expected rawProfile
	}{
		{
			JSON: []byte(`null`),
		},



		{
			JSON: []byte(`{}`),
		},



		{
			JSON: []byte(`{"apple":"one","banana":2,"cherry":[1,2,3]}`),
		},



		{
			JSON: []byte(
				`{` +
					`"network":null` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":false` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":true` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":-2` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":-1` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":0` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":1` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":2` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":""` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":"hello world"` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"network":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProfile{
				Network: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"username":null` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":false` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":true` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":-2` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":-1` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":0` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":1` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":2` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":""` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":"hello world"` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"username":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProfile{
				UserName: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"url":null` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":false` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":true` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-2` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-1` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":0` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":1` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":2` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":""` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"hello world"` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProfile{
				URL: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"network":"ONE"` +
					`,` +
					`"username":"TWO"` +
					`,` +
					`"url":"THREE"` +
				`}`,
			),
			Expected: rawProfile{
				Network:  []byte(`"ONE"`),
				UserName: []byte(`"TWO"`),
				URL:      []byte(`"THREE"`),
			},
		},
	}

	for testNumber, test := range tests {
		var actual rawProfile

		err := json.Unmarshal(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual raw-profile is not what was expected.", testNumber)

			t.Logf("EXPECTED.NETWORK:\n%s", expected.Network)
			t.Logf("EXPECTED.USERNAME:\n%s", expected.UserName)
			t.Logf("EXPECTED.URL:\n%s", expected.URL)

			t.Logf("ACTUAL.NETWORK:\n%s", actual.Network)
			t.Logf("ACTUAL.USERNAME:\n%s", actual.UserName)
			t.Logf("ACTUAL.URL:\n%s", actual.URL)

			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
