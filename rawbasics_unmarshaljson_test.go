package jsonresume

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestRawBasics_unmarshalJSON(t *testing.T) {
	tests := []struct{
		JSON []byte
		Expected rawBasics
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
					`"email":null` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":false` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":true` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":-2` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":-1` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":0` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":1` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":2` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":""` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":"hello world"` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"email":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawBasics{
				EMail: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"image":null` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":false` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":true` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":-2` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":-1` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":0` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":1` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":2` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":""` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":"hello world"` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"image":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawBasics{
				Image: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"label":null` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":false` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":true` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":-2` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":-1` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":0` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":1` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":2` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":""` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":"hello world"` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"label":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawBasics{
				Label: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"location":null` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":false` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":true` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":-2` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":-1` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":0` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":1` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":2` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":""` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":"hello world"` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawBasics{
				Location: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"name":null` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":false` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":true` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":-2` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":-1` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":0` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":1` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":2` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":""` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"hello world"` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawBasics{
				Name: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"phone":null` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":false` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":true` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":-2` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":-1` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":0` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":1` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":2` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":""` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":"hello world"` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"phone":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawBasics{
				Phone: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"profiles":null` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":false` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":true` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":-2` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":-1` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":0` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":1` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":2` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":""` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":"hello world"` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"profiles":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawBasics{
				Profiles: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"summary":null` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":false` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":true` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":-2` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":-1` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":0` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":1` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":2` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":""` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":"hello world"` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawBasics{
				Summary: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"url":null` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":false` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":true` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-2` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-1` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":0` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":1` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":2` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":""` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"hello world"` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawBasics{
				URL: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"email":"ONE"` +
					`,` +
					`"image":"TWO"` +
					`,` +
					`"label":"THREE"` +
					`,` +
					`"location":"FOUR"` +
					`,` +
					`"name":"FIVE"` +
					`,` +
					`"phone":"SIX"` +
					`,` +
					`"profiles":"SEVEN"` +
					`,` +
					`"summary":"EIGHT"` +
					`,` +
					`"url":"NINE"` +
				`}`,
			),
			Expected: rawBasics{
				EMail:    []byte(`"ONE"`),
				Image:    []byte(`"TWO"`),
				Label:    []byte(`"THREE"`),
				Location: []byte(`"FOUR"`),
				Name:     []byte(`"FIVE"`),
				Phone:    []byte(`"SIX"`),
				Profiles: []byte(`"SEVEN"`),
				Summary:  []byte(`"EIGHT"`),
				URL:      []byte(`"NINE"`),
			},
		},
	}

	for testNumber, test := range tests {
		var actual rawBasics

		err := json.Unmarshal(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual raw-basics is not what was expected.", testNumber)

			t.Logf("EXPECTED.EMAIL:\n%s", expected.EMail)
			t.Logf("EXPECTED.IMAGE:\n%s", expected.Image)
			t.Logf("EXPECTED.LABEL:\n%s", expected.Label)
			t.Logf("EXPECTED.LOCATION:\n%s", expected.Location)
			t.Logf("EXPECTED.NAME:\n%s", expected.Name)
			t.Logf("EXPECTED.PHONE:\n%s", expected.Phone)
			t.Logf("EXPECTED.PROFILES:\n%s", expected.Profiles)
			t.Logf("EXPECTED.SUMMARY:\n%s", expected.Summary)
			t.Logf("EXPECTED.URL:\n%s", expected.URL)

			t.Logf("ACTUAL.EMAIL:\n%s", actual.EMail)
			t.Logf("ACTUAL.IMAGE:\n%s", actual.Image)
			t.Logf("ACTUAL.LABEL:\n%s", actual.Label)
			t.Logf("ACTUAL.LOCATION:\n%s", actual.Location)
			t.Logf("ACTUAL.NAME:\n%s", actual.Name)
			t.Logf("ACTUAL.PHONE:\n%s", actual.Phone)
			t.Logf("ACTUAL.PROFILES:\n%s", actual.Profiles)
			t.Logf("ACTUAL.SUMMARY:\n%s", actual.Summary)
			t.Logf("ACTUAL.URL:\n%s", actual.URL)

			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
