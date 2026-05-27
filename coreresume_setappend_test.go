package jsonresume

import (
	"reflect"
	"testing"
)

func TestCoreResume_Set(t *testing.T) {

	// Each Set* method should replace the field with a single-element slice (or single value for Basics/Meta).
	// We first Append to populate, then Set to verify it replaces rather than appends.

	tests := []struct {
		Name     string
		Do       func(r *Resume)
		Field    func(r *Resume) interface{}
		Expected interface{}
	}{
		{
			Name: "SetAwardID",
			Do: func(r *Resume) {
				r.AppendAwardID("http://example.com/award/old")
				r.SetAwardID("http://example.com/award/new")
			},
			Field:    func(r *Resume) interface{} { return r.Awards },
			Expected: []ProtoAward{SomeAwardID("http://example.com/award/new")},
		},
		{
			Name: "SetCertificateID",
			Do: func(r *Resume) {
				r.AppendCertificateID("http://example.com/cert/old")
				r.SetCertificateID("http://example.com/cert/new")
			},
			Field:    func(r *Resume) interface{} { return r.Certificates },
			Expected: []ProtoCertificate{SomeCertificateID("http://example.com/cert/new")},
		},
		{
			Name: "SetEducationID",
			Do: func(r *Resume) {
				r.AppendEducationID("http://example.com/edu/old")
				r.SetEducationID("http://example.com/edu/new")
			},
			Field:    func(r *Resume) interface{} { return r.Education },
			Expected: []ProtoEducation{SomeEducationID("http://example.com/edu/new")},
		},
		{
			Name: "SetInterestID",
			Do: func(r *Resume) {
				r.AppendInterestID("http://example.com/interest/old")
				r.SetInterestID("http://example.com/interest/new")
			},
			Field:    func(r *Resume) interface{} { return r.Interests },
			Expected: []ProtoInterest{SomeInterestID("http://example.com/interest/new")},
		},
		{
			Name: "SetLanguageID",
			Do: func(r *Resume) {
				r.AppendLanguageID("http://example.com/lang/old")
				r.SetLanguageID("http://example.com/lang/new")
			},
			Field:    func(r *Resume) interface{} { return r.Languages },
			Expected: []ProtoLanguage{SomeLanguageID("http://example.com/lang/new")},
		},
		{
			Name: "SetProjectID",
			Do: func(r *Resume) {
				r.AppendProjectID("http://example.com/project/old")
				r.SetProjectID("http://example.com/project/new")
			},
			Field:    func(r *Resume) interface{} { return r.Projects },
			Expected: []ProtoProject{SomeProjectID("http://example.com/project/new")},
		},
		{
			Name: "SetPublicationID",
			Do: func(r *Resume) {
				r.AppendPublicationID("http://example.com/pub/old")
				r.SetPublicationID("http://example.com/pub/new")
			},
			Field:    func(r *Resume) interface{} { return r.Publications },
			Expected: []ProtoPublication{SomePublicationID("http://example.com/pub/new")},
		},
		{
			Name: "SetReferenceID",
			Do: func(r *Resume) {
				r.AppendReferenceID("http://example.com/ref/old")
				r.SetReferenceID("http://example.com/ref/new")
			},
			Field:    func(r *Resume) interface{} { return r.References },
			Expected: []ProtoReference{SomeReferenceID("http://example.com/ref/new")},
		},
		{
			Name: "SetSkillID",
			Do: func(r *Resume) {
				r.AppendSkillID("http://example.com/skill/old")
				r.SetSkillID("http://example.com/skill/new")
			},
			Field:    func(r *Resume) interface{} { return r.Skills },
			Expected: []ProtoSkill{SomeSkillID("http://example.com/skill/new")},
		},
		{
			Name: "SetVolunteerID",
			Do: func(r *Resume) {
				r.AppendVolunteerID("http://example.com/vol/old")
				r.SetVolunteerID("http://example.com/vol/new")
			},
			Field:    func(r *Resume) interface{} { return r.Volunteer },
			Expected: []ProtoExperience{SomeExperienceID("http://example.com/vol/new")},
		},
		{
			Name: "SetWorkID",
			Do: func(r *Resume) {
				r.AppendWorkID("http://example.com/work/old")
				r.SetWorkID("http://example.com/work/new")
			},
			Field:    func(r *Resume) interface{} { return r.Work },
			Expected: []ProtoExperience{SomeExperienceID("http://example.com/work/new")},
		},
		{
			Name: "SetBasicsID",
			Do: func(r *Resume) {
				r.SetBasicsID("http://example.com/basics/1")
			},
			Field:    func(r *Resume) interface{} { return r.Basics },
			Expected: ProtoBasics(SomeBasicsID("http://example.com/basics/1")),
		},
		{
			Name: "SetMetaID",
			Do: func(r *Resume) {
				r.SetMetaID("http://example.com/meta/1")
			},
			Field:    func(r *Resume) interface{} { return r.Meta },
			Expected: ProtoMeta(SomeMetaID("http://example.com/meta/1")),
		},
	}

	for testNumber, test := range tests {
		var r Resume
		test.Do(&r)

		actual := test.Field(&r)
		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("For test #%d (%s), the actual value is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}
}

func TestCoreResume_Append(t *testing.T) {

	// Each Append* method should accumulate entries, not replace.

	tests := []struct {
		Name     string
		Do       func(r *Resume)
		Field    func(r *Resume) interface{}
		Expected interface{}
	}{
		{
			Name: "AppendCertificateID",
			Do: func(r *Resume) {
				r.AppendCertificateID("http://example.com/cert/1")
				r.AppendCertificateID("http://example.com/cert/2")
			},
			Field: func(r *Resume) interface{} { return r.Certificates },
			Expected: []ProtoCertificate{
				SomeCertificateID("http://example.com/cert/1"),
				SomeCertificateID("http://example.com/cert/2"),
			},
		},
		{
			Name: "AppendEducationID",
			Do: func(r *Resume) {
				r.AppendEducationID("http://example.com/edu/1")
				r.AppendEducationID("http://example.com/edu/2")
			},
			Field: func(r *Resume) interface{} { return r.Education },
			Expected: []ProtoEducation{
				SomeEducationID("http://example.com/edu/1"),
				SomeEducationID("http://example.com/edu/2"),
			},
		},
		{
			Name: "AppendInterestID",
			Do: func(r *Resume) {
				r.AppendInterestID("http://example.com/interest/1")
				r.AppendInterestID("http://example.com/interest/2")
			},
			Field: func(r *Resume) interface{} { return r.Interests },
			Expected: []ProtoInterest{
				SomeInterestID("http://example.com/interest/1"),
				SomeInterestID("http://example.com/interest/2"),
			},
		},
		{
			Name: "AppendProjectID",
			Do: func(r *Resume) {
				r.AppendProjectID("http://example.com/project/1")
				r.AppendProjectID("http://example.com/project/2")
			},
			Field: func(r *Resume) interface{} { return r.Projects },
			Expected: []ProtoProject{
				SomeProjectID("http://example.com/project/1"),
				SomeProjectID("http://example.com/project/2"),
			},
		},
		{
			Name: "AppendPublicationID",
			Do: func(r *Resume) {
				r.AppendPublicationID("http://example.com/pub/1")
				r.AppendPublicationID("http://example.com/pub/2")
			},
			Field: func(r *Resume) interface{} { return r.Publications },
			Expected: []ProtoPublication{
				SomePublicationID("http://example.com/pub/1"),
				SomePublicationID("http://example.com/pub/2"),
			},
		},
		{
			Name: "AppendReferenceID",
			Do: func(r *Resume) {
				r.AppendReferenceID("http://example.com/ref/1")
				r.AppendReferenceID("http://example.com/ref/2")
			},
			Field: func(r *Resume) interface{} { return r.References },
			Expected: []ProtoReference{
				SomeReferenceID("http://example.com/ref/1"),
				SomeReferenceID("http://example.com/ref/2"),
			},
		},
	}

	for testNumber, test := range tests {
		var r Resume
		test.Do(&r)

		actual := test.Field(&r)
		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("For test #%d (%s), the actual value is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", actual)
		}
	}
}

// TestCoreResume_SetAppend_nilReceiver verifies that all Set* and Append*
// methods on a nil *CoreResume do not panic.
func TestCoreResume_SetAppend_nilReceiver(t *testing.T) {

	var r *CoreResume // nil

	// Append methods — must not panic on nil receiver.
	r.AppendAwardID("http://example.com/a")
	r.AppendCertificateID("http://example.com/c")
	r.AppendEducationID("http://example.com/e")
	r.AppendInterestID("http://example.com/i")
	r.AppendLanguageID("http://example.com/l")
	r.AppendProjectID("http://example.com/p")
	r.AppendPublicationID("http://example.com/pub")
	r.AppendReferenceID("http://example.com/ref")
	r.AppendSkillID("http://example.com/s")
	r.AppendVolunteerID("http://example.com/v")
	r.AppendWorkID("http://example.com/w")

	// Set methods — must not panic on nil receiver.
	r.SetAwardID("http://example.com/a")
	r.SetCertificateID("http://example.com/c")
	r.SetEducationID("http://example.com/e")
	r.SetInterestID("http://example.com/i")
	r.SetLanguageID("http://example.com/l")
	r.SetProjectID("http://example.com/p")
	r.SetPublicationID("http://example.com/pub")
	r.SetReferenceID("http://example.com/ref")
	r.SetSkillID("http://example.com/s")
	r.SetVolunteerID("http://example.com/v")
	r.SetWorkID("http://example.com/w")
	r.SetBasicsID("http://example.com/b")
	r.SetMetaID("http://example.com/m")
}
