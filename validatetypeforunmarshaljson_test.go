package jsonresume

import (
	"testing"
)

func TestValidateTypeForUnmarshalJSON(t *testing.T) {

	tests := []struct {
		Bytes       []byte
		Entity      string
		ValidTypes  []string
		ExpectError bool
	}{
		// 0: nil bytes — no validation, no error
		{
			Bytes:      nil,
			Entity:     "thing",
			ValidTypes: []string{"Thing"},
		},

		// 1: empty bytes — no validation, no error
		{
			Bytes:      []byte{},
			Entity:     "thing",
			ValidTypes: []string{"Thing"},
		},

		// 2: valid type — exact match
		{
			Bytes:      []byte(`"Award"`),
			Entity:     "award",
			ValidTypes: []string{"Award", "cv:Award", "https://w3id.org/fep/6158#Award"},
		},

		// 3: valid type — compact form
		{
			Bytes:      []byte(`"cv:Award"`),
			Entity:     "award",
			ValidTypes: []string{"Award", "cv:Award", "https://w3id.org/fep/6158#Award"},
		},

		// 4: valid type — expanded form
		{
			Bytes:      []byte(`"https://w3id.org/fep/6158#Award"`),
			Entity:     "award",
			ValidTypes: []string{"Award", "cv:Award", "https://w3id.org/fep/6158#Award"},
		},

		// 5: invalid type — not in valid list
		{
			Bytes:       []byte(`"Bogus"`),
			Entity:      "award",
			ValidTypes:  []string{"Award", "cv:Award", "https://w3id.org/fep/6158#Award"},
			ExpectError: true,
		},

		// 6: malformed JSON — not a valid JSON string
		{
			Bytes:       []byte(`not-json`),
			Entity:      "award",
			ValidTypes:  []string{"Award"},
			ExpectError: true,
		},

		// 7: wrong JSON type — number instead of string
		{
			Bytes:       []byte(`42`),
			Entity:      "award",
			ValidTypes:  []string{"Award"},
			ExpectError: true,
		},

		// 8: wrong JSON type — object instead of string
		{
			Bytes:       []byte(`{"type":"Award"}`),
			Entity:      "award",
			ValidTypes:  []string{"Award"},
			ExpectError: true,
		},

		// 9: wrong JSON type — array instead of string
		{
			Bytes:       []byte(`["Award"]`),
			Entity:      "award",
			ValidTypes:  []string{"Award"},
			ExpectError: true,
		},

		// 10: valid type — single valid type in list
		{
			Bytes:      []byte(`"Resume"`),
			Entity:     "resume",
			ValidTypes: []string{"Resume"},
		},

		// 11: empty string type — not in valid list
		{
			Bytes:       []byte(`""`),
			Entity:      "thing",
			ValidTypes:  []string{"Thing"},
			ExpectError: true,
		},

		// 12: no valid types provided — any type is invalid
		{
			Bytes:       []byte(`"Award"`),
			Entity:      "award",
			ValidTypes:  []string{},
			ExpectError: true,
		},
	}

	for testNumber, test := range tests {

		err := validateTypeForUnmarshalJSON(test.Bytes, test.Entity, test.ValidTypes...)

		if test.ExpectError {
			if nil == err {
				t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
				t.Logf("BYTES:  %s", test.Bytes)
				t.Logf("ENTITY: %s", test.Entity)
				continue
			}
			continue
		}

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR:  %s", err)
			t.Logf("BYTES:  %s", test.Bytes)
			t.Logf("ENTITY: %s", test.Entity)
			continue
		}
	}
}
