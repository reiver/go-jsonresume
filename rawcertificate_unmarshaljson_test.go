package jsonresume

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestRawCertificate_unmarshalJSON(t *testing.T) {
	tests := []struct{
		JSON []byte
		Expected rawCertificate
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
					`"date":null` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":false` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":true` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":-2` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":-1` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":0` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":1` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":2` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":""` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":"hello world"` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"date":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawCertificate{
				Date: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"name":null` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":false` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":true` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":-2` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":-1` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":0` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":1` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":2` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":""` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"hello world"` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawCertificate{
				Name: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"issuer":null` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":false` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":true` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":-2` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":-1` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":0` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":1` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":2` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":""` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":"hello world"` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"issuer":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawCertificate{
				Issuer: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"url":null` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":false` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":true` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-2` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-1` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":0` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":1` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":2` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":""` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"hello world"` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawCertificate{
				URL: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"date":"ONE"` +
					`,` +
					`"name":"TWO"` +
					`,` +
					`"issuer":"THREE"` +
					`,` +
					`"url":"FOUR"` +
				`}`,
			),
			Expected: rawCertificate{
				Date:   []byte(`"ONE"`),
				Name:   []byte(`"TWO"`),
				Issuer: []byte(`"THREE"`),
				URL:    []byte(`"FOUR"`),
			},
		},
	}

	for testNumber, test := range tests {
		var actual rawCertificate

		err := json.Unmarshal(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual raw-certificate is not what was expected.", testNumber)

			t.Logf("EXPECTED.DATE:\n%s", expected.Date)
			t.Logf("EXPECTED.NAME:\n%s", expected.Name)
			t.Logf("EXPECTED.ISSUER:\n%s", expected.Issuer)
			t.Logf("EXPECTED.URL:\n%s", expected.URL)

			t.Logf("ACTUAL.DATE:\n%s", actual.Date)
			t.Logf("ACTUAL.NAME:\n%s", actual.Name)
			t.Logf("ACTUAL.ISSUER:\n%s", actual.Issuer)
			t.Logf("ACTUAL.URL:\n%s", actual.URL)

			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
