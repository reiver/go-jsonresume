package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

func TestBasics_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON         string
		ExpectedName nul.Nullable[string]
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Basics"}`,
		},
		{
			JSON: `{"type":"cv:Basics"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Basics"}`,
		},
		{
			JSON: `{"id":"http://example.com/basics/1","type":"Basics"}`,
		},
		{
			JSON: `{"id":"http://example.com/basics/1","type":"https://w3id.org/fep/6158#Basics"}`,
		},
		{
			JSON:         `{"type":"Basics","name":"John Doe"}`,
			ExpectedName: nul.Something("John Doe"),
		},
		{
			JSON:         `{"type":"cv:Basics","name":"John Doe"}`,
			ExpectedName: nul.Something("John Doe"),
		},
		{
			JSON:         `{"type":"https://w3id.org/fep/6158#Basics","name":"John Doe"}`,
			ExpectedName: nul.Something("John Doe"),
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Basics

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

func TestBasics_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"basics"}`,
		},
		{
			JSON: `{"type":"BASICS"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","name":"John Doe"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Basics

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
