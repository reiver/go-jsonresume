package jsonresume_test

import (
	"reflect"
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
