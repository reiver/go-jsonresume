package jsonresume

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestRawExperience_unmarshalJSON(t *testing.T) {
	tests := []struct{
		JSON []byte
		Expected rawExperience
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
			Expected: rawExperience{
				Description: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":false` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":true` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":-2` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":-1` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":0` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":1` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":2` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":""` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":"hello world"` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"description":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawExperience{
				Description: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"endDate":null` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":false` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":true` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":-2` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":-1` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":0` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":1` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":2` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":""` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":"hello world"` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawExperience{
				EndDate: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"highlights":null` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":false` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":true` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":-2` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":-1` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":0` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":1` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":2` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":""` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":"hello world"` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"highlights":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawExperience{
				Highlights: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"location":null` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":false` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":true` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":-2` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":-1` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":0` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":1` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":2` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":""` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":"hello world"` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"location":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawExperience{
				Location: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"name":null` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":false` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":true` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":-2` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":-1` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":0` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":1` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":2` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":""` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"hello world"` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"name":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawExperience{
				Name: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"organization":null` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":false` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":true` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":-2` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":-1` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":0` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":1` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":2` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":""` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":"hello world"` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"organization":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawExperience{
				Organization: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"position":null` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":false` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":true` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":-2` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":-1` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":0` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":1` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":2` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":""` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":"hello world"` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"position":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawExperience{
				Position: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"startDate":null` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":false` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":true` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":-2` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":-1` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":0` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":1` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":2` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":""` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":"hello world"` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawExperience{
				StartDate: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"summary":null` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":false` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":true` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":-2` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":-1` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":0` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":1` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":2` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":""` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":"hello world"` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"summary":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawExperience{
				Summary: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"url":null` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":false` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":true` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-2` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-1` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":0` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":1` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":2` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":""` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"hello world"` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawExperience{
				URL: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawExperience{
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
					`"highlights":"THREE"` +
					`,` +
					`"location":"FOUR"` +
					`,` +
					`"name":"FIVE"` +
					`,` +
					`"organization":"SIX"` +
					`,` +
					`"position":"SEVEN"` +
					`,` +
					`"startDate":"EIGHT"` +
					`,` +
					`"summary":"NINE"` +
					`,` +
					`"url":"TEN"` +
				`}`,
			),
			Expected: rawExperience{
				Description:  []byte(`"ONE"`),
				EndDate:      []byte(`"TWO"`),
				Highlights:   []byte(`"THREE"`),
				Location:     []byte(`"FOUR"`),
				Name:         []byte(`"FIVE"`),
				Organization: []byte(`"SIX"`),
				Position:     []byte(`"SEVEN"`),
				StartDate:    []byte(`"EIGHT"`),
				Summary:      []byte(`"NINE"`),
				URL:          []byte(`"TEN"`),
			},
		},
	}

	for testNumber, test := range tests {
		var actual rawExperience

		err := json.Unmarshal(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual raw-experience is not what was expected.", testNumber)

			t.Logf("EXPECTED.DESCRIPTION:\n%s", expected.Description)
			t.Logf("EXPECTED.ENDDATE:\n%s", expected.EndDate)
			t.Logf("EXPECTED.HIGHLIGHTS:\n%s", expected.Highlights)
			t.Logf("EXPECTED.LOCATION:\n%s", expected.Location)
			t.Logf("EXPECTED.NAME:\n%s", expected.Name)
			t.Logf("EXPECTED.ORGANIZATION:\n%s", expected.Organization)
			t.Logf("EXPECTED.POSITION:\n%s", expected.Position)
			t.Logf("EXPECTED.STARTDATE:\n%s", expected.StartDate)
			t.Logf("EXPECTED.SUMMARY:\n%s", expected.Summary)
			t.Logf("EXPECTED.URL:\n%s", expected.URL)

			t.Logf("ACTUAL.DESCRIPTION:\n%s", actual.Description)
			t.Logf("ACTUAL.ENDDATE:\n%s", actual.EndDate)
			t.Logf("ACTUAL.HIGHLIGHTS:\n%s", actual.Highlights)
			t.Logf("ACTUAL.LOCATION:\n%s", actual.Location)
			t.Logf("ACTUAL.NAME:\n%s", actual.Name)
			t.Logf("ACTUAL.ORGANIZATION:\n%s", actual.Organization)
			t.Logf("ACTUAL.POSITION:\n%s", actual.Position)
			t.Logf("ACTUAL.STARTDATE:\n%s", actual.StartDate)
			t.Logf("ACTUAL.SUMMARY:\n%s", actual.Summary)
			t.Logf("ACTUAL.URL:\n%s", actual.URL)

			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
