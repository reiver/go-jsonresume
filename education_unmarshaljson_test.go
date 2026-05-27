package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestEducation_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON string
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Education"}`,
		},
		{
			JSON: `{"type":"cv:Education"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Education"}`,
		},
		{
			JSON: `{"id":"http://example.com/edu/1","type":"Education"}`,
		},
		{
			JSON: `{"id":"http://example.com/edu/1","type":"https://w3id.org/fep/6158#Education"}`,
		},
		{
			JSON: `{"type":"Education","institution":"SFU"}`,
		},
		{
			JSON: `{"type":"cv:Education","institution":"SFU"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Education","institution":"SFU"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Education

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}

func TestEducation_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"education"}`,
		},
		{
			JSON: `{"type":"EDUCATION"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","institution":"SFU"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Education

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
