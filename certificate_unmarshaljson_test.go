package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestCertificate_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON string
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Certificate"}`,
		},
		{
			JSON: `{"type":"cv:Certificate"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Certificate"}`,
		},
		{
			JSON: `{"id":"http://example.com/cert/1","type":"Certificate"}`,
		},
		{
			JSON: `{"id":"http://example.com/cert/1","type":"https://w3id.org/fep/6158#Certificate"}`,
		},
		{
			JSON: `{"type":"Certificate","name":"CP+"}`,
		},
		{
			JSON: `{"type":"cv:Certificate","name":"CP+"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Certificate","name":"CP+"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Certificate

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}

func TestCertificate_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"certificate"}`,
		},
		{
			JSON: `{"type":"CERTIFICATE"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","name":"CP+"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Certificate

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
