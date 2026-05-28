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

// unmarshalAnyType dispatches UnmarshalJSON to the correct Any type.
func unmarshalAnyType(typeName string, data []byte) error {
	switch typeName {
	case "Award":
		var v AnyAward
		return v.UnmarshalJSON(data)
	case "Basics":
		var v AnyBasics
		return v.UnmarshalJSON(data)
	case "Certificate":
		var v AnyCertificate
		return v.UnmarshalJSON(data)
	case "Education":
		var v AnyEducation
		return v.UnmarshalJSON(data)
	case "Experience":
		var v AnyExperience
		return v.UnmarshalJSON(data)
	case "Interest":
		var v AnyInterest
		return v.UnmarshalJSON(data)
	case "Language":
		var v AnyLanguage
		return v.UnmarshalJSON(data)
	case "Location":
		var v AnyLocation
		return v.UnmarshalJSON(data)
	case "Meta":
		var v AnyMeta
		return v.UnmarshalJSON(data)
	case "Profile":
		var v AnyProfile
		return v.UnmarshalJSON(data)
	case "Project":
		var v AnyProject
		return v.UnmarshalJSON(data)
	case "Publication":
		var v AnyPublication
		return v.UnmarshalJSON(data)
	case "Reference":
		var v AnyReference
		return v.UnmarshalJSON(data)
	case "Skill":
		var v AnySkill
		return v.UnmarshalJSON(data)
	default:
		panic("unknown type: " + typeName)
	}
}

// TestUnmarshalRaw_fieldError verifies that every field in every type
// rejects an invalid value (a number where a string/array/object is expected),
// producing an error rather than silently accepting bad data.
//
// This is a meaningful data-integrity check: each field has its own unmarshal
// block, and a bug in any one of them could cause silent data loss.
func TestUnmarshalRaw_fieldError(t *testing.T) {

	// typeKey is "type" for most types, "@type" for Project.
	type fieldTest struct {
		TypeName string
		TypeKey  string
		Field    string
	}

	tests := []fieldTest{
		// ID field for every type (number where string expected).
		{"Award", "type", "id"},
		{"Basics", "type", "id"},
		{"Certificate", "type", "id"},
		{"Education", "type", "id"},
		{"Experience", "type", "id"},
		{"Interest", "type", "id"},
		{"Language", "type", "id"},
		{"Location", "type", "id"},
		{"Meta", "type", "id"},
		{"Profile", "type", "id"},
		{"Project", "@type", "id"},
		{"Publication", "type", "id"},
		{"Reference", "type", "id"},
		{"Skill", "type", "id"},

		// Award: 4 fields
		{"Award", "type", "awarder"},
		{"Award", "type", "date"},
		{"Award", "type", "summary"},
		{"Award", "type", "title"},

		// Basics: 9 fields (email, image, label, location, name, phone, profiles, summary, url)
		{"Basics", "type", "email"},
		{"Basics", "type", "image"},
		{"Basics", "type", "label"},
		{"Basics", "type", "location"},
		{"Basics", "type", "name"},
		{"Basics", "type", "phone"},
		{"Basics", "type", "profiles"},
		{"Basics", "type", "summary"},
		{"Basics", "type", "url"},

		// Certificate: 4 fields
		{"Certificate", "type", "date"},
		{"Certificate", "type", "issuer"},
		{"Certificate", "type", "name"},
		{"Certificate", "type", "url"},

		// Education: 8 fields
		{"Education", "type", "area"},
		{"Education", "type", "courses"},
		{"Education", "type", "endDate"},
		{"Education", "type", "institution"},
		{"Education", "type", "score"},
		{"Education", "type", "startDate"},
		{"Education", "type", "studyType"},
		{"Education", "type", "url"},

		// Experience: 10 fields
		{"Experience", "type", "description"},
		{"Experience", "type", "endDate"},
		{"Experience", "type", "highlights"},
		{"Experience", "type", "location"},
		{"Experience", "type", "name"},
		{"Experience", "type", "organization"},
		{"Experience", "type", "position"},
		{"Experience", "type", "startDate"},
		{"Experience", "type", "summary"},
		{"Experience", "type", "url"},

		// Interest: 2 fields
		{"Interest", "type", "keywords"},
		{"Interest", "type", "name"},

		// Language: 2 fields
		{"Language", "type", "fluency"},
		{"Language", "type", "language"},

		// Location: 5 fields
		{"Location", "type", "address"},
		{"Location", "type", "city"},
		{"Location", "type", "countryCode"},
		{"Location", "type", "postalCode"},
		{"Location", "type", "region"},

		// Meta: 3 fields
		{"Meta", "type", "canonical"},
		{"Meta", "type", "lastModified"},
		{"Meta", "type", "version"},

		// Profile: 3 fields
		{"Profile", "type", "network"},
		{"Profile", "type", "url"},
		{"Profile", "type", "username"},

		// Project: 10 fields (uses "@type")
		{"Project", "@type", "description"},
		{"Project", "@type", "endDate"},
		{"Project", "@type", "entity"},
		{"Project", "@type", "highlights"},
		{"Project", "@type", "keywords"},
		{"Project", "@type", "name"},
		{"Project", "@type", "roles"},
		{"Project", "@type", "startDate"},
		{"Project", "@type", "type"},
		{"Project", "@type", "url"},

		// Publication: 5 fields
		{"Publication", "type", "name"},
		{"Publication", "type", "publisher"},
		{"Publication", "type", "releaseDate"},
		{"Publication", "type", "summary"},
		{"Publication", "type", "url"},

		// Reference: 2 fields
		{"Reference", "type", "name"},
		{"Reference", "type", "reference"},

		// Skill: 3 fields
		{"Skill", "type", "keywords"},
		{"Skill", "type", "level"},
		{"Skill", "type", "name"},
	}

	for _, test := range tests {
		name := test.TypeName + "." + test.Field
		t.Run(name, func(t *testing.T) {
			// Build JSON with the type key and one malformed field (number where string expected).
			json := `{"` + test.TypeKey + `":"` + test.TypeName + `","` + test.Field + `":123}`

			err := unmarshalAnyType(test.TypeName, []byte(json))
			if nil == err {
				t.Errorf("Expected an error for malformed %s but did not get one.", test.Field)
				t.Logf("JSON: %s", json)
			}
		})
	}
}
