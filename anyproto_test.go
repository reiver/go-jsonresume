package jsonresume

import (
	"reflect"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

func TestAny_ProtoNode(t *testing.T) {

	id := jsonld.SomeID("http://example.com/thing/1")
	typ := jsonld.SomeType("Thing")

	tests := []struct {
		Name     string
		Actual   activitypub.AnyNode
		Expected activitypub.AnyNode
	}{
		// 0
		{
			Name:   "AnyAward",
			Actual: AnyAward{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 1
		{
			Name:   "AnyBasics",
			Actual: AnyBasics{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 2
		{
			Name:   "AnyCertificate",
			Actual: AnyCertificate{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 3
		{
			Name:   "AnyEducation",
			Actual: AnyEducation{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 4
		{
			Name:   "AnyExperience",
			Actual: AnyExperience{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 5
		{
			Name:   "AnyInterest",
			Actual: AnyInterest{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 6
		{
			Name:   "AnyLanguage",
			Actual: AnyLanguage{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 7
		{
			Name:   "AnyLocation",
			Actual: AnyLocation{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 8
		{
			Name:   "AnyMeta",
			Actual: AnyMeta{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 9
		{
			Name:   "AnyProfile",
			Actual: AnyProfile{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 10
		{
			Name:   "AnyProject",
			Actual: AnyProject{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 11
		{
			Name:   "AnyPublication",
			Actual: AnyPublication{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 12
		{
			Name:   "AnyReference",
			Actual: AnyReference{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 13
		{
			Name:   "AnyResume",
			Actual: AnyResume{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
			},
		},

		// 14
		{
			Name:   "AnySkill",
			Actual: AnySkill{ID: id, Type: typ}.ProtoNode(),
			Expected: activitypub.AnyNode{
				ID:   id,
				Type: typ,
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

func TestAny_ProtoEntity(t *testing.T) {

	id := jsonld.SomeID("http://example.com/thing/1")
	typ := jsonld.SomeType("Thing")
	name := nul.Something("Test Name")

	// Types that propagate Name to ProtoEntity.
	{
		tests := []struct {
			Name   string
			Actual activitypub.AnyEntity
		}{
			{Name: "AnyBasics",      Actual: AnyBasics{ID: id, Type: typ, CoreBasics: CoreBasics{Name: name}}.ProtoEntity()},
			{Name: "AnyCertificate", Actual: AnyCertificate{ID: id, Type: typ, CoreCertificate: CoreCertificate{Name: name}}.ProtoEntity()},
			{Name: "AnyExperience",  Actual: AnyExperience{ID: id, Type: typ, CoreExperience: CoreExperience{Name: name}}.ProtoEntity()},
			{Name: "AnyInterest",    Actual: AnyInterest{ID: id, Type: typ, CoreInterest: CoreInterest{Name: name}}.ProtoEntity()},
			{Name: "AnyProject",     Actual: AnyProject{ID: id, Type: typ, CoreProject: CoreProject{Name: name}}.ProtoEntity()},
			{Name: "AnyPublication", Actual: AnyPublication{ID: id, Type: typ, CorePublication: CorePublication{Name: name}}.ProtoEntity()},
			{Name: "AnyReference",   Actual: AnyReference{ID: id, Type: typ, CoreReference: CoreReference{Name: name}}.ProtoEntity()},
		{Name: "AnySkill",      Actual: AnySkill{ID: id, Type: typ, CoreSkill: CoreSkill{Name: name}}.ProtoEntity()},
		}

		expected := activitypub.AnyEntity{
			ID:   id,
			Type: typ,
			CoreEntity: activitypub.CoreEntity{
				Name: name,
			},
		}

		for testNumber, test := range tests {
			if !reflect.DeepEqual(expected, test.Actual) {
				t.Errorf("For test #%d (%s), ProtoEntity() did not propagate Name.", testNumber, test.Name)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", test.Actual)
			}
		}
	}

	// Types that do NOT propagate Name to ProtoEntity (ID+Type only).
	{
		tests := []struct {
			Name   string
			Actual activitypub.AnyEntity
		}{
			{Name: "AnyAward",    Actual: AnyAward{ID: id, Type: typ}.ProtoEntity()},
			{Name: "AnyEducation", Actual: AnyEducation{ID: id, Type: typ}.ProtoEntity()},
			{Name: "AnyLanguage", Actual: AnyLanguage{ID: id, Type: typ}.ProtoEntity()},
			{Name: "AnyLocation", Actual: AnyLocation{ID: id, Type: typ}.ProtoEntity()},
			{Name: "AnyMeta",     Actual: AnyMeta{ID: id, Type: typ}.ProtoEntity()},
			{Name: "AnyProfile",  Actual: AnyProfile{ID: id, Type: typ}.ProtoEntity()},
			{Name: "AnyResume",   Actual: AnyResume{ID: id, Type: typ}.ProtoEntity()},
		}

		expected := activitypub.AnyEntity{
			ID:   id,
			Type: typ,
		}

		for testNumber, test := range tests {
			if !reflect.DeepEqual(expected, test.Actual) {
				t.Errorf("For test #%d (%s), ProtoEntity() should only have ID+Type.", testNumber, test.Name)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", test.Actual)
			}
		}
	}
}

func TestAny_ProtoObject(t *testing.T) {

	id := jsonld.SomeID("http://example.com/thing/1")
	typ := jsonld.SomeType("Thing")
	name := nul.Something("Test Name")
	summary := nul.Something("A summary")
	url := []activitypub.ProtoLink{activitypub.HRef("https://example.com")}
	image := []activitypub.ProtoImageOrProtoLink{activitypub.HRef("https://example.com/photo.jpg")}

	tests := []struct {
		Name     string
		Actual   activitypub.AnyObject
		Expected activitypub.AnyObject
	}{
		// AnyAward — propagates Summary only (no Name, no URL).
		{
			Name: "AnyAward",
			Actual: AnyAward{
				ID: id, Type: typ,
				CoreAward: CoreAward{Summary: summary},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreObject: activitypub.CoreObject{Summary: summary},
			},
		},

		// AnyBasics — propagates Name, Image, Summary, URL.
		{
			Name: "AnyBasics",
			Actual: AnyBasics{
				ID: id, Type: typ,
				CoreBasics: CoreBasics{Name: name, Image: image, Summary: summary, URL: url},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreEntity: activitypub.CoreEntity{Name: name},
				CoreObject: activitypub.CoreObject{Image: image, Summary: summary, URL: url},
			},
		},

		// AnyCertificate — propagates Name, URL.
		{
			Name: "AnyCertificate",
			Actual: AnyCertificate{
				ID: id, Type: typ,
				CoreCertificate: CoreCertificate{Name: name, URL: url},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreEntity: activitypub.CoreEntity{Name: name},
				CoreObject: activitypub.CoreObject{URL: url},
			},
		},

		// AnyEducation — propagates URL only (no Name).
		{
			Name: "AnyEducation",
			Actual: AnyEducation{
				ID: id, Type: typ,
				CoreEducation: CoreEducation{URL: url},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreObject: activitypub.CoreObject{URL: url},
			},
		},

		// AnyExperience — propagates Name, Summary, URL.
		{
			Name: "AnyExperience",
			Actual: AnyExperience{
				ID: id, Type: typ,
				CoreExperience: CoreExperience{Name: name, Summary: summary, URL: url},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreEntity: activitypub.CoreEntity{Name: name},
				CoreObject: activitypub.CoreObject{Summary: summary, URL: url},
			},
		},

		// AnyInterest — propagates Name only (no CoreObject fields).
		{
			Name: "AnyInterest",
			Actual: AnyInterest{
				ID: id, Type: typ,
				CoreInterest: CoreInterest{Name: name},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},

		// AnyLanguage — ID+Type only (no propagated fields).
		{
			Name:     "AnyLanguage",
			Actual:   AnyLanguage{ID: id, Type: typ}.ProtoObject(),
			Expected: activitypub.AnyObject{ID: id, Type: typ},
		},

		// AnyLocation — ID+Type only (no propagated fields).
		{
			Name:     "AnyLocation",
			Actual:   AnyLocation{ID: id, Type: typ}.ProtoObject(),
			Expected: activitypub.AnyObject{ID: id, Type: typ},
		},

		// AnyMeta — ID+Type only (no propagated fields).
		{
			Name:     "AnyMeta",
			Actual:   AnyMeta{ID: id, Type: typ}.ProtoObject(),
			Expected: activitypub.AnyObject{ID: id, Type: typ},
		},

		// AnyProfile — propagates URL only (no Name).
		{
			Name: "AnyProfile",
			Actual: AnyProfile{
				ID: id, Type: typ,
				CoreProfile: CoreProfile{URL: url},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreObject: activitypub.CoreObject{URL: url},
			},
		},

		// AnyProject — propagates Name, URL.
		{
			Name: "AnyProject",
			Actual: AnyProject{
				ID: id, Type: typ,
				CoreProject: CoreProject{Name: name, URL: url},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreEntity: activitypub.CoreEntity{Name: name},
				CoreObject: activitypub.CoreObject{URL: url},
			},
		},

		// AnyPublication — propagates Name, Summary, URL.
		{
			Name: "AnyPublication",
			Actual: AnyPublication{
				ID: id, Type: typ,
				CorePublication: CorePublication{Name: name, Summary: summary, URL: url},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreEntity: activitypub.CoreEntity{Name: name},
				CoreObject: activitypub.CoreObject{Summary: summary, URL: url},
			},
		},

		// AnyReference — propagates Name only (no CoreObject fields).
		{
			Name: "AnyReference",
			Actual: AnyReference{
				ID: id, Type: typ,
				CoreReference: CoreReference{Name: name},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},

		// AnyResume — ID+Type only (no propagated fields).
		{
			Name:     "AnyResume",
			Actual:   AnyResume{ID: id, Type: typ}.ProtoObject(),
			Expected: activitypub.AnyObject{ID: id, Type: typ},
		},

		// AnySkill — propagates Name only (no CoreObject fields).
		{
			Name: "AnySkill",
			Actual: AnySkill{
				ID: id, Type: typ,
				CoreSkill: CoreSkill{Name: name},
			}.ProtoObject(),
			Expected: activitypub.AnyObject{
				ID: id, Type: typ,
				CoreEntity: activitypub.CoreEntity{Name: name},
			},
		},
	}

	for testNumber, test := range tests {
		if !reflect.DeepEqual(test.Expected, test.Actual) {
			t.Errorf("For test #%d (%s), ProtoObject() did not return expected value.", testNumber, test.Name)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", test.Actual)
		}
	}
}

func TestAny_ProtoSpecific(t *testing.T) {

	id := jsonld.SomeID("http://example.com/thing/1")
	typ := jsonld.SomeType("Award")

	// All type-specific Proto methods on Any types return receiver.

	{
		v := AnyAward{ID: id, Type: typ, CoreAward: CoreAward{Title: nul.Something("Best")}}
		if !reflect.DeepEqual(v, v.ProtoAward()) {
			t.Errorf("AnyAward.ProtoAward() did not return receiver")
		}
	}
	{
		v := AnyBasics{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoBasics()) {
			t.Errorf("AnyBasics.ProtoBasics() did not return receiver")
		}
	}
	{
		v := AnyCertificate{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoCertificate()) {
			t.Errorf("AnyCertificate.ProtoCertificate() did not return receiver")
		}
	}
	{
		v := AnyEducation{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoEducation()) {
			t.Errorf("AnyEducation.ProtoEducation() did not return receiver")
		}
	}
	{
		v := AnyExperience{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoExperience()) {
			t.Errorf("AnyExperience.ProtoExperience() did not return receiver")
		}
	}
	{
		v := AnyInterest{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoInterest()) {
			t.Errorf("AnyInterest.ProtoInterest() did not return receiver")
		}
	}
	{
		v := AnyLanguage{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoLanguage()) {
			t.Errorf("AnyLanguage.ProtoLanguage() did not return receiver")
		}
	}
	{
		v := AnyLocation{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoLocation()) {
			t.Errorf("AnyLocation.ProtoLocation() did not return receiver")
		}
	}
	{
		v := AnyMeta{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoMeta()) {
			t.Errorf("AnyMeta.ProtoMeta() did not return receiver")
		}
	}
	{
		v := AnyProfile{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoProfile()) {
			t.Errorf("AnyProfile.ProtoProfile() did not return receiver")
		}
	}
	{
		v := AnyProject{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoProject()) {
			t.Errorf("AnyProject.ProtoProject() did not return receiver")
		}
	}
	{
		v := AnyPublication{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoPublication()) {
			t.Errorf("AnyPublication.ProtoPublication() did not return receiver")
		}
	}
	{
		v := AnyReference{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoReference()) {
			t.Errorf("AnyReference.ProtoReference() did not return receiver")
		}
	}
	{
		v := AnyResume{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoResume()) {
			t.Errorf("AnyResume.ProtoResume() did not return receiver")
		}
	}
	{
		v := AnySkill{ID: id, Type: typ}
		if !reflect.DeepEqual(v, v.ProtoSkill()) {
			t.Errorf("AnySkill.ProtoSkill() did not return receiver")
		}
	}
}
