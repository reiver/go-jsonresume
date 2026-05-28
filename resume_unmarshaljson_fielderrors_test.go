package jsonresume_test

import (
	"strings"
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestResume_UnmarshalJSON_fieldTypeErrors(t *testing.T) {

	tests := []struct {
		Name      string
		JSON      string
		ErrorHint string // substring expected in the error message
	}{
		// awards expects an array of objects or strings, not a plain number
		{
			Name:      "awards is number",
			JSON:      `{"type":"Resume","awards":42}`,
			ErrorHint: "awards",
		},
		// basics expects an object or string, not an array
		{
			Name:      "basics is array",
			JSON:      `{"type":"Resume","basics":[1,2,3]}`,
			ErrorHint: "basics",
		},
		// basics expects an object or string, not a number
		{
			Name:      "basics is number",
			JSON:      `{"type":"Resume","basics":99}`,
			ErrorHint: "basics",
		},
		// meta expects an object or string, not a number
		{
			Name:      "meta is number",
			JSON:      `{"type":"Resume","meta":42}`,
			ErrorHint: "meta",
		},
		// meta expects an object or string, not an array
		{
			Name:      "meta is array",
			JSON:      `{"type":"Resume","meta":[1,2]}`,
			ErrorHint: "meta",
		},
		// skills expects an array of objects or strings, not a plain number
		{
			Name:      "skills is number",
			JSON:      `{"type":"Resume","skills":42}`,
			ErrorHint: "skills",
		},
		// work expects an array of objects or strings, not a plain string
		{
			Name:      "work is boolean",
			JSON:      `{"type":"Resume","work":true}`,
			ErrorHint: "work",
		},
		// education expects an array, not a boolean
		{
			Name:      "education is boolean",
			JSON:      `{"type":"Resume","education":false}`,
			ErrorHint: "education",
		},
		// id expects a string, not a number
		{
			Name:      "id is number",
			JSON:      `{"type":"Resume","id":123}`,
			ErrorHint: "id",
		},
		// $schema expects a string, not a number
		{
			Name:      "$schema is number",
			JSON:      `{"type":"Resume","$schema":123}`,
			ErrorHint: "$schema",
		},
		// certificates expects an array, not a number
		{
			Name:      "certificates is number",
			JSON:      `{"type":"Resume","certificates":42}`,
			ErrorHint: "certificates",
		},
		// interests expects an array, not a number
		{
			Name:      "interests is number",
			JSON:      `{"type":"Resume","interests":42}`,
			ErrorHint: "interests",
		},
		// languages expects an array, not a number
		{
			Name:      "languages is number",
			JSON:      `{"type":"Resume","languages":42}`,
			ErrorHint: "languages",
		},
		// projects expects an array, not a number
		{
			Name:      "projects is number",
			JSON:      `{"type":"Resume","projects":42}`,
			ErrorHint: "projects",
		},
		// publications expects an array, not a number
		{
			Name:      "publications is number",
			JSON:      `{"type":"Resume","publications":42}`,
			ErrorHint: "publications",
		},
		// references expects an array, not a number
		{
			Name:      "references is number",
			JSON:      `{"type":"Resume","references":42}`,
			ErrorHint: "references",
		},
		// volunteer expects an array, not a number
		{
			Name:      "volunteer is number",
			JSON:      `{"type":"Resume","volunteer":42}`,
			ErrorHint: "volunteer",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual jsonresume.Resume

			err := jsonld.Unmarshal([]byte(test.JSON), &actual)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				t.Logf("JSON:\n%s", test.JSON)
				return
			}
			if !strings.Contains(err.Error(), test.ErrorHint) {
				t.Errorf("Error should mention %q but got: %s", test.ErrorHint, err)
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
