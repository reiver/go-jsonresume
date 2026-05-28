package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

func TestLanguage_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON             string
		ExpectedLanguage nul.Nullable[string]
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Language"}`,
		},
		{
			JSON: `{"type":"cv:Language"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Language"}`,
		},
		{
			JSON: `{"id":"http://example.com/language/1","type":"Language"}`,
		},
		{
			JSON: `{"id":"http://example.com/language/1","type":"https://w3id.org/fep/6158#Language"}`,
		},
		{
			JSON:             `{"type":"Language","language":"English"}`,
			ExpectedLanguage: nul.Something("English"),
		},
		{
			JSON:             `{"type":"cv:Language","language":"English"}`,
			ExpectedLanguage: nul.Something("English"),
		},
		{
			JSON:             `{"type":"https://w3id.org/fep/6158#Language","language":"English"}`,
			ExpectedLanguage: nul.Something("English"),
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Language

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		{
			expected := test.ExpectedLanguage
			if expected != actual.Language {
				t.Errorf("For test #%d, the actual Language is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual.Language)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}

func TestLanguage_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"language"}`,
		},
		{
			JSON: `{"type":"LANGUAGE"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","language":"English"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Language

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
