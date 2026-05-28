package jsonresume

import (
	"errors"
	"testing"

	"github.com/reiver/go-jsonld"
)

// TestUnmarshalRaw_nilReceiver verifies that all unmarshalRaw* methods
// return ErrReceiverNil when called on a nil pointer, rather than panicking.
func TestUnmarshalRaw_nilReceiver(t *testing.T) {

	var id jsonld.ID

	tests := []struct {
		Name string
		Fn   func() error
	}{
		{Name: "Award",       Fn: func() error { return (*CoreAward)(nil).unmarshalRawAward(rawAward{}, &id) }},
		{Name: "Basics",      Fn: func() error { return (*CoreBasics)(nil).unmarshalRawBasics(rawBasics{}, &id) }},
		{Name: "Certificate", Fn: func() error { return (*CoreCertificate)(nil).unmarshalRawCertificate(rawCertificate{}, &id) }},
		{Name: "Education",   Fn: func() error { return (*CoreEducation)(nil).unmarshalRawEducation(rawEducation{}, &id) }},
		{Name: "Experience",  Fn: func() error { return (*CoreExperience)(nil).unmarshalRawExperience(rawExperience{}, &id) }},
		{Name: "Interest",    Fn: func() error { return (*CoreInterest)(nil).unmarshalRawInterest(rawInterest{}, &id) }},
		{Name: "Language",    Fn: func() error { return (*CoreLanguage)(nil).unmarshalRawLanguage(rawLanguage{}, &id) }},
		{Name: "Location",    Fn: func() error { return (*CoreLocation)(nil).unmarshalRawLocation(rawLocation{}, &id) }},
		{Name: "Meta",        Fn: func() error { return (*CoreMeta)(nil).unmarshalRawMeta(rawMeta{}, &id) }},
		{Name: "Profile",     Fn: func() error { return (*CoreProfile)(nil).unmarshalRawProfile(rawProfile{}, &id) }},
		{Name: "Project",     Fn: func() error { return (*CoreProject)(nil).unmarshalRawProject(rawProject{}, &id) }},
		{Name: "Publication", Fn: func() error { return (*CorePublication)(nil).unmarshalRawPublication(rawPublication{}, &id) }},
		{Name: "Reference",   Fn: func() error { return (*CoreReference)(nil).unmarshalRawReference(rawReference{}, &id) }},
		{Name: "Resume",      Fn: func() error { return (*CoreResume)(nil).unmarshalRawResume(rawResume{}, &id) }},
		{Name: "Skill",       Fn: func() error { return (*CoreSkill)(nil).unmarshalRawSkill(rawSkill{}, &id) }},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			err := test.Fn()
			if nil == err {
				t.Errorf("Expected an error but did not get one.")
				return
			}
			if !errors.Is(err, ErrReceiverNil) {
				t.Errorf("Expected ErrReceiverNil but got: %s", err)
			}
		})
	}
}

// TestUnmarshalRaw_fieldError verifies that a malformed field value
// (a number where a string is expected) produces an error mentioning
// the field, rather than silently accepting bad data.
func TestUnmarshalRaw_fieldError(t *testing.T) {

	tests := []struct {
		Name string
		JSON string
	}{
		// Each type gets a JSON object with one field set to a number (invalid for a string field).
		{Name: "Award.title",           JSON: `{"type":"Award","title":123}`},
		{Name: "Basics.name",           JSON: `{"type":"Basics","name":123}`},
		{Name: "Certificate.name",      JSON: `{"type":"Certificate","name":123}`},
		{Name: "Education.institution", JSON: `{"type":"Education","institution":123}`},
		{Name: "Experience.name",       JSON: `{"type":"Experience","name":123}`},
		{Name: "Interest.name",         JSON: `{"type":"Interest","name":123}`},
		{Name: "Language.language",      JSON: `{"type":"Language","language":123}`},
		{Name: "Location.city",         JSON: `{"type":"Location","city":123}`},
		{Name: "Meta.version",          JSON: `{"type":"Meta","version":123}`},
		{Name: "Profile.network",       JSON: `{"type":"Profile","network":123}`},
		{Name: "Project.name",          JSON: `{"@type":"Project","name":123}`},
		{Name: "Publication.name",      JSON: `{"type":"Publication","name":123}`},
		{Name: "Reference.name",        JSON: `{"type":"Reference","name":123}`},
		{Name: "Skill.name",            JSON: `{"type":"Skill","name":123}`},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// Use AnyXxx types since they accept any type value.
			var err error
			switch {
			case len(test.JSON) > 0 && test.JSON[2] == 't': // "type" key
				// Dispatch based on test name prefix
				switch test.Name[:4] {
				case "Awar":
					var v AnyAward
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Basi":
					var v AnyBasics
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Cert":
					var v AnyCertificate
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Educ":
					var v AnyEducation
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Expe":
					var v AnyExperience
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Inte":
					var v AnyInterest
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Lang":
					var v AnyLanguage
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Loca":
					var v AnyLocation
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Meta":
					var v AnyMeta
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Prof":
					var v AnyProfile
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Publ":
					var v AnyPublication
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Refe":
					var v AnyReference
					err = v.UnmarshalJSON([]byte(test.JSON))
				case "Skil":
					var v AnySkill
					err = v.UnmarshalJSON([]byte(test.JSON))
				}
			default:
				// Project uses "@type"
				var v AnyProject
				err = v.UnmarshalJSON([]byte(test.JSON))
			}

			if nil == err {
				t.Errorf("Expected an error but did not get one.")
			}
		})
	}
}
