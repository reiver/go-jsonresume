package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

func TestInterest_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON         string
		ExpectedName nul.Nullable[string]
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Interest"}`,
		},
		{
			JSON: `{"type":"cv:Interest"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Interest"}`,
		},
		{
			JSON: `{"id":"http://example.com/interest/1","type":"Interest"}`,
		},
		{
			JSON: `{"id":"http://example.com/interest/1","type":"https://w3id.org/fep/6158#Interest"}`,
		},
		{
			JSON:         `{"type":"Interest","name":"Machine Learning"}`,
			ExpectedName: nul.Something("Machine Learning"),
		},
		{
			JSON:         `{"type":"cv:Interest","name":"Machine Learning"}`,
			ExpectedName: nul.Something("Machine Learning"),
		},
		{
			JSON:         `{"type":"https://w3id.org/fep/6158#Interest","name":"Machine Learning"}`,
			ExpectedName: nul.Something("Machine Learning"),
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Interest

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

func TestInterest_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"interest"}`,
		},
		{
			JSON: `{"type":"INTEREST"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","name":"Machine Learning"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Interest

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
