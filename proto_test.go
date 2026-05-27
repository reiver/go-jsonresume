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

	tests := []struct {
		Name     string
		Actual   activitypub.AnyEntity
		Expected activitypub.AnyEntity
	}{
		// Types that propagate Name to CoreEntity:

		{
			Name:   "Basics",
			Actual: Basics{ID: id, CoreBasics: CoreBasics{Name: name}}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeBasics),
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},
		{
			Name:   "Certificate",
			Actual: Certificate{ID: id, CoreCertificate: CoreCertificate{Name: name}}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeCertificate),
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},
		{
			Name:   "Experience",
			Actual: Experience{ID: id, CoreExperience: CoreExperience{Name: name}}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeExperience),
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},
		{
			Name:   "Interest",
			Actual: Interest{ID: id, CoreInterest: CoreInterest{Name: name}}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeInterest),
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},
		{
			Name:   "Project",
			Actual: Project{ID: id, CoreProject: CoreProject{Name: name}}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeProject),
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},
		{
			Name:   "Publication",
			Actual: Publication{ID: id, CorePublication: CorePublication{Name: name}}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypePublication),
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},
		{
			Name:   "Reference",
			Actual: Reference{ID: id, CoreReference: CoreReference{Name: name}}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeReference),
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},

		// Types that do NOT propagate Name (ID+Type only):

		{
			Name:   "Award",
			Actual: Award{ID: id}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeAward),
			},
		},
		{
			Name:   "Education",
			Actual: Education{ID: id}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeEducation),
			},
		},
		{
			Name:   "Language",
			Actual: Language{ID: id}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeLanguage),
			},
		},
		{
			Name:   "Location",
			Actual: Location{ID: id}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeLocation),
			},
		},
		{
			Name:   "Meta",
			Actual: Meta{ID: id}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeMeta),
			},
		},
		{
			Name:   "Profile",
			Actual: Profile{ID: id}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeProfile),
			},
		},
		{
			Name:   "Resume",
			Actual: Resume{ID: id}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
				ID:   id,
				Type: jsonld.SomeType(TypeResume),
			},
		},
		{
			Name:   "Skill",
			Actual: Skill{ID: id}.ProtoEntity(),
			Expected: activitypub.AnyEntity{
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

func TestProtoObject(t *testing.T) {

	id := jsonld.SomeID("http://example.com/thing/1")
	name := nul.Something("Test Name")
	summary := nul.Something("A summary")
	url := []activitypub.ProtoLink{activitypub.HRef("https://example.com")}
	image := []activitypub.ProtoImageOrProtoLink{activitypub.HRef("https://example.com/photo.jpg")}

	tests := []struct {
		Name     string
		Actual   activitypub.AnyObject
		Expected activitypub.AnyObject
	}{
		// Award: Summary only
		{
			Name:   "Award",
			Actual: Award{ID: id, CoreAward: CoreAward{Summary: summary}}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeAward),
				CoreObject: activitypub.CoreObject{
					Summary: summary,
				},
			},
		},

		// Basics: Name, Image, Summary, URL
		{
			Name: "Basics",
			Actual: Basics{ID: id, CoreBasics: CoreBasics{
				Name: name, Image: image, Summary: summary, URL: url,
			}}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeBasics),
				CoreEntity: activitypub.CoreEntity{Name: name},
				CoreObject: activitypub.CoreObject{
					Image:   image,
					Summary: summary,
					URL:     url,
				},
			},
		},

		// Certificate: Name, URL
		{
			Name:   "Certificate",
			Actual: Certificate{ID: id, CoreCertificate: CoreCertificate{Name: name, URL: url}}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeCertificate),
				CoreEntity: activitypub.CoreEntity{Name: name},
				CoreObject: activitypub.CoreObject{URL: url},
			},
		},

		// Education: URL only
		{
			Name:   "Education",
			Actual: Education{ID: id, CoreEducation: CoreEducation{URL: url}}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeEducation),
				CoreObject: activitypub.CoreObject{URL: url},
			},
		},

		// Experience: Name, Summary, URL
		{
			Name: "Experience",
			Actual: Experience{ID: id, CoreExperience: CoreExperience{
				Name: name, Summary: summary, URL: url,
			}}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeExperience),
				CoreEntity: activitypub.CoreEntity{Name: name},
				CoreObject: activitypub.CoreObject{
					Summary: summary,
					URL:     url,
				},
			},
		},

		// Interest: Name only
		{
			Name:   "Interest",
			Actual: Interest{ID: id, CoreInterest: CoreInterest{Name: name}}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeInterest),
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},

		// Language: ID+Type only
		{
			Name:   "Language",
			Actual: Language{ID: id}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeLanguage),
			},
		},

		// Location: ID+Type only
		{
			Name:   "Location",
			Actual: Location{ID: id}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeLocation),
			},
		},

		// Meta: ID+Type only
		{
			Name:   "Meta",
			Actual: Meta{ID: id}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeMeta),
			},
		},

		// Profile: URL only
		{
			Name:   "Profile",
			Actual: Profile{ID: id, CoreProfile: CoreProfile{URL: url}}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeProfile),
				CoreObject: activitypub.CoreObject{URL: url},
			},
		},

		// Project: Name, URL
		{
			Name:   "Project",
			Actual: Project{ID: id, CoreProject: CoreProject{Name: name, URL: url}}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeProject),
				CoreEntity: activitypub.CoreEntity{Name: name},
				CoreObject: activitypub.CoreObject{URL: url},
			},
		},

		// Publication: Name, Summary, URL
		{
			Name: "Publication",
			Actual: Publication{ID: id, CorePublication: CorePublication{
				Name: name, Summary: summary, URL: url,
			}}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypePublication),
				CoreEntity: activitypub.CoreEntity{Name: name},
				CoreObject: activitypub.CoreObject{
					Summary: summary,
					URL:     url,
				},
			},
		},

		// Reference: Name only
		{
			Name:   "Reference",
			Actual: Reference{ID: id, CoreReference: CoreReference{Name: name}}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeReference),
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},

		// Resume: ID+Type only
		{
			Name:   "Resume",
			Actual: Resume{ID: id}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID:   id,
				Type: jsonld.SomeType(TypeResume),
			},
		},

		// Skill: ID+Type only
		{
			Name:   "Skill",
			Actual: Skill{ID: id}.ProtoObject(),
			Expected: activitypub.AnyObject{
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

	// Basics
	{
		core := CoreBasics{
			Name:    nul.Something("Joe Blow"),
			Label:   activitypub.SomeStrings("Programmer"),
			EMail:   activitypub.SomeStrings("joe@example.com"),
			Phone:   activitypub.SomeStrings("555-1234"),
			Summary: nul.Something("A developer"),
			URL:     []activitypub.ProtoLink{activitypub.HRef("https://joe.example.com")},
			Image:   []activitypub.ProtoImageOrProtoLink{activitypub.HRef("https://joe.example.com/photo.jpg")},
		}

		actual := Basics{ID: id, CoreBasics: core}.ProtoBasics()

		expected := AnyBasics{
			ID:         id,
			Type:       jsonld.SomeType(TypeBasics),
			CoreBasics: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Basics.ProtoBasics(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Certificate
	{
		core := CoreCertificate{
			Name:   nul.Something("AWS Solutions Architect"),
			Date:   nul.Something("2024-01-15"),
			Issuer: nul.Something("Amazon"),
			URL:    []activitypub.ProtoLink{activitypub.HRef("https://aws.amazon.com/cert")},
		}

		actual := Certificate{ID: id, CoreCertificate: core}.ProtoCertificate()

		expected := AnyCertificate{
			ID:              id,
			Type:            jsonld.SomeType(TypeCertificate),
			CoreCertificate: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Certificate.ProtoCertificate(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Education
	{
		core := CoreEducation{
			Institution: nul.Something("MIT"),
			Area:        activitypub.SomeStrings("Computer Science"),
			StudyType:   activitypub.SomeString("Bachelor"),
			StartDate:   nul.Something("2018"),
			EndDate:     nul.Something("2022"),
			Score:       nul.Something("3.9"),
			URL:         []activitypub.ProtoLink{activitypub.HRef("https://mit.edu")},
		}

		actual := Education{ID: id, CoreEducation: core}.ProtoEducation()

		expected := AnyEducation{
			ID:            id,
			Type:          jsonld.SomeType(TypeEducation),
			CoreEducation: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Education.ProtoEducation(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Interest
	{
		core := CoreInterest{
			Name:     nul.Something("Open Source"),
			Keywords: activitypub.SomeStrings("Linux", "FreeBSD"),
		}

		actual := Interest{ID: id, CoreInterest: core}.ProtoInterest()

		expected := AnyInterest{
			ID:           id,
			Type:         jsonld.SomeType(TypeInterest),
			CoreInterest: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Interest.ProtoInterest(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Language
	{
		core := CoreLanguage{
			Language: nul.Something("English"),
			Fluency:  nul.Something("Native"),
		}

		actual := Language{ID: id, CoreLanguage: core}.ProtoLanguage()

		expected := AnyLanguage{
			ID:           id,
			Type:         jsonld.SomeType(TypeLanguage),
			CoreLanguage: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Language.ProtoLanguage(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Location
	{
		core := CoreLocation{
			Address:     nul.Something("123 Main St"),
			City:        nul.Something("Springfield"),
			CountryCode: nul.Something("US"),
			PostalCode:  nul.Something("62704"),
			Region:      nul.Something("Illinois"),
		}

		actual := Location{ID: id, CoreLocation: core}.ProtoLocation()

		expected := AnyLocation{
			ID:           id,
			Type:         jsonld.SomeType(TypeLocation),
			CoreLocation: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Location.ProtoLocation(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Meta
	{
		core := CoreMeta{
			Canonical:    nul.Something("https://example.com/resume.json"),
			Version:      nul.Something("v1.0.0"),
			LastModified: nul.Something("2024-05-21T00:00:00Z"),
		}

		actual := Meta{ID: id, CoreMeta: core}.ProtoMeta()

		expected := AnyMeta{
			ID:       id,
			Type:     jsonld.SomeType(TypeMeta),
			CoreMeta: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Meta.ProtoMeta(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Profile
	{
		core := CoreProfile{
			Network:  nul.Something("GitHub"),
			UserName: nul.Something("joeblow"),
			URL:      []activitypub.ProtoLink{activitypub.HRef("https://github.com/joeblow")},
		}

		actual := Profile{ID: id, CoreProfile: core}.ProtoProfile()

		expected := AnyProfile{
			ID:          id,
			Type:        jsonld.SomeType(TypeProfile),
			CoreProfile: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Profile.ProtoProfile(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Publication
	{
		core := CorePublication{
			Name:        nul.Something("My Paper"),
			Publisher:   nul.Something("ACM"),
			ReleaseDate: nul.Something("2023-06-15"),
			Summary:     nul.Something("A research paper"),
			URL:         []activitypub.ProtoLink{activitypub.HRef("https://acm.org/paper")},
		}

		actual := Publication{ID: id, CorePublication: core}.ProtoPublication()

		expected := AnyPublication{
			ID:              id,
			Type:            jsonld.SomeType(TypePublication),
			CorePublication: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Publication.ProtoPublication(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Reference
	{
		core := CoreReference{
			Name:      nul.Something("Jane Smith"),
			Reference: nul.Something("Joe is a great colleague"),
		}

		actual := Reference{ID: id, CoreReference: core}.ProtoReference()

		expected := AnyReference{
			ID:            id,
			Type:          jsonld.SomeType(TypeReference),
			CoreReference: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Reference.ProtoReference(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}

	// Skill
	{
		core := CoreSkill{
			Name:     nul.Something("Go"),
			Level:    nul.Something("Advanced"),
			Keywords: activitypub.SomeStrings("concurrency", "testing"),
		}

		actual := Skill{ID: id, CoreSkill: core}.ProtoSkill()

		expected := AnySkill{
			ID:        id,
			Type:      jsonld.SomeType(TypeSkill),
			CoreSkill: core,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For Skill.ProtoSkill(), the actual value is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}
}
