package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-jsonresume"
)

// TestResume_roundtrip verifies that a fully populated Resume survives
// marshal (via jsonld.Marshal) → unmarshal (via jsonld.Unmarshal into AnyResume)
// without losing data in the scalar fields.
func TestResume_roundtrip(t *testing.T) {

	var cv jsonresume.Resume

	cv.ID = jsonld.SomeID("http://example.com/resume/1")

	cv.Schema = nul.Something("https://jsonresume.org/schema")

	cv.Basics = jsonresume.Basics{
		CoreBasics: jsonresume.CoreBasics{
			EMail:   activitypub.SomeString("joe@example.com"),
			Label:   activitypub.SomeString("Programmer"),
			Name:    nul.Something("Joe Blow"),
			Phone:   activitypub.SomeString("555-1234"),
			Summary: nul.Something("An experienced developer."),
			Image:   []activitypub.ProtoImageOrProtoLink{activitypub.HRef("https://example.com/photo.jpg")},
			URL:     []activitypub.ProtoLink{activitypub.HRef("https://joe.example.com")},
			Location: []jsonresume.ProtoLocation{
				jsonresume.Location{
					CoreLocation: jsonresume.CoreLocation{
						Address:     nul.Something("123 Main St"),
						City:        nul.Something("Springfield"),
						CountryCode: nul.Something("US"),
						PostalCode:  nul.Something("62704"),
						Region:      nul.Something("Illinois"),
					},
				},
			},
			Profiles: []jsonresume.ProtoProfile{
				jsonresume.Profile{
					CoreProfile: jsonresume.CoreProfile{
						Network:  nul.Something("Mastodon"),
						UserName: nul.Something("joeblow"),
						URL:      []activitypub.ProtoLink{activitypub.HRef("https://mastodon.social/@joeblow")},
					},
				},
			},
		},
	}

	cv.Awards = []jsonresume.ProtoAward{
		jsonresume.Award{
			CoreAward: jsonresume.CoreAward{
				Title:   nul.Something("Best Employee"),
				Date:    nul.Something("2024-01-15"),
				Awarder: nul.Something("ACME Corp"),
				Summary: nul.Something("Outstanding performance."),
			},
		},
	}

	cv.Certificates = []jsonresume.ProtoCertificate{
		jsonresume.Certificate{
			CoreCertificate: jsonresume.CoreCertificate{
				Name:   nul.Something("AWS Solutions Architect"),
				Date:   nul.Something("2023-06-01"),
				Issuer: nul.Something("Amazon"),
				URL:    []activitypub.ProtoLink{activitypub.HRef("https://aws.amazon.com/cert/123")},
			},
		},
	}

	cv.Education = []jsonresume.ProtoEducation{
		jsonresume.Education{
			CoreEducation: jsonresume.CoreEducation{
				Institution: nul.Something("MIT"),
				Area:        activitypub.SomeString("Computer Science"),
				StudyType:   activitypub.SomeString("Bachelor"),
				StartDate:   nul.Something("2010-09-01"),
				EndDate:     nul.Something("2014-06-15"),
				Score:       nul.Something("3.9"),
				Courses:     activitypub.SomeStrings("CS101", "CS201"),
				URL:         []activitypub.ProtoLink{activitypub.HRef("https://mit.edu")},
			},
		},
	}

	cv.Work = []jsonresume.ProtoExperience{
		jsonresume.Experience{
			CoreExperience: jsonresume.CoreExperience{
				Name:        nul.Something("ACME Corp"),
				Position:    activitypub.SomeString("Senior Developer"),
				StartDate:   nul.Something("2015-01-01"),
				EndDate:     nul.Something("2024-12-31"),
				Summary:     nul.Something("Led the backend team."),
				Description: nul.Something("Full-stack development."),
				Location:    nul.Something("Vancouver, BC"),
				Highlights:  activitypub.SomeStrings("Shipped v2.0", "Mentored juniors"),
				URL:         []activitypub.ProtoLink{activitypub.HRef("https://acme.example.com")},
			},
		},
	}

	cv.Volunteer = []jsonresume.ProtoExperience{
		jsonresume.Experience{
			CoreExperience: jsonresume.CoreExperience{
				Organization: nul.Something("Code for Good"),
				Position:     activitypub.SomeString("Volunteer Developer"),
				StartDate:    nul.Something("2020-03-01"),
				Summary:      nul.Something("Built internal tools."),
			},
		},
	}

	cv.Skills = []jsonresume.ProtoSkill{
		jsonresume.Skill{
			CoreSkill: jsonresume.CoreSkill{
				Name:     nul.Something("Backend Development"),
				Level:    nul.Something("Senior"),
				Keywords: activitypub.SomeStrings("Go", "PostgreSQL", "gRPC"),
			},
		},
	}

	cv.Languages = []jsonresume.ProtoLanguage{
		jsonresume.Language{
			CoreLanguage: jsonresume.CoreLanguage{
				Language: nul.Something("English"),
				Fluency:  nul.Something("Native"),
			},
		},
	}

	cv.Interests = []jsonresume.ProtoInterest{
		jsonresume.Interest{
			CoreInterest: jsonresume.CoreInterest{
				Name:     nul.Something("Open Source"),
				Keywords: activitypub.SomeStrings("ActivityPub", "Fediverse"),
			},
		},
	}

	cv.Projects = []jsonresume.ProtoProject{
		jsonresume.Project{
			CoreProject: jsonresume.CoreProject{
				Name:        nul.Something("go-jsonresume"),
				Description: nul.Something("JSON Resume as Go types with JSON-LD support."),
				Entity:      nul.Something("Personal"),
				StartDate:   nul.Something("2024-01-01"),
				EndDate:     nul.Something("2024-12-31"),
				Roles:       activitypub.SomeStrings("Lead", "Developer"),
				Keywords:    activitypub.SomeStrings("Go", "JSON-LD"),
				Highlights:  activitypub.SomeStrings("Full JSON Resume coverage"),
				ProjectType: nul.Something("application"),
				URL:         []activitypub.ProtoLink{activitypub.HRef("https://github.com/reiver/go-jsonresume")},
			},
		},
	}

	cv.Publications = []jsonresume.ProtoPublication{
		jsonresume.Publication{
			CorePublication: jsonresume.CorePublication{
				Name:        nul.Something("JSON-LD for Resumes"),
				Publisher:   nul.Something("Tech Blog"),
				ReleaseDate: nul.Something("2024-06-15"),
				Summary:     nul.Something("How to use JSON-LD with JSON Resume."),
				URL:         []activitypub.ProtoLink{activitypub.HRef("https://techblog.example.com/article")},
			},
		},
	}

	cv.References = []jsonresume.ProtoReference{
		jsonresume.Reference{
			CoreReference: jsonresume.CoreReference{
				Name:      nul.Something("Jane Doe"),
				Reference: nul.Something("Joe is an exceptional developer."),
			},
		},
	}

	cv.Meta = jsonresume.Meta{
		CoreMeta: jsonresume.CoreMeta{
			Canonical:    nul.Something("https://example.com/resume.json"),
			Version:      nul.Something("v1.0.0"),
			LastModified: nul.Something("2024-12-25T00:00:00"),
		},
	}

	// Marshal
	bytes, err := jsonld.Marshal(cv)
	if nil != err {
		t.Fatalf("Marshal failed: %s", err)
	}

	// Unmarshal into AnyResume
	var result jsonresume.AnyResume
	err = jsonld.Unmarshal(bytes, &result)
	if nil != err {
		t.Fatalf("Unmarshal failed: %s\nJSON:\n%s", err, string(bytes))
	}

	// Verify ID
	{
		expected := "http://example.com/resume/1"
		actual := result.ID.GetElse("")
		if expected != actual {
			t.Errorf("ID: expected %q, got %q", expected, actual)
		}
	}

	// Verify Schema
	assertNullableString(t, "Schema", nul.Something("https://jsonresume.org/schema"), result.Schema)

	// Verify Basics
	{
		basics, ok := result.Basics.(jsonresume.AnyBasics)
		if !ok {
			t.Fatalf("Basics: expected AnyBasics, got %T", result.Basics)
		}
		assertNullableString(t, "Basics.Name", nul.Something("Joe Blow"), basics.Name)
		assertNullableString(t, "Basics.Summary", nul.Something("An experienced developer."), basics.Summary)

		if len(basics.Image) != 1 {
			t.Errorf("Basics.Image: expected 1, got %d", len(basics.Image))
		}
		if len(basics.URL) != 1 {
			t.Errorf("Basics.URL: expected 1, got %d", len(basics.URL))
		}

		if len(basics.Location) != 1 {
			t.Errorf("Basics.Location: expected 1, got %d", len(basics.Location))
		} else {
			loc, ok := basics.Location[0].(jsonresume.AnyLocation)
			if !ok {
				t.Errorf("Basics.Location[0]: expected AnyLocation, got %T", basics.Location[0])
			} else {
				assertNullableString(t, "Location.Address", nul.Something("123 Main St"), loc.Address)
				assertNullableString(t, "Location.City", nul.Something("Springfield"), loc.City)
				assertNullableString(t, "Location.CountryCode", nul.Something("US"), loc.CountryCode)
				assertNullableString(t, "Location.PostalCode", nul.Something("62704"), loc.PostalCode)
				assertNullableString(t, "Location.Region", nul.Something("Illinois"), loc.Region)
			}
		}

		if len(basics.Profiles) != 1 {
			t.Errorf("Basics.Profiles: expected 1, got %d", len(basics.Profiles))
		} else {
			profile, ok := basics.Profiles[0].(jsonresume.AnyProfile)
			if !ok {
				t.Errorf("Basics.Profiles[0]: expected AnyProfile, got %T", basics.Profiles[0])
			} else {
				assertNullableString(t, "Profile.Network", nul.Something("Mastodon"), profile.Network)
				assertNullableString(t, "Profile.UserName", nul.Something("joeblow"), profile.UserName)
				if len(profile.URL) != 1 {
					t.Errorf("Profile.URL: expected 1, got %d", len(profile.URL))
				}
			}
		}
	}

	// Verify Awards
	assertSliceLen(t, "Awards", 1, len(result.Awards))
	if len(result.Awards) == 1 {
		award, ok := result.Awards[0].(jsonresume.AnyAward)
		if !ok {
			t.Errorf("Awards[0]: expected AnyAward, got %T", result.Awards[0])
		} else {
			assertNullableString(t, "Award.Title", nul.Something("Best Employee"), award.Title)
			assertNullableString(t, "Award.Date", nul.Something("2024-01-15"), award.Date)
			assertNullableString(t, "Award.Awarder", nul.Something("ACME Corp"), award.Awarder)
		}
	}

	// Verify Certificates
	assertSliceLen(t, "Certificates", 1, len(result.Certificates))
	if len(result.Certificates) == 1 {
		cert, ok := result.Certificates[0].(jsonresume.AnyCertificate)
		if !ok {
			t.Errorf("Certificates[0]: expected AnyCertificate, got %T", result.Certificates[0])
		} else {
			assertNullableString(t, "Certificate.Name", nul.Something("AWS Solutions Architect"), cert.Name)
			assertNullableString(t, "Certificate.Issuer", nul.Something("Amazon"), cert.Issuer)
			if len(cert.URL) != 1 {
				t.Errorf("Certificate.URL: expected 1, got %d", len(cert.URL))
			}
		}
	}

	// Verify Education
	assertSliceLen(t, "Education", 1, len(result.Education))
	if len(result.Education) == 1 {
		edu, ok := result.Education[0].(jsonresume.AnyEducation)
		if !ok {
			t.Errorf("Education[0]: expected AnyEducation, got %T", result.Education[0])
		} else {
			assertNullableString(t, "Education.Institution", nul.Something("MIT"), edu.Institution)
			assertNullableString(t, "Education.StartDate", nul.Something("2010-09-01"), edu.StartDate)
			assertNullableString(t, "Education.Score", nul.Something("3.9"), edu.Score)
			if len(edu.URL) != 1 {
				t.Errorf("Education.URL: expected 1, got %d", len(edu.URL))
			}
		}
	}

	// Verify Work
	assertSliceLen(t, "Work", 1, len(result.Work))
	if len(result.Work) == 1 {
		work, ok := result.Work[0].(jsonresume.AnyExperience)
		if !ok {
			t.Errorf("Work[0]: expected AnyExperience, got %T", result.Work[0])
		} else {
			assertNullableString(t, "Work.Name", nul.Something("ACME Corp"), work.Name)
			assertNullableString(t, "Work.StartDate", nul.Something("2015-01-01"), work.StartDate)
			assertNullableString(t, "Work.Summary", nul.Something("Led the backend team."), work.Summary)
			assertNullableString(t, "Work.Location", nul.Something("Vancouver, BC"), work.Location)
			if len(work.URL) != 1 {
				t.Errorf("Work.URL: expected 1, got %d", len(work.URL))
			}
		}
	}

	// Verify Volunteer
	assertSliceLen(t, "Volunteer", 1, len(result.Volunteer))
	if len(result.Volunteer) == 1 {
		vol, ok := result.Volunteer[0].(jsonresume.AnyExperience)
		if !ok {
			t.Errorf("Volunteer[0]: expected AnyExperience, got %T", result.Volunteer[0])
		} else {
			assertNullableString(t, "Volunteer.Organization", nul.Something("Code for Good"), vol.Organization)
			assertNullableString(t, "Volunteer.Summary", nul.Something("Built internal tools."), vol.Summary)
		}
	}

	// Verify Skills
	assertSliceLen(t, "Skills", 1, len(result.Skills))
	if len(result.Skills) == 1 {
		skill, ok := result.Skills[0].(jsonresume.AnySkill)
		if !ok {
			t.Errorf("Skills[0]: expected AnySkill, got %T", result.Skills[0])
		} else {
			assertNullableString(t, "Skill.Name", nul.Something("Backend Development"), skill.Name)
			assertNullableString(t, "Skill.Level", nul.Something("Senior"), skill.Level)
		}
	}

	// Verify Languages
	assertSliceLen(t, "Languages", 1, len(result.Languages))
	if len(result.Languages) == 1 {
		lang, ok := result.Languages[0].(jsonresume.AnyLanguage)
		if !ok {
			t.Errorf("Languages[0]: expected AnyLanguage, got %T", result.Languages[0])
		} else {
			assertNullableString(t, "Language.Language", nul.Something("English"), lang.Language)
			assertNullableString(t, "Language.Fluency", nul.Something("Native"), lang.Fluency)
		}
	}

	// Verify Interests
	assertSliceLen(t, "Interests", 1, len(result.Interests))
	if len(result.Interests) == 1 {
		interest, ok := result.Interests[0].(jsonresume.AnyInterest)
		if !ok {
			t.Errorf("Interests[0]: expected AnyInterest, got %T", result.Interests[0])
		} else {
			assertNullableString(t, "Interest.Name", nul.Something("Open Source"), interest.Name)
		}
	}

	// Verify Projects
	assertSliceLen(t, "Projects", 1, len(result.Projects))
	if len(result.Projects) == 1 {
		proj, ok := result.Projects[0].(jsonresume.AnyProject)
		if !ok {
			t.Errorf("Projects[0]: expected AnyProject, got %T", result.Projects[0])
		} else {
			assertNullableString(t, "Project.Name", nul.Something("go-jsonresume"), proj.Name)
			assertNullableString(t, "Project.Description", nul.Something("JSON Resume as Go types with JSON-LD support."), proj.Description)
			assertNullableString(t, "Project.Entity", nul.Something("Personal"), proj.Entity)
			assertNullableString(t, "Project.ProjectType", nul.Something("application"), proj.ProjectType)
			if len(proj.URL) != 1 {
				t.Errorf("Project.URL: expected 1, got %d", len(proj.URL))
			}
		}
	}

	// Verify Publications
	assertSliceLen(t, "Publications", 1, len(result.Publications))
	if len(result.Publications) == 1 {
		pub, ok := result.Publications[0].(jsonresume.AnyPublication)
		if !ok {
			t.Errorf("Publications[0]: expected AnyPublication, got %T", result.Publications[0])
		} else {
			assertNullableString(t, "Publication.Name", nul.Something("JSON-LD for Resumes"), pub.Name)
			assertNullableString(t, "Publication.Publisher", nul.Something("Tech Blog"), pub.Publisher)
			assertNullableString(t, "Publication.ReleaseDate", nul.Something("2024-06-15"), pub.ReleaseDate)
			if len(pub.URL) != 1 {
				t.Errorf("Publication.URL: expected 1, got %d", len(pub.URL))
			}
		}
	}

	// Verify References
	assertSliceLen(t, "References", 1, len(result.References))
	if len(result.References) == 1 {
		ref, ok := result.References[0].(jsonresume.AnyReference)
		if !ok {
			t.Errorf("References[0]: expected AnyReference, got %T", result.References[0])
		} else {
			assertNullableString(t, "Reference.Name", nul.Something("Jane Doe"), ref.Name)
			assertNullableString(t, "Reference.Reference", nul.Something("Joe is an exceptional developer."), ref.Reference)
		}
	}

	// Verify Meta
	{
		meta, ok := result.Meta.(jsonresume.AnyMeta)
		if !ok {
			t.Fatalf("Meta: expected AnyMeta, got %T", result.Meta)
		}
		assertNullableString(t, "Meta.Canonical", nul.Something("https://example.com/resume.json"), meta.Canonical)
		assertNullableString(t, "Meta.Version", nul.Something("v1.0.0"), meta.Version)
		assertNullableString(t, "Meta.LastModified", nul.Something("2024-12-25T00:00:00"), meta.LastModified)
	}
}

func assertNullableString(t *testing.T, field string, expected, actual nul.Nullable[string]) {
	t.Helper()
	if expected != actual {
		t.Errorf("%s: expected %#v, got %#v", field, expected, actual)
	}
}

func assertSliceLen(t *testing.T, field string, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Errorf("%s: expected length %d, got %d", field, expected, actual)
	}
}
