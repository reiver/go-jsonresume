package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

func TestReference_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON         string
		ExpectedName nul.Nullable[string]
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Reference"}`,
		},
		{
			JSON: `{"type":"cv:Reference"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Reference"}`,
		},
		{
			JSON: `{"id":"http://example.com/reference/1","type":"Reference"}`,
		},
		{
			JSON: `{"id":"http://example.com/reference/1","type":"https://w3id.org/fep/6158#Reference"}`,
		},
		{
			JSON:         `{"type":"Reference","name":"Jane Smith"}`,
			ExpectedName: nul.Something("Jane Smith"),
		},
		{
			JSON:         `{"type":"cv:Reference","name":"Jane Smith"}`,
			ExpectedName: nul.Something("Jane Smith"),
		},
		{
			JSON:         `{"type":"https://w3id.org/fep/6158#Reference","name":"Jane Smith"}`,
			ExpectedName: nul.Something("Jane Smith"),
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Reference

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		{
			expected := test.ExpectedName
			if expected != actual.Name {
				t.Errorf("For test #%d, the actual Name is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual.Name)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}

func TestReference_UnmarshalJSON_typeRejected(t *testing.T) {

	tests := []struct {
		JSON string
	}{
		{
			JSON: `{"type":"Person"}`,
		},
		{
			JSON: `{"type":"Resume"}`,
		},
		{
			JSON: `{"type":""}`,
		},
		{
			JSON: `{"type":"reference"}`,
		},
		{
			JSON: `{"type":"REFERENCE"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","name":"Jane Smith"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Reference

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
