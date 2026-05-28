package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

func TestProject_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON                string
		ExpectedName        nul.Nullable[string]
		ExpectedProjectType nul.Nullable[string]
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"@type":"Project"}`,
		},
		{
			JSON: `{"@type":"cv:Project"}`,
		},
		{
			JSON: `{"@type":"https://w3id.org/fep/6158#Project"}`,
		},
		{
			JSON: `{"id":"http://example.com/project/1","@type":"Project"}`,
		},
		{
			JSON: `{"id":"http://example.com/project/1","@type":"https://w3id.org/fep/6158#Project"}`,
		},
		{
			JSON:         `{"@type":"Project","name":"Microdon"}`,
			ExpectedName: nul.Something("Microdon"),
		},
		{
			JSON:         `{"@type":"cv:Project","name":"Microdon"}`,
			ExpectedName: nul.Something("Microdon"),
		},
		{
			JSON:         `{"@type":"https://w3id.org/fep/6158#Project","name":"Microdon"}`,
			ExpectedName: nul.Something("Microdon"),
		},
		{
			JSON:                `{"@type":"Project","type":"application"}`,
			ExpectedProjectType: nul.Something("application"),
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Project

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

		{
			expected := test.ExpectedProjectType
			if expected != actual.ProjectType {
				t.Errorf("For test #%d, the actual ProjectType is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual.ProjectType)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}

func TestProject_UnmarshalJSON_typeRejected(t *testing.T) {

	tests := []struct {
		JSON string
	}{
		{
			JSON: `{"@type":"Person"}`,
		},
		{
			JSON: `{"@type":"Resume"}`,
		},
		{
			JSON: `{"@type":""}`,
		},
		{
			JSON: `{"@type":"project"}`,
		},
		{
			JSON: `{"@type":"PROJECT"}`,
		},
		{
			JSON: `{"@type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"@type":"Person","name":"Microdon"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Project

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
