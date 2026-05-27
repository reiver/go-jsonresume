package jsonresume

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestRawPublication_unmarshalJSON(t *testing.T) {
	tests := []struct{
		JSON []byte
		Expected rawPublication
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
					`"name":null` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":false` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":true` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":-2` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":-1` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":0` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":1` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":2` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":""` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"hello world"` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawPublication{
				Name: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"publisher":null` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":false` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":true` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":-2` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":-1` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":0` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":1` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":2` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":""` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":"hello world"` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publisher":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawPublication{
				Publisher: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"releaseDate":null` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":false` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":true` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":-2` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":-1` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":0` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":1` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":2` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":""` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":"hello world"` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"releaseDate":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawPublication{
				ReleaseDate: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"summary":null` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":false` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":true` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":-2` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":-1` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":0` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":1` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":2` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":""` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":"hello world"` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawPublication{
				Summary: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"url":null` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":false` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":true` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-2` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-1` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":0` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":1` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":2` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":""` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"hello world"` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawPublication{
				URL: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"name":"ONE"` +
					`,` +
					`"publisher":"TWO"` +
					`,` +
					`"releaseDate":"THREE"` +
					`,` +
					`"summary":"FOUR"` +
					`,` +
					`"url":"FIVE"` +
				`}`,
			),
			Expected: rawPublication{
				Name:        []byte(`"ONE"`),
				Publisher:   []byte(`"TWO"`),
				ReleaseDate: []byte(`"THREE"`),
				Summary:     []byte(`"FOUR"`),
				URL:         []byte(`"FIVE"`),
			},
		},
	}

	for testNumber, test := range tests {
		var actual rawPublication

		err := json.Unmarshal(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual raw-publication is not what was expected.", testNumber)

			t.Logf("EXPECTED.NAME:\n%s", expected.Name)
			t.Logf("EXPECTED.PUBLISHER:\n%s", expected.Publisher)
			t.Logf("EXPECTED.RELEASEDATE:\n%s", expected.ReleaseDate)
			t.Logf("EXPECTED.SUMMARY:\n%s", expected.Summary)
			t.Logf("EXPECTED.URL:\n%s", expected.URL)

			t.Logf("ACTUAL.NAME:\n%s", actual.Name)
			t.Logf("ACTUAL.PUBLISHER:\n%s", actual.Publisher)
			t.Logf("ACTUAL.RELEASEDATE:\n%s", actual.ReleaseDate)
			t.Logf("ACTUAL.SUMMARY:\n%s", actual.Summary)
			t.Logf("ACTUAL.URL:\n%s", actual.URL)

			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
