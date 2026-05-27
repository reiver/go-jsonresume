package jsonresume

import (
	"reflect"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

func TestID_ProtoNode(t *testing.T) {

	const iri = "http://example.com/thing/1"
	expectedID := jsonld.SomeID(iri)

	tests := []struct {
		Name     string
		Actual   activitypub.AnyNode
		Expected activitypub.AnyNode
	}{
		// 0
		{
			Name:   "AwardID",
			Actual: SomeAwardID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeAward),
			},
		},

		// 1
		{
			Name:   "BasicsID",
			Actual: SomeBasicsID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeBasics),
			},
		},

		// 2
		{
			Name:   "CertificateID",
			Actual: SomeCertificateID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeCertificate),
			},
		},

		// 3
		{
			Name:   "EducationID",
			Actual: SomeEducationID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeEducation),
			},
		},

		// 4
		{
			Name:   "ExperienceID",
			Actual: SomeExperienceID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeExperience),
			},
		},

		// 5
		{
			Name:   "InterestID",
			Actual: SomeInterestID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeInterest),
			},
		},

		// 6
		{
			Name:   "LanguageID",
			Actual: SomeLanguageID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeLanguage),
			},
		},

		// 7
		{
			Name:   "LocationID",
			Actual: SomeLocationID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeLocation),
			},
		},

		// 8
		{
			Name:   "MetaID",
			Actual: SomeMetaID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeMeta),
			},
		},

		// 9
		{
			Name:   "ProfileID",
			Actual: SomeProfileID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeProfile),
			},
		},

		// 10
		{
			Name:   "ProjectID",
			Actual: SomeProjectID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeProject),
			},
		},

		// 11
		{
			Name:   "PublicationID",
			Actual: SomePublicationID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypePublication),
			},
		},

		// 12
		{
			Name:   "ReferenceID",
			Actual: SomeReferenceID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeReference),
			},
		},

		// 13
		{
			Name:   "ResumeID",
			Actual: SomeResumeID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeResume),
			},
		},

		// 14
		{
			Name:   "SkillID",
			Actual: SomeSkillID(iri).ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   expectedID,
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

func TestID_ProtoEntity(t *testing.T) {

	const iri = "http://example.com/thing/1"
	expectedID := jsonld.SomeID(iri)

	tests := []struct {
		Name     string
		Actual   activitypub.AnyEntity
		Expected activitypub.AnyEntity
	}{
		{Name: "AwardID", Actual: SomeAwardID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeAward)}},
		{Name: "BasicsID", Actual: SomeBasicsID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeBasics)}},
		{Name: "CertificateID", Actual: SomeCertificateID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeCertificate)}},
		{Name: "EducationID", Actual: SomeEducationID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeEducation)}},
		{Name: "ExperienceID", Actual: SomeExperienceID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeExperience)}},
		{Name: "InterestID", Actual: SomeInterestID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeInterest)}},
		{Name: "LanguageID", Actual: SomeLanguageID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeLanguage)}},
		{Name: "LocationID", Actual: SomeLocationID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeLocation)}},
		{Name: "MetaID", Actual: SomeMetaID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeMeta)}},
		{Name: "ProfileID", Actual: SomeProfileID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeProfile)}},
		{Name: "ProjectID", Actual: SomeProjectID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeProject)}},
		{Name: "PublicationID", Actual: SomePublicationID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypePublication)}},
		{Name: "ReferenceID", Actual: SomeReferenceID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeReference)}},
		{Name: "ResumeID", Actual: SomeResumeID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeResume)}},
		{Name: "SkillID", Actual: SomeSkillID(iri).ProtoEntity(), Expected: activitypub.AnyEntity{ID: expectedID, Type: jsonld.SomeType(TypeSkill)}},
	}

	for testNumber, test := range tests {
		if !reflect.DeepEqual(test.Expected, test.Actual) {
			t.Errorf("For test #%d (%s), the actual value is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", test.Actual)
		}
	}
}

func TestID_ProtoObject(t *testing.T) {

	const iri = "http://example.com/thing/1"
	expectedID := jsonld.SomeID(iri)

	tests := []struct {
		Name     string
		Actual   activitypub.AnyObject
		Expected activitypub.AnyObject
	}{
		{Name: "AwardID", Actual: SomeAwardID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeAward)}},
		{Name: "BasicsID", Actual: SomeBasicsID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeBasics)}},
		{Name: "CertificateID", Actual: SomeCertificateID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeCertificate)}},
		{Name: "EducationID", Actual: SomeEducationID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeEducation)}},
		{Name: "ExperienceID", Actual: SomeExperienceID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeExperience)}},
		{Name: "InterestID", Actual: SomeInterestID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeInterest)}},
		{Name: "LanguageID", Actual: SomeLanguageID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeLanguage)}},
		{Name: "LocationID", Actual: SomeLocationID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeLocation)}},
		{Name: "MetaID", Actual: SomeMetaID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeMeta)}},
		{Name: "ProfileID", Actual: SomeProfileID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeProfile)}},
		{Name: "ProjectID", Actual: SomeProjectID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeProject)}},
		{Name: "PublicationID", Actual: SomePublicationID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypePublication)}},
		{Name: "ReferenceID", Actual: SomeReferenceID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeReference)}},
		{Name: "ResumeID", Actual: SomeResumeID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeResume)}},
		{Name: "SkillID", Actual: SomeSkillID(iri).ProtoObject(), Expected: activitypub.AnyObject{ID: expectedID, Type: jsonld.SomeType(TypeSkill)}},
	}

	for testNumber, test := range tests {
		if !reflect.DeepEqual(test.Expected, test.Actual) {
			t.Errorf("For test #%d (%s), the actual value is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", test.Actual)
		}
	}
}

func TestID_ProtoSpecific(t *testing.T) {

	const iri = "http://example.com/thing/1"
	expectedID := jsonld.SomeID(iri)

	tests := []struct {
		Name     string
		Actual   interface{}
		Expected interface{}
	}{
		// 0
		{
			Name:   "AwardID.ProtoAward",
			Actual: SomeAwardID(iri).ProtoAward(),
			Expected: AnyAward{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeAward),
			},
		},

		// 1
		{
			Name:   "CertificateID.ProtoCertificate",
			Actual: SomeCertificateID(iri).ProtoCertificate(),
			Expected: AnyCertificate{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeCertificate),
			},
		},

		// 2
		{
			Name:   "ProjectID.ProtoProject",
			Actual: SomeProjectID(iri).ProtoProject(),
			Expected: AnyProject{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeProject),
			},
		},

		// 3
		{
			Name:   "ResumeID.ProtoResume",
			Actual: SomeResumeID(iri).ProtoResume(),
			Expected: AnyResume{
				ID:   expectedID,
				Type: jsonld.SomeType(TypeResume),
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

func TestID_String(t *testing.T) {

	const iri = "http://example.com/thing/1"

	tests := []struct {
		Name   string
		Actual string
	}{
		// 0
		{
			Name:   "AwardID",
			Actual: SomeAwardID(iri).String(),
		},

		// 1
		{
			Name:   "BasicsID",
			Actual: SomeBasicsID(iri).String(),
		},

		// 2
		{
			Name:   "CertificateID",
			Actual: SomeCertificateID(iri).String(),
		},

		// 3
		{
			Name:   "EducationID",
			Actual: SomeEducationID(iri).String(),
		},

		// 4
		{
			Name:   "ExperienceID",
			Actual: SomeExperienceID(iri).String(),
		},

		// 5
		{
			Name:   "InterestID",
			Actual: SomeInterestID(iri).String(),
		},

		// 6
		{
			Name:   "LanguageID",
			Actual: SomeLanguageID(iri).String(),
		},

		// 7
		{
			Name:   "LocationID",
			Actual: SomeLocationID(iri).String(),
		},

		// 8
		{
			Name:   "MetaID",
			Actual: SomeMetaID(iri).String(),
		},

		// 9
		{
			Name:   "ProfileID",
			Actual: SomeProfileID(iri).String(),
		},

		// 10
		{
			Name:   "ProjectID",
			Actual: SomeProjectID(iri).String(),
		},

		// 11
		{
			Name:   "PublicationID",
			Actual: SomePublicationID(iri).String(),
		},

		// 12
		{
			Name:   "ReferenceID",
			Actual: SomeReferenceID(iri).String(),
		},

		// 13
		{
			Name:   "ResumeID",
			Actual: SomeResumeID(iri).String(),
		},

		// 14
		{
			Name:   "SkillID",
			Actual: SomeSkillID(iri).String(),
		},
	}

	for testNumber, test := range tests {
		if iri != test.Actual {
			t.Errorf("For test #%d (%s), the actual value is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED: %q", iri)
			t.Logf("ACTUAL:   %q", test.Actual)
		}
	}
}
