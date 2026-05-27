package jsonresume_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

// TestAnyType_UnmarshalJSON_typeAccepted verifies that all Any types accept
// any type value without rejection, unlike the main types which validate.
// It also verifies that the type value is correctly stored.
func TestAnyType_UnmarshalJSON_typeAccepted(t *testing.T) {

	typeInputs := []struct {
		JSON          string
		ExpectedTypes []string
	}{
		// 0: correct type
		{
			JSON:          `{"type":"Award"}`,
			ExpectedTypes: []string{"Award"},
		},

		// 1: wrong type
		{
			JSON:          `{"type":"Person"}`,
			ExpectedTypes: []string{"Person"},
		},

		// 2: empty string type
		{
			JSON:          `{"type":""}`,
			ExpectedTypes: []string{""},
		},

		// 3: no type at all
		{
			JSON:          `{}`,
			ExpectedTypes: nil,
		},

		// 4: lowercase type
		{
			JSON:          `{"type":"award"}`,
			ExpectedTypes: []string{"award"},
		},

		// 5: array of types
		{
			JSON:          `{"type":["Award","Thing"]}`,
			ExpectedTypes: []string{"Award", "Thing"},
		},

		// 6: array with single type
		{
			JSON:          `{"type":["Award"]}`,
			ExpectedTypes: []string{"Award"},
		},

		// 7: array with unrelated types
		{
			JSON:          `{"type":["Person","Agent"]}`,
			ExpectedTypes: []string{"Person", "Agent"},
		},
	}

	projectInputs := []struct {
		JSON          string
		ExpectedTypes []string
	}{
		// 0: correct type
		{
			JSON:          `{"@type":"Project"}`,
			ExpectedTypes: []string{"Project"},
		},

		// 1: wrong type
		{
			JSON:          `{"@type":"Person"}`,
			ExpectedTypes: []string{"Person"},
		},

		// 2: empty string type
		{
			JSON:          `{"@type":""}`,
			ExpectedTypes: []string{""},
		},

		// 3: no type at all
		{
			JSON:          `{}`,
			ExpectedTypes: nil,
		},

		// 4: lowercase type
		{
			JSON:          `{"@type":"project"}`,
			ExpectedTypes: []string{"project"},
		},

		// 5: array of types
		{
			JSON:          `{"@type":["Project","Thing"]}`,
			ExpectedTypes: []string{"Project", "Thing"},
		},

		// 6: array with single type
		{
			JSON:          `{"@type":["Project"]}`,
			ExpectedTypes: []string{"Project"},
		},

		// 7: array with unrelated types
		{
			JSON:          `{"@type":["Person","Agent"]}`,
			ExpectedTypes: []string{"Person", "Agent"},
		},
	}

	checkTypes := func(t *testing.T, name string, inputNumber int, json string, actual jsonld.Types, expected []string) {
		t.Helper()

		actualStrings := actual.Strings()

		if nil == expected && nil == actualStrings {
			return
		}

		if !reflect.DeepEqual(expected, actualStrings) {
			t.Errorf("%s: input #%d, the stored type is not what was expected.", name, inputNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actualStrings)
			t.Logf("JSON:     %s", json)
		}
	}

	for inputNumber, input := range typeInputs {
		{
			var v jsonresume.AnyAward
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyAward: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyAward", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyBasics
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyBasics: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyBasics", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyCertificate
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyCertificate: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyCertificate", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyEducation
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyEducation: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyEducation", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyExperience
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyExperience: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyExperience", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyInterest
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyInterest: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyInterest", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyLanguage
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyLanguage: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyLanguage", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyLocation
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyLocation: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyLocation", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyMeta
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyMeta: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyMeta", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyProfile
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyProfile: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyProfile", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyPublication
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyPublication: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyPublication", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyReference
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyReference: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyReference", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnyResume
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnyResume: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnyResume", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
		{
			var v jsonresume.AnySkill
			if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
				t.Errorf("AnySkill: input #%d, did not expect an error but got one.", inputNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("JSON:  %s", input.JSON)
				continue
			}
			checkTypes(t, "AnySkill", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
		}
	}

	// AnyProject uses "@type" instead of "type".
	for inputNumber, input := range projectInputs {
		var v jsonresume.AnyProject
		if err := jsonld.Unmarshal([]byte(input.JSON), &v); nil != err {
			t.Errorf("AnyProject: input #%d, did not expect an error but got one.", inputNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:  %s", input.JSON)
			continue
		}
		checkTypes(t, "AnyProject", inputNumber, input.JSON, v.Type, input.ExpectedTypes)
	}
}

// TestAnyType_UnmarshalJSON_nilReceiver verifies that calling UnmarshalJSON
// on a nil pointer returns ErrReceiverNil rather than panicking.
func TestAnyType_UnmarshalJSON_nilReceiver(t *testing.T) {

	tests := []struct {
		Name string
		Fn   func([]byte) error
	}{
		{Name: "AnyAward",       Fn: (*jsonresume.AnyAward)(nil).UnmarshalJSON},
		{Name: "AnyBasics",      Fn: (*jsonresume.AnyBasics)(nil).UnmarshalJSON},
		{Name: "AnyCertificate", Fn: (*jsonresume.AnyCertificate)(nil).UnmarshalJSON},
		{Name: "AnyEducation",   Fn: (*jsonresume.AnyEducation)(nil).UnmarshalJSON},
		{Name: "AnyExperience",  Fn: (*jsonresume.AnyExperience)(nil).UnmarshalJSON},
		{Name: "AnyInterest",    Fn: (*jsonresume.AnyInterest)(nil).UnmarshalJSON},
		{Name: "AnyLanguage",    Fn: (*jsonresume.AnyLanguage)(nil).UnmarshalJSON},
		{Name: "AnyLocation",    Fn: (*jsonresume.AnyLocation)(nil).UnmarshalJSON},
		{Name: "AnyMeta",        Fn: (*jsonresume.AnyMeta)(nil).UnmarshalJSON},
		{Name: "AnyProfile",     Fn: (*jsonresume.AnyProfile)(nil).UnmarshalJSON},
		{Name: "AnyProject",     Fn: (*jsonresume.AnyProject)(nil).UnmarshalJSON},
		{Name: "AnyPublication", Fn: (*jsonresume.AnyPublication)(nil).UnmarshalJSON},
		{Name: "AnyReference",   Fn: (*jsonresume.AnyReference)(nil).UnmarshalJSON},
		{Name: "AnyResume",      Fn: (*jsonresume.AnyResume)(nil).UnmarshalJSON},
		{Name: "AnySkill",       Fn: (*jsonresume.AnySkill)(nil).UnmarshalJSON},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			err := test.Fn([]byte(`{}`))
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				return
			}
			if !errors.Is(err, jsonresume.ErrReceiverNil) {
				t.Errorf("Expected ErrReceiverNil but got: %s", err)
			}
		})
	}
}

// TestAnyType_UnmarshalJSON_directMalformed calls UnmarshalJSON directly
// (bypassing the outer jsonld.Unmarshal) to verify that the inner
// jsonld.Unmarshal error path returns an error rather than panicking
// or silently proceeding with a zero-value raw struct.
func TestAnyType_UnmarshalJSON_directMalformed(t *testing.T) {

	malformed := []byte(`{"type":"Award","title":`)

	tests := []struct {
		Name string
		Fn   func([]byte) error
	}{
		{Name: "AnyAward",       Fn: new(jsonresume.AnyAward).UnmarshalJSON},
		{Name: "AnyBasics",      Fn: new(jsonresume.AnyBasics).UnmarshalJSON},
		{Name: "AnyCertificate", Fn: new(jsonresume.AnyCertificate).UnmarshalJSON},
		{Name: "AnyEducation",   Fn: new(jsonresume.AnyEducation).UnmarshalJSON},
		{Name: "AnyExperience",  Fn: new(jsonresume.AnyExperience).UnmarshalJSON},
		{Name: "AnyInterest",    Fn: new(jsonresume.AnyInterest).UnmarshalJSON},
		{Name: "AnyLanguage",    Fn: new(jsonresume.AnyLanguage).UnmarshalJSON},
		{Name: "AnyLocation",    Fn: new(jsonresume.AnyLocation).UnmarshalJSON},
		{Name: "AnyMeta",        Fn: new(jsonresume.AnyMeta).UnmarshalJSON},
		{Name: "AnyProfile",     Fn: new(jsonresume.AnyProfile).UnmarshalJSON},
		{Name: "AnyProject",     Fn: new(jsonresume.AnyProject).UnmarshalJSON},
		{Name: "AnyPublication", Fn: new(jsonresume.AnyPublication).UnmarshalJSON},
		{Name: "AnyReference",   Fn: new(jsonresume.AnyReference).UnmarshalJSON},
		{Name: "AnyResume",      Fn: new(jsonresume.AnyResume).UnmarshalJSON},
		{Name: "AnySkill",       Fn: new(jsonresume.AnySkill).UnmarshalJSON},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			err := test.Fn(malformed)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
			}
		})
	}
}

// TestAnyType_UnmarshalJSON_malformedType verifies that a non-string type
// field value (e.g. a number) causes an error in the type unmarshal step,
// rather than being silently accepted or causing a panic.
func TestAnyType_UnmarshalJSON_malformedType(t *testing.T) {

	// "type" is a number — valid JSON object but invalid for jsonld.Types.
	typeJSON := []byte(`{"type":123}`)
	// AnyProject uses "@type" instead of "type".
	projectJSON := []byte(`{"@type":123}`)

	tests := []struct {
		Name  string
		Fn    func([]byte) error
		Input []byte
	}{
		{Name: "AnyAward",       Fn: new(jsonresume.AnyAward).UnmarshalJSON,       Input: typeJSON},
		{Name: "AnyBasics",      Fn: new(jsonresume.AnyBasics).UnmarshalJSON,      Input: typeJSON},
		{Name: "AnyCertificate", Fn: new(jsonresume.AnyCertificate).UnmarshalJSON, Input: typeJSON},
		{Name: "AnyEducation",   Fn: new(jsonresume.AnyEducation).UnmarshalJSON,   Input: typeJSON},
		{Name: "AnyExperience",  Fn: new(jsonresume.AnyExperience).UnmarshalJSON,  Input: typeJSON},
		{Name: "AnyInterest",    Fn: new(jsonresume.AnyInterest).UnmarshalJSON,    Input: typeJSON},
		{Name: "AnyLanguage",    Fn: new(jsonresume.AnyLanguage).UnmarshalJSON,    Input: typeJSON},
		{Name: "AnyLocation",    Fn: new(jsonresume.AnyLocation).UnmarshalJSON,    Input: typeJSON},
		{Name: "AnyMeta",        Fn: new(jsonresume.AnyMeta).UnmarshalJSON,        Input: typeJSON},
		{Name: "AnyProfile",     Fn: new(jsonresume.AnyProfile).UnmarshalJSON,     Input: typeJSON},
		{Name: "AnyProject",     Fn: new(jsonresume.AnyProject).UnmarshalJSON,     Input: projectJSON},
		{Name: "AnyPublication", Fn: new(jsonresume.AnyPublication).UnmarshalJSON, Input: typeJSON},
		{Name: "AnyReference",   Fn: new(jsonresume.AnyReference).UnmarshalJSON,   Input: typeJSON},
		{Name: "AnyResume",      Fn: new(jsonresume.AnyResume).UnmarshalJSON,      Input: typeJSON},
		{Name: "AnySkill",       Fn: new(jsonresume.AnySkill).UnmarshalJSON,       Input: typeJSON},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			err := test.Fn(test.Input)
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				return
			}
			if !strings.Contains(err.Error(), "type") {
				t.Errorf("Expected error to mention 'type' but got: %s", err)
			}
		})
	}
}
