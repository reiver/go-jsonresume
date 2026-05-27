package jsonresume_test

import (
	"strings"
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestResume_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON string
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Resume"}`,
		},
		{
			JSON: `{"type":"cv:Resume"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Resume"}`,
		},
		{
			JSON: `{"id":"http://example.com/resume/1","type":"Resume"}`,
		},
		{
			JSON: `{"id":"http://example.com/resume/1","type":"https://w3id.org/fep/6158#Resume"}`,
		},
		{
			JSON: `{"type":"Resume","awards":[{"title":"Best Employee"}]}`,
		},
		{
			JSON: `{"type":"cv:Resume","awards":[{"title":"Best Employee"}]}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Resume","awards":[{"title":"Best Employee"}]}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Resume

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		{
			marshaled := actual.String()

			if !strings.Contains(marshaled, `"type": "Resume"`) {
				t.Errorf("For test #%d, marshaled output does not contain expected type.", testNumber)
				t.Logf("MARSHALED:\n%s", marshaled)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}

func TestResume_UnmarshalJSON_typeRejected(t *testing.T) {

	tests := []struct {
		JSON string
	}{
		{
			JSON: `{"type":"Person"}`,
		},
		{
			JSON: `{"type":"Award"}`,
		},
		{
			JSON: `{"type":""}`,
		},
		{
			JSON: `{"type":"resume"}`,
		},
		{
			JSON: `{"type":"RESUME"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","awards":[{"title":"Best Employee"}]}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Resume

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
