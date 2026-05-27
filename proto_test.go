package jsonresume

import (
	"reflect"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

func TestProtoNode(t *testing.T) {

	id := jsonld.SomeID("http://example.com/thing/1")

	tests := []struct {
		Name     string
		Actual   activitypub.AnyNode
		Expected activitypub.AnyNode
	}{
		// 0
		{
			Name:   "Award",
			Actual: Award{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeAward),
			},
		},

		// 1
		{
			Name:   "Basics",
			Actual: Basics{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeBasics),
			},
		},

		// 2
		{
			Name:   "Certificate",
			Actual: Certificate{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeCertificate),
			},
		},

		// 3
		{
			Name:   "Education",
			Actual: Education{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeEducation),
			},
		},

		// 4
		{
			Name:   "Experience",
			Actual: Experience{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeExperience),
			},
		},

		// 5
		{
			Name:   "Interest",
			Actual: Interest{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeInterest),
			},
		},

		// 6
		{
			Name:   "Language",
			Actual: Language{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeLanguage),
			},
		},

		// 7
		{
			Name:   "Location",
			Actual: Location{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeLocation),
			},
		},

		// 8
		{
			Name:   "Meta",
			Actual: Meta{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeMeta),
			},
		},

		// 9
		{
			Name:   "Profile",
			Actual: Profile{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeProfile),
			},
		},

		// 10
		{
			Name:   "Project",
			Actual: Project{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeProject),
			},
		},

		// 11
		{
			Name:   "Publication",
			Actual: Publication{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypePublication),
			},
		},

		// 12
		{
			Name:   "Reference",
			Actual: Reference{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeReference),
			},
		},

		// 13
		{
			Name:   "Resume",
			Actual: Resume{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeResume),
			},
		},

		// 14
		{
			Name:   "Skill",
			Actual: Skill{ID: id}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: jsonld.SomeType(TypeSkill),
			},
		},
	}

	for testNumber, test := range tests {
		if !reflect.DeepEqual(test.Expected, test.Actual) {
			t.Errorf("For test #%d (%s), the actual value is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", test.Actual)
		}
	}
}

func TestProtoEntity(t *testing.T) {

	id := jsonld.SomeID("http://example.com/thing/1")
	name := nul.Something("Test Name")

	// Test a type that propagates Name (Experience) and one that doesn't (Award).

	{
		actual := Experience{
			ID: id,
			CoreExperience: CoreExperience{
				Name: name,
			},
		}.ProtoEntity()

		expected := activitypub.AnyEntity{
			ID:   id,
			Type: jsonld.SomeType(TypeExperience),
			CoreEntity: activitypub.CoreEntity{
				Name: name,
			},
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Experience.ProtoEntity(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	{
		actual := Award{ID: id}.ProtoEntity()

		expected := activitypub.AnyEntity{
			ID:   id,
			Type: jsonld.SomeType(TypeAward),
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Award.ProtoEntity(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}
}

func TestProtoObject(t *testing.T) {

	id := jsonld.SomeID("http://example.com/thing/1")
	name := nul.Something("Test Name")
	summary := nul.Something("A summary")
	url := []activitypub.ProtoLink{activitypub.HRef("https://example.com")}

	// Test Experience — propagates Name, Summary, URL.
	{
		actual := Experience{
			ID: id,
			CoreExperience: CoreExperience{
				Name:    name,
				Summary: summary,
				URL:     url,
			},
		}.ProtoObject()

		expected := activitypub.AnyObject{
			ID:   id,
			Type: jsonld.SomeType(TypeExperience),
			CoreEntity: activitypub.CoreEntity{
				Name: name,
			},
			CoreObject: activitypub.CoreObject{
				Summary: summary,
				URL:     url,
			},
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Experience.ProtoObject(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Test Award — propagates Summary only (no Name, no URL).
	{
		actual := Award{
			ID: id,
			CoreAward: CoreAward{
				Summary: summary,
			},
		}.ProtoObject()

		expected := activitypub.AnyObject{
			ID:   id,
			Type: jsonld.SomeType(TypeAward),
			CoreObject: activitypub.CoreObject{
				Summary: summary,
			},
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Award.ProtoObject(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Test Project — propagates Name and URL (no Summary).
	{
		actual := Project{
			ID: id,
			CoreProject: CoreProject{
				Name: name,
				URL:  url,
			},
		}.ProtoObject()

		expected := activitypub.AnyObject{
			ID:   id,
			Type: jsonld.SomeType(TypeProject),
			CoreEntity: activitypub.CoreEntity{
				Name: name,
			},
			CoreObject: activitypub.CoreObject{
				URL: url,
			},
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Project.ProtoObject(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}
}

func TestProtoSpecific(t *testing.T) {

	id := jsonld.SomeID("http://example.com/thing/1")

	// Test that the type-specific Proto method copies the full Core struct.

	// Award
	{
		core := CoreAward{
			Awarder: nul.Something("SuperCo"),
			Date:    nul.Something("2024-05-21"),
			Title:   nul.Something("Best Employee"),
			Summary: nul.Something("Awarded for excellence"),
		}

		actual := Award{ID: id, CoreAward: core}.ProtoAward()

		expected := AnyAward{
			ID:        id,
			Type:      jsonld.SomeType(TypeAward),
			CoreAward: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Award.ProtoAward(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Experience
	{
		core := CoreExperience{
			Name:     nul.Something("Acme Corp"),
			Position: activitypub.SomeStrings("Engineer"),
			Summary:  nul.Something("Did things"),
		}

		actual := Experience{ID: id, CoreExperience: core}.ProtoExperience()

		expected := AnyExperience{
			ID:             id,
			Type:           jsonld.SomeType(TypeExperience),
			CoreExperience: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Experience.ProtoExperience(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Project
	{
		core := CoreProject{
			Name:        nul.Something("Microdon"),
			Description: nul.Something("An AP server"),
			ProjectType: nul.Something("application"),
		}

		actual := Project{ID: id, CoreProject: core}.ProtoProject()

		expected := AnyProject{
			ID:          id,
			Type:        jsonld.SomeType(TypeProject),
			CoreProject: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Project.ProtoProject(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Resume
	{
		core := CoreResume{
			Schema: nul.Something("https://jsonresume.org/schema"),
		}

		actual := Resume{ID: id, CoreResume: core}.ProtoResume()

		expected := AnyResume{
			ID:         id,
			Type:       jsonld.SomeType(TypeResume),
			CoreResume: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Resume.ProtoResume(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}
}
