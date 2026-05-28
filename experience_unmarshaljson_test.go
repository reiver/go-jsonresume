package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

func TestExperience_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON         string
		ExpectedName nul.Nullable[string]
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Experience"}`,
		},
		{
			JSON: `{"type":"cv:Experience"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Experience"}`,
		},
		{
			JSON: `{"id":"http://example.com/exp/1","type":"Experience"}`,
		},
		{
			JSON: `{"id":"http://example.com/exp/1","type":"https://w3id.org/fep/6158#Experience"}`,
		},
		{
			JSON:         `{"type":"Experience","name":"SuperCo"}`,
			ExpectedName: nul.Something("SuperCo"),
		},
		{
			JSON:         `{"type":"cv:Experience","name":"SuperCo"}`,
			ExpectedName: nul.Something("SuperCo"),
		},
		{
			JSON:         `{"type":"https://w3id.org/fep/6158#Experience","name":"SuperCo"}`,
			ExpectedName: nul.Something("SuperCo"),
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Experience

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

func TestExperience_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"experience"}`,
		},
		{
			JSON: `{"type":"EXPERIENCE"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","name":"SuperCo"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Experience

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
