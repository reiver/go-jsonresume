package jsonresume

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestRawEducation_unmarshalJSON(t *testing.T) {
	tests := []struct{
		JSON []byte
		Expected rawEducation
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
					`"area":null` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":false` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":true` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":-2` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":-1` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":0` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":1` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":2` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":""` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":"hello world"` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"area":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawEducation{
				Area: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"courses":null` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":false` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":true` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":-2` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":-1` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":0` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":1` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":2` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":""` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":"hello world"` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"courses":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawEducation{
				Courses: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"endDate":null` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":false` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":true` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":-2` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":-1` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":0` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":1` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":2` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":""` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":"hello world"` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"endDate":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawEducation{
				EndDate: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"institution":null` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":false` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":true` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":-2` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":-1` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":0` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":1` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":2` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":""` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":"hello world"` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"institution":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawEducation{
				Institution: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"score":null` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":false` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":true` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":-2` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":-1` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":0` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":1` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":2` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":""` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":"hello world"` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"score":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawEducation{
				Score: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"startDate":null` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":false` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":true` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":-2` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":-1` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":0` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":1` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":2` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":""` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":"hello world"` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"startDate":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawEducation{
				StartDate: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"studyType":null` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":false` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":true` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":-2` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":-1` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":0` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":1` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":2` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":""` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":"hello world"` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"studyType":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawEducation{
				StudyType: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"url":null` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`null`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":false` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`false`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":true` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`true`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-2` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`-2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":-1` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`-1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":0` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`0`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":1` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`1`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":2` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`2`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":""` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`""`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"hello world"` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`"hello world"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"https://example.com/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`"https://example.com/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":"HTTPS://EXAMPLE.COM/id/123"` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`"HTTPS://EXAMPLE.COM/id/123"`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`[null,false,true,-2,-1,0,1,2,"","hello world","https://example.com/id/12","HTTPS://EXAMPLE.COM/id/123"]`),
			},
		},
		{
			JSON: []byte(
				`{` +
					`"url":{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}` +
				`}`,
			),
			Expected: rawEducation{
				URL: []byte(`{"once":1,"twice":"two","thrice":"3","fource":"FOUR"}`),
			},
		},



		{
			JSON: []byte(
				`{` +
					`"area":"ONE"` +
					`,` +
					`"courses":"TWO"` +
					`,` +
					`"endDate":"THREE"` +
					`,` +
					`"institution":"FOUR"` +
					`,` +
					`"score":"FIVE"` +
					`,` +
					`"startDate":"SIX"` +
					`,` +
					`"studyType":"SEVEN"` +
					`,` +
					`"url":"EIGHT"` +
				`}`,
			),
			Expected: rawEducation{
				Area:        []byte(`"ONE"`),
				Courses:     []byte(`"TWO"`),
				EndDate:     []byte(`"THREE"`),
				Institution: []byte(`"FOUR"`),
				Score:       []byte(`"FIVE"`),
				StartDate:   []byte(`"SIX"`),
				StudyType:   []byte(`"SEVEN"`),
				URL:         []byte(`"EIGHT"`),
			},
		},
	}

	for testNumber, test := range tests {
		var actual rawEducation

		err := json.Unmarshal(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual raw-education is not what was expected.", testNumber)

			t.Logf("EXPECTED.AREA:\n%s", expected.Area)
			t.Logf("EXPECTED.COURSES:\n%s", expected.Courses)
			t.Logf("EXPECTED.ENDDATE:\n%s", expected.EndDate)
			t.Logf("EXPECTED.INSTITUTION:\n%s", expected.Institution)
			t.Logf("EXPECTED.SCORE:\n%s", expected.Score)
			t.Logf("EXPECTED.STARTDATE:\n%s", expected.StartDate)
			t.Logf("EXPECTED.STUDYTYPE:\n%s", expected.StudyType)
			t.Logf("EXPECTED.URL:\n%s", expected.URL)

			t.Logf("ACTUAL.AREA:\n%s", actual.Area)
			t.Logf("ACTUAL.COURSES:\n%s", actual.Courses)
			t.Logf("ACTUAL.ENDDATE:\n%s", actual.EndDate)
			t.Logf("ACTUAL.INSTITUTION:\n%s", actual.Institution)
			t.Logf("ACTUAL.SCORE:\n%s", actual.Score)
			t.Logf("ACTUAL.STARTDATE:\n%s", actual.StartDate)
			t.Logf("ACTUAL.STUDYTYPE:\n%s", actual.StudyType)
			t.Logf("ACTUAL.URL:\n%s", actual.URL)

			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
