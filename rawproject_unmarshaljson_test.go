package jsonresume

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestRawProject_unmarshalJSON(t *testing.T) {
	tests := []struct{
		JSON []byte
		Expected rawProject
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
					`"description":null` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":false` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":true` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":-2` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":-1` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":0` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":1` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":2` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":""` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":"hello world"` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"endDate":null` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":false` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":true` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":-2` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":-1` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":0` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":1` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":2` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":""` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":"hello world"` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProject{
				EndDate: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"entity":null` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":false` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":true` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":-2` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":-1` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":0` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":1` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":2` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":""` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":"hello world"` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"entity":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProject{
				Entity: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"highlights":null` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":false` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":true` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":-2` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":-1` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":0` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":1` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":2` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":""` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":"hello world"` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProject{
				Highlights: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"keywords":null` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":false` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":true` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":-2` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":-1` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":0` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":1` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":2` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":""` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":"hello world"` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"keywords":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProject{
				Keywords: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"name":null` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":false` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":true` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":-2` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":-1` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":0` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":1` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":2` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":""` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"hello world"` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProject{
				Name: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"roles":null` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":false` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":true` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":-2` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":-1` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":0` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":1` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":2` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":""` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":"hello world"` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"roles":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProject{
				Roles: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"startDate":null` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":false` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":true` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":-2` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":-1` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":0` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":1` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":2` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":""` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":"hello world"` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProject{
				StartDate: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"url":null` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":false` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":true` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-2` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-1` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":0` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":1` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":2` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":""` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"hello world"` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawProject{
				URL: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"description":"ONE"` +
					`,` +
					`"endDate":"TWO"` +
					`,` +
					`"entity":"THREE"` +
					`,` +
					`"highlights":"FOUR"` +
					`,` +
					`"keywords":"FIVE"` +
					`,` +
					`"name":"SIX"` +
					`,` +
					`"roles":"SEVEN"` +
					`,` +
					`"startDate":"EIGHT"` +
					`,` +
					`"url":"NINE"` +
				`}`,
			),
			Expected: rawProject{
				Description: []byte(`"ONE"`),
				EndDate:     []byte(`"TWO"`),
				Entity:      []byte(`"THREE"`),
				Highlights:  []byte(`"FOUR"`),
				Keywords:    []byte(`"FIVE"`),
				Name:        []byte(`"SIX"`),
				Roles:       []byte(`"SEVEN"`),
				StartDate:   []byte(`"EIGHT"`),
				URL:         []byte(`"NINE"`),
			},
		},
	}

	for testNumber, test := range tests {
		var actual rawProject

		err := json.Unmarshal(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual raw-project is not what was expected.", testNumber)

			t.Logf("EXPECTED.DESCRIPTION:\n%s", expected.Description)
			t.Logf("EXPECTED.ENDDATE:\n%s", expected.EndDate)
			t.Logf("EXPECTED.ENTITY:\n%s", expected.Entity)
			t.Logf("EXPECTED.HIGHLIGHTS:\n%s", expected.Highlights)
			t.Logf("EXPECTED.KEYWORDS:\n%s", expected.Keywords)
			t.Logf("EXPECTED.NAME:\n%s", expected.Name)
			t.Logf("EXPECTED.ROLES:\n%s", expected.Roles)
			t.Logf("EXPECTED.STARTDATE:\n%s", expected.StartDate)
			t.Logf("EXPECTED.URL:\n%s", expected.URL)

			t.Logf("ACTUAL.DESCRIPTION:\n%s", actual.Description)
			t.Logf("ACTUAL.ENDDATE:\n%s", actual.EndDate)
			t.Logf("ACTUAL.ENTITY:\n%s", actual.Entity)
			t.Logf("ACTUAL.HIGHLIGHTS:\n%s", actual.Highlights)
			t.Logf("ACTUAL.KEYWORDS:\n%s", actual.Keywords)
			t.Logf("ACTUAL.NAME:\n%s", actual.Name)
			t.Logf("ACTUAL.ROLES:\n%s", actual.Roles)
			t.Logf("ACTUAL.STARTDATE:\n%s", actual.StartDate)
			t.Logf("ACTUAL.URL:\n%s", actual.URL)

			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
