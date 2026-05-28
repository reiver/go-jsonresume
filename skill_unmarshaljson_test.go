package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

func TestSkill_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON         string
		ExpectedName nul.Nullable[string]
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Skill"}`,
		},
		{
			JSON: `{"type":"cv:Skill"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Skill"}`,
		},
		{
			JSON: `{"id":"http://example.com/skill/1","type":"Skill"}`,
		},
		{
			JSON: `{"id":"http://example.com/skill/1","type":"https://w3id.org/fep/6158#Skill"}`,
		},
		{
			JSON:         `{"type":"Skill","name":"Go Programming"}`,
			ExpectedName: nul.Something("Go Programming"),
		},
		{
			JSON:         `{"type":"cv:Skill","name":"Go Programming"}`,
			ExpectedName: nul.Something("Go Programming"),
		},
		{
			JSON:         `{"type":"https://w3id.org/fep/6158#Skill","name":"Go Programming"}`,
			ExpectedName: nul.Something("Go Programming"),
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Skill

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

func TestSkill_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"skill"}`,
		},
		{
			JSON: `{"type":"SKILL"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","name":"Go Programming"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Skill

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
