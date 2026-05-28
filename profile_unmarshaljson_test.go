package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

func TestProfile_UnmarshalJSON_typeAccepted(t *testing.T) {

	tests := []struct {
		JSON            string
		ExpectedNetwork nul.Nullable[string]
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"type":"Profile"}`,
		},
		{
			JSON: `{"type":"cv:Profile"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Profile"}`,
		},
		{
			JSON: `{"id":"http://example.com/profile/1","type":"Profile"}`,
		},
		{
			JSON: `{"id":"http://example.com/profile/1","type":"https://w3id.org/fep/6158#Profile"}`,
		},
		{
			JSON:            `{"type":"Profile","network":"Mastodon"}`,
			ExpectedNetwork: nul.Something("Mastodon"),
		},
		{
			JSON:            `{"type":"cv:Profile","network":"Mastodon"}`,
			ExpectedNetwork: nul.Something("Mastodon"),
		},
		{
			JSON:            `{"type":"https://w3id.org/fep/6158#Profile","network":"Mastodon"}`,
			ExpectedNetwork: nul.Something("Mastodon"),
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Profile

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		{
			expected := test.ExpectedNetwork
			if expected != actual.Network {
				t.Errorf("For test #%d, the actual Network is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual.Network)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}

func TestProfile_UnmarshalJSON_typeRejected(t *testing.T) {

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
			JSON: `{"type":"profile"}`,
		},
		{
			JSON: `{"type":"PROFILE"}`,
		},
		{
			JSON: `{"type":"https://w3id.org/fep/6158#Person"}`,
		},
		{
			JSON: `{"type":"Person","network":"Mastodon"}`,
		},
	}

	for testNumber, test := range tests {
		var actual jsonresume.Profile

		err := jsonld.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
