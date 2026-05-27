package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestResume_UnmarshalJSON_malformedJSON(t *testing.T) {

	tests := []struct {
		Name string
		JSON string
	}{
		{
			Name: "truncated object",
			JSON: `{"type":"Resume","awards":[{"title":"Be`,
		},
		{
			Name: "truncated array",
			JSON: `{"type":"Resume","awards":[`,
		},
		{
			Name: "missing closing brace",
			JSON: `{"type":"Resume"`,
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
			var actual jsonresume.Resume

			err := jsonld.Unmarshal([]byte(test.JSON), &actual)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				t.Logf("JSON:\n%s", test.JSON)
			}
		})
	}
}

// TestResume_UnmarshalJSON_directMalformed calls UnmarshalJSON directly
// (bypassing the outer jsonld.Unmarshal) to verify that the inner
// jsonld.Unmarshal error path returns an error wrapping the cause,
// rather than panicking or silently proceeding with a zero-value rawResume.
func TestResume_UnmarshalJSON_directMalformed(t *testing.T) {

	tests := []struct {
		Name  string
		Bytes []byte
	}{
		{
			Name:  "truncated object",
			Bytes: []byte(`{"type":"Resume","awards":[`),
		},
		{
			Name:  "bare string",
			Bytes: []byte(`not json`),
		},
		{
			Name:  "empty",
			Bytes: []byte(``),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual jsonresume.Resume

			err := actual.UnmarshalJSON(test.Bytes)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				t.Logf("BYTES: %q", test.Bytes)
			}
		})
	}
}

func TestResume_UnmarshalJSON_typeFieldMalformed(t *testing.T) {

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
			JSON: `{"type":["Resume"]}`,
		},
		{
			Name: "type is object",
			JSON: `{"type":{"name":"Resume"}}`,
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

func TestAnyResume_UnmarshalJSON_malformedJSON(t *testing.T) {

	tests := []struct {
		Name string
		JSON string
	}{
		{
			Name: "truncated object",
			JSON: `{"type":"Resume","awards":[{"title":"Be`,
		},
		{
			Name: "truncated array",
			JSON: `{"type":"Resume","awards":[`,
		},
		{
			Name: "missing closing brace",
			JSON: `{"type":"Resume"`,
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
			var actual jsonresume.AnyResume

			err := jsonld.Unmarshal([]byte(test.JSON), &actual)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				t.Logf("JSON:\n%s", test.JSON)
			}
		})
	}
}
