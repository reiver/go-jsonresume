package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestResume_UnmarshalJSON_fieldTypeErrors(t *testing.T) {

	tests := []struct {
		Name string
		JSON string
	}{
		// awards expects an array of objects or strings, not a plain number
		{
			Name: "awards is number",
			JSON: `{"type":"Resume","awards":42}`,
		},
		// basics expects an object or string, not an array
		{
			Name: "basics is array",
			JSON: `{"type":"Resume","basics":[1,2,3]}`,
		},
		// basics expects an object or string, not a number
		{
			Name: "basics is number",
			JSON: `{"type":"Resume","basics":99}`,
		},
		// meta expects an object or string, not a number
		{
			Name: "meta is number",
			JSON: `{"type":"Resume","meta":42}`,
		},
		// meta expects an object or string, not an array
		{
			Name: "meta is array",
			JSON: `{"type":"Resume","meta":[1,2]}`,
		},
		// skills expects an array of objects or strings, not a plain number
		{
			Name: "skills is number",
			JSON: `{"type":"Resume","skills":42}`,
		},
		// work expects an array of objects or strings, not a plain string
		{
			Name: "work is boolean",
			JSON: `{"type":"Resume","work":true}`,
		},
		// education expects an array, not a boolean
		{
			Name: "education is boolean",
			JSON: `{"type":"Resume","education":false}`,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual jsonresume.Resume

			err := jsonld.Unmarshal([]byte(test.JSON), &actual)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				t.Logf("JSON:\n%s", test.JSON)
			}
		})
	}
}

func TestAward_UnmarshalJSON_malformedJSON(t *testing.T) {

	tests := []struct {
		Name string
		JSON string
	}{
		{
			Name: "truncated object",
			JSON: `{"type":"Award","title":"Best`,
		},
		{
			Name: "missing closing brace",
			JSON: `{"type":"Award"`,
		},
		{
			Name: "bare string",
			JSON: `hello`,
		},
		{
			Name: "empty string",
			JSON: ``,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual jsonresume.Award

			err := jsonld.Unmarshal([]byte(test.JSON), &actual)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				t.Logf("JSON:\n%s", test.JSON)
			}
		})
	}
}

func TestAward_UnmarshalJSON_typeFieldMalformed(t *testing.T) {

	tests := []struct {
		Name string
		JSON string
	}{
		{
			Name: "type is number",
			JSON: `{"type":123}`,
		},
		{
			Name: "type is boolean",
			JSON: `{"type":true}`,
		},
		{
			Name: "type is array",
			JSON: `{"type":["Award"]}`,
		},
		{
			Name: "type is object",
			JSON: `{"type":{"name":"Award"}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual jsonresume.Award

			err := jsonld.Unmarshal([]byte(test.JSON), &actual)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				t.Logf("JSON:\n%s", test.JSON)
			}
		})
	}
}

func TestMeta_UnmarshalJSON_malformedJSON(t *testing.T) {

	tests := []struct {
		Name string
		JSON string
	}{
		{
			Name: "truncated object",
			JSON: `{"type":"Meta","version":"v1`,
		},
		{
			Name: "missing closing brace",
			JSON: `{"type":"Meta"`,
		},
		{
			Name: "bare string",
			JSON: `hello`,
		},
		{
			Name: "empty string",
			JSON: ``,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual jsonresume.Meta

			err := jsonld.Unmarshal([]byte(test.JSON), &actual)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				t.Logf("JSON:\n%s", test.JSON)
			}
		})
	}
}

func TestMeta_UnmarshalJSON_typeFieldMalformed(t *testing.T) {

	tests := []struct {
		Name string
		JSON string
	}{
		{
			Name: "type is number",
			JSON: `{"type":123}`,
		},
		{
			Name: "type is boolean",
			JSON: `{"type":true}`,
		},
		{
			Name: "type is array",
			JSON: `{"type":["Meta"]}`,
		},
		{
			Name: "type is object",
			JSON: `{"type":{"name":"Meta"}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual jsonresume.Meta

			err := jsonld.Unmarshal([]byte(test.JSON), &actual)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				t.Logf("JSON:\n%s", test.JSON)
			}
		})
	}
}
