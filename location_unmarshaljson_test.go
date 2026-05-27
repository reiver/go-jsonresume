package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestLocation_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON string
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Location"}`,
		},
		{
			JSON: `{"type":"cv:Location"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Location"}`,
		},
		{
			JSON: `{"id":"http://example.com/location/1","type":"Location"}`,
		},
		{
			JSON: `{"id":"http://example.com/location/1","type":"https://w3id.org/fep/6158#Location"}`,
		},
		{
			JSON: `{"type":"Location","city":"Vancouver"}`,
		},
		{
			JSON: `{"type":"cv:Location","city":"Vancouver"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Location","city":"Vancouver"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Location

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}

func TestLocation_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"location"}`,
		},
		{
			JSON: `{"type":"LOCATION"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","city":"Vancouver"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Location

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
