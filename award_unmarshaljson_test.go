package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

func TestAward_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON          string
		ExpectedTitle nul.Nullable[string]
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Award"}`,
		},
		{
			JSON: `{"type":"cv:Award"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Award"}`,
		},
		{
			JSON: `{"id":"http://example.com/award/1","type":"Award"}`,
		},
		{
			JSON: `{"id":"http://example.com/award/1","type":"https://w3id.org/fep/6158#Award"}`,
		},
		{
			JSON:          `{"type":"Award","title":"Employee of the Year"}`,
			ExpectedTitle: nul.Something("Employee of the Year"),
		},
		{
			JSON:          `{"type":"cv:Award","title":"Employee of the Year"}`,
			ExpectedTitle: nul.Something("Employee of the Year"),
		},
		{
			JSON:          `{"type":"https://w3id.org/fep/6158#Award","title":"Employee of the Year"}`,
			ExpectedTitle: nul.Something("Employee of the Year"),
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Award

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		{
			expected := test.ExpectedTitle
			if expected != actual.Title {
				t.Errorf("For test #%d, the actual Title is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual.Title)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}

func TestAward_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"award"}`,
		},
		{
			JSON: `{"type":"AWARD"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","title":"Employee of the Year"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Award

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
