package jsonresume

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestRawResume_unmarshalJSON(t *testing.T) {
	tests := []struct{
		JSON []byte
		Expected rawResume
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
					`"awards":null` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":false` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":true` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":-2` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":-1` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":0` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":1` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":2` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":""` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"awards":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Awards: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"basics":null` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":false` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":true` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":-2` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":-1` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":0` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":1` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":2` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":""` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"basics":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Basics: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"certificates":null` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":false` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":true` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":-2` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":-1` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":0` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":1` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":2` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":""` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"certificates":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Certificates: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"education":null` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":false` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":true` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":-2` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":-1` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":0` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":1` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":2` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":""` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"education":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Education: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"interests":null` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":false` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":true` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":-2` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":-1` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":0` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":1` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":2` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":""` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"interests":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Interests: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"languages":null` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":false` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":true` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":-2` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":-1` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":0` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":1` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":2` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":""` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"languages":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Languages: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"meta":null` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":false` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":true` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":-2` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":-1` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":0` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":1` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":2` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":""` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"meta":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Meta: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"projects":null` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":false` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":true` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":-2` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":-1` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":0` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":1` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":2` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":""` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"projects":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Projects: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"publications":null` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":false` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":true` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":-2` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":-1` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":0` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":1` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":2` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":""` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"publications":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Publications: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"references":null` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":false` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":true` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":-2` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":-1` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":0` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":1` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":2` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":""` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"references":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				References: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"skills":null` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":false` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":true` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":-2` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":-1` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":0` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":1` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":2` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":""` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"skills":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Skills: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"volunteer":null` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":false` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":true` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":-2` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":-1` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":0` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":1` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":2` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":""` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"volunteer":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Volunteer: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"work":null` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":false` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":true` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":-2` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":-1` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":0` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":1` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":2` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":""` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":"hello world"` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"work":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawResume{
				Work: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"awards":"ONE"` +
					`,` +
					`"basics":"TWO"` +
					`,` +
					`"education":"THREE"` +
					`,` +
					`"interests":"FOUR"` +
					`,` +
					`"languages":"FIVE"` +
					`,` +
					`"meta":"SIX"` +
					`,` +
					`"projects":"SEVEN"` +
					`,` +
					`"publications":"EIGHT"` +
					`,` +
					`"references":"NINE"` +
					`,` +
					`"skills":"TEN"` +
					`,` +
					`"volunteer":"ELEVEN"` +
					`,` +
					`"work":"TWELVE"` +
				`}`,
			),
			Expected: rawResume{
				Awards:       []byte(`"ONE"`),
				Basics:       []byte(`"TWO"`),
				Education:    []byte(`"THREE"`),
				Interests:    []byte(`"FOUR"`),
				Languages:    []byte(`"FIVE"`),
				Meta:         []byte(`"SIX"`),
				Projects:     []byte(`"SEVEN"`),
				Publications: []byte(`"EIGHT"`),
				References:   []byte(`"NINE"`),
				Skills:       []byte(`"TEN"`),
				Volunteer:    []byte(`"ELEVEN"`),
				Work:         []byte(`"TWELVE"`),
			},
		},
	}

	for testNumber, test := range tests {
		var actual rawResume

		err := json.Unmarshal(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual raw-resume is not what was expected.", testNumber)

			t.Logf("EXPECTED.AWARDS:\n%s", expected.Awards)
			t.Logf("EXPECTED.BASICS:\n%s", expected.Basics)
			t.Logf("EXPECTED.CERTIFICATES:\n%s", expected.Certificates)
			t.Logf("EXPECTED.EDUCATION:\n%s", expected.Education)
			t.Logf("EXPECTED.INTERESTS:\n%s", expected.Interests)
			t.Logf("EXPECTED.LANGUAGES:\n%s", expected.Languages)
			t.Logf("EXPECTED.META:\n%s", expected.Meta)
			t.Logf("EXPECTED.PROJECTS:\n%s", expected.Projects)
			t.Logf("EXPECTED.PUBLICATIONS:\n%s", expected.Publications)
			t.Logf("EXPECTED.REFERENCES:\n%s", expected.References)
			t.Logf("EXPECTED.SKILLS:\n%s", expected.Skills)
			t.Logf("EXPECTED.VOLUNTEER:\n%s", expected.Volunteer)
			t.Logf("EXPECTED.WORK:\n%s", expected.Work)

			t.Logf("ACTUAL.AWARDS:\n%s", actual.Awards)
			t.Logf("ACTUAL.BASICS:\n%s", actual.Basics)
			t.Logf("ACTUAL.CERTIFICATES:\n%s", actual.Certificates)
			t.Logf("ACTUAL.EDUCATION:\n%s", actual.Education)
			t.Logf("ACTUAL.INTERESTS:\n%s", actual.Interests)
			t.Logf("ACTUAL.LANGUAGES:\n%s", actual.Languages)
			t.Logf("ACTUAL.META:\n%s", actual.Meta)
			t.Logf("ACTUAL.PROJECTS:\n%s", actual.Projects)
			t.Logf("ACTUAL.PUBLICATIONS:\n%s", actual.Publications)
			t.Logf("ACTUAL.REFERENCES:\n%s", actual.References)
			t.Logf("ACTUAL.SKILLS:\n%s", actual.Skills)
			t.Logf("ACTUAL.VOLUNTEER:\n%s", actual.Volunteer)
			t.Logf("ACTUAL.WORK:\n%s", actual.Work)

			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
