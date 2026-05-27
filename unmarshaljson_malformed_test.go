package jsonresume_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

func TestUnmarshalJSON_malformedJSON(t *testing.T) {

	inputs := []struct {
		Name string
		JSON string
	}{
		{
			Name: "truncated object",
			JSON: `{"type":"Award","title":"Be`,
		},
		{
			Name: "missing closing brace",
			JSON: `{"type":"Award"`,
		},
		{
			Name: "bare string",
			JSON: `hello`,
		},
		{
			Name: "empty string",
			JSON: ``,
		},
		{
			Name: "bare number",
			JSON: `42`,
		},
	}

	types := []struct {
		Name      string
		Unmarshal func([]byte) error
	}{
		// 0
		{
			Name: "Award",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Award
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 1
		{
			Name: "Basics",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Basics
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 2
		{
			Name: "Certificate",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Certificate
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 3
		{
			Name: "Education",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Education
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 4
		{
			Name: "Experience",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Experience
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 5
		{
			Name: "Interest",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Interest
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 6
		{
			Name: "Language",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Language
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 7
		{
			Name: "Location",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Location
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 8
		{
			Name: "Meta",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Meta
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 9
		{
			Name: "Profile",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Profile
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 10
		{
			Name: "Project",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Project
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 11
		{
			Name: "Publication",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Publication
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 12
		{
			Name: "Reference",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Reference
				return jsonld.Unmarshal(b, &v)
			},
		},

		// 13
		{
			Name: "Skill",
			Unmarshal: func(b []byte) error {
				var v jsonresume.Skill
				return jsonld.Unmarshal(b, &v)
			},
		},
	}

	for _, typ := range types {
		for _, input := range inputs {
			t.Run(typ.Name+"/"+input.Name, func(t *testing.T) {
				err := typ.Unmarshal([]byte(input.JSON))
				if nil == err {
					t.Errorf("Expected an error but did not get one.")
					t.Logf("JSON: %s", input.JSON)
				}
			})
		}
	}
}
