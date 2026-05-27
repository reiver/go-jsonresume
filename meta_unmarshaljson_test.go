package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestMeta_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON string
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Meta"}`,
		},
		{
			JSON: `{"type":"cv:Meta"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Meta"}`,
		},
		{
			JSON: `{"id":"http://example.com/meta/1","type":"Meta"}`,
		},
		{
			JSON: `{"id":"http://example.com/meta/1","type":"https://w3id.org/fep/6158#Meta"}`,
		},
		{
			JSON: `{"type":"Meta","version":"v1.0.0"}`,
		},
		{
			JSON: `{"type":"cv:Meta","version":"v1.0.0"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Meta","version":"v1.0.0"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Meta

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}

func TestMeta_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"meta"}`,
		},
		{
			JSON: `{"type":"META"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","version":"v1.0.0"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Meta

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
