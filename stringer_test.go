package jsonresume

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-nul"
)

func TestString_empty(t *testing.T) {

	tests := []struct {
		Name     string
		Actual   string
		Expected string
	}{
		// 0
		{
			Name:   "Award",
			Actual: Award{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"awarder\": \"cv:awarder\",\n" +
				"    \"date\": \"cv:date\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"title\": \"cv:title\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"summary\": \"as:summary\"\n" +
				"  },\n" +
				"  \"type\": \"Award\"\n" +
				"}",
		},

		// 1
		{
			Name:   "AnyAward",
			Actual: AnyAward{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"awarder\": \"cv:awarder\",\n" +
				"    \"date\": \"cv:date\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"title\": \"cv:title\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"summary\": \"as:summary\"\n" +
				"  }\n" +
				"}",
		},

		// 3
		{
			Name:   "Basics",
			Actual: Basics{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"email\": \"cv:email\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"label\": \"cv:label\",\n" +
				"    \"location\": \"cv:location\",\n" +
				"    \"phone\": \"cv:phone\",\n" +
				"    \"profiles\": \"cv:profiles\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"image\": \"as:image\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"type\": \"Basics\"\n" +
				"}",
		},

		// 4
		{
			Name:   "AnyBasics",
			Actual: AnyBasics{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"email\": \"cv:email\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"label\": \"cv:label\",\n" +
				"    \"location\": \"cv:location\",\n" +
				"    \"phone\": \"cv:phone\",\n" +
				"    \"profiles\": \"cv:profiles\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"image\": \"as:image\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  }\n" +
				"}",
		},

		// 5
		{
			Name:   "Certificate",
			Actual: Certificate{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"date\": \"cv:date\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"issuer\": \"cv:issuer\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"type\": \"Certificate\"\n" +
				"}",
		},

		// 6
		{
			Name:   "AnyCertificate",
			Actual: AnyCertificate{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"date\": \"cv:date\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"issuer\": \"cv:issuer\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  }\n" +
				"}",
		},

		// 7
		{
			Name:   "Education",
			Actual: Education{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"area\": \"cv:area\",\n" +
				"    \"courses\": \"cv:courses\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"institution\": \"cv:institution\",\n" +
				"    \"score\": \"cv:score\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"studyType\": \"cv:studyType\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"type\": \"Education\"\n" +
				"}",
		},

		// 8
		{
			Name:   "AnyEducation",
			Actual: AnyEducation{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"area\": \"cv:area\",\n" +
				"    \"courses\": \"cv:courses\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"institution\": \"cv:institution\",\n" +
				"    \"score\": \"cv:score\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"studyType\": \"cv:studyType\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  }\n" +
				"}",
		},

		// 9
		{
			Name:   "Experience",
			Actual: Experience{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"description\": \"cv:description\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"highlights\": \"cv:highlights\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"location\": \"cv:location\",\n" +
				"    \"organization\": \"cv:organization\",\n" +
				"    \"position\": \"cv:position\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"type\": \"Experience\"\n" +
				"}",
		},

		// 10
		{
			Name:   "AnyExperience",
			Actual: AnyExperience{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"description\": \"cv:description\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"highlights\": \"cv:highlights\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"location\": \"cv:location\",\n" +
				"    \"organization\": \"cv:organization\",\n" +
				"    \"position\": \"cv:position\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  }\n" +
				"}",
		},

		// 11
		{
			Name:   "Interest",
			Actual: Interest{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  },\n" +
				"  \"type\": \"Interest\"\n" +
				"}",
		},

		// 12
		{
			Name:   "AnyInterest",
			Actual: AnyInterest{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  }\n" +
				"}",
		},

		// 13
		{
			Name:   "Language",
			Actual: Language{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"fluency\": \"cv:fluency\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"language\": \"cv:language\",\n" +
				"    \"type\": \"cv:type\"\n" +
				"  },\n" +
				"  \"type\": \"Language\"\n" +
				"}",
		},

		// 14
		{
			Name:   "AnyLanguage",
			Actual: AnyLanguage{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"fluency\": \"cv:fluency\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"language\": \"cv:language\",\n" +
				"    \"type\": \"cv:type\"\n" +
				"  }\n" +
				"}",
		},

		// 15
		{
			Name:   "Location",
			Actual: Location{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"address\": \"cv:address\",\n" +
				"    \"city\": \"cv:city\",\n" +
				"    \"countryCode\": \"cv:countryCode\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"postalCode\": \"cv:postalCode\",\n" +
				"    \"region\": \"cv:region\",\n" +
				"    \"type\": \"cv:type\"\n" +
				"  },\n" +
				"  \"type\": \"Location\"\n" +
				"}",
		},

		// 16
		{
			Name:   "AnyLocation",
			Actual: AnyLocation{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"address\": \"cv:address\",\n" +
				"    \"city\": \"cv:city\",\n" +
				"    \"countryCode\": \"cv:countryCode\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"postalCode\": \"cv:postalCode\",\n" +
				"    \"region\": \"cv:region\",\n" +
				"    \"type\": \"cv:type\"\n" +
				"  }\n" +
				"}",
		},

		// 17
		{
			Name:   "Meta",
			Actual: Meta{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"canonical\": \"cv:canonical\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"lastModified\": \"cv:lastModified\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"version\": \"cv:version\"\n" +
				"  },\n" +
				"  \"type\": \"Meta\"\n" +
				"}",
		},

		// 18
		{
			Name:   "AnyMeta",
			Actual: AnyMeta{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"canonical\": \"cv:canonical\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"lastModified\": \"cv:lastModified\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"version\": \"cv:version\"\n" +
				"  }\n" +
				"}",
		},

		// 19
		{
			Name:   "Profile",
			Actual: Profile{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"network\": \"cv:network\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"username\": \"cv:username\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"type\": \"Profile\"\n" +
				"}",
		},

		// 20
		{
			Name:   "AnyProfile",
			Actual: AnyProfile{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"network\": \"cv:network\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"username\": \"cv:username\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  }\n" +
				"}",
		},

		// 21
		{
			Name:   "Project",
			Actual: Project{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"@type\": \"cv:@type\",\n" +
				"    \"description\": \"cv:description\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"entity\": \"cv:entity\",\n" +
				"    \"highlights\": \"cv:highlights\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"roles\": \"cv:roles\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"@type\": \"Project\"\n" +
				"}",
		},

		// 22
		{
			Name:   "AnyProject",
			Actual: AnyProject{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"@type\": \"cv:@type\",\n" +
				"    \"description\": \"cv:description\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"entity\": \"cv:entity\",\n" +
				"    \"highlights\": \"cv:highlights\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"roles\": \"cv:roles\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  }\n" +
				"}",
		},

		// 23
		{
			Name:   "Publication",
			Actual: Publication{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"publisher\": \"cv:publisher\",\n" +
				"    \"releaseDate\": \"cv:releaseDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"type\": \"Publication\"\n" +
				"}",
		},

		// 24
		{
			Name:   "AnyPublication",
			Actual: AnyPublication{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"publisher\": \"cv:publisher\",\n" +
				"    \"releaseDate\": \"cv:releaseDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  }\n" +
				"}",
		},

		// 25
		{
			Name:   "Reference",
			Actual: Reference{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"reference\": \"cv:reference\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  },\n" +
				"  \"type\": \"Reference\"\n" +
				"}",
		},

		// 26
		{
			Name:   "AnyReference",
			Actual: AnyReference{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"reference\": \"cv:reference\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  }\n" +
				"}",
		},

		// 27
		{
			Name:   "Resume",
			Actual: Resume{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"$schema\": \"cv:$schema\",\n" +
				"    \"awards\": \"cv:awards\",\n" +
				"    \"basics\": \"cv:basics\",\n" +
				"    \"certificates\": \"cv:certificates\",\n" +
				"    \"education\": \"cv:education\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"interests\": \"cv:interests\",\n" +
				"    \"languages\": \"cv:languages\",\n" +
				"    \"meta\": \"cv:meta\",\n" +
				"    \"projects\": \"cv:projects\",\n" +
				"    \"publications\": \"cv:publications\",\n" +
				"    \"references\": \"cv:references\",\n" +
				"    \"skills\": \"cv:skills\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"volunteer\": \"cv:volunteer\",\n" +
				"    \"work\": \"cv:work\"\n" +
				"  },\n" +
				"  \"type\": \"Resume\"\n" +
				"}",
		},

		// 28
		{
			Name:   "AnyResume",
			Actual: AnyResume{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"$schema\": \"cv:$schema\",\n" +
				"    \"awards\": \"cv:awards\",\n" +
				"    \"basics\": \"cv:basics\",\n" +
				"    \"certificates\": \"cv:certificates\",\n" +
				"    \"education\": \"cv:education\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"interests\": \"cv:interests\",\n" +
				"    \"languages\": \"cv:languages\",\n" +
				"    \"meta\": \"cv:meta\",\n" +
				"    \"projects\": \"cv:projects\",\n" +
				"    \"publications\": \"cv:publications\",\n" +
				"    \"references\": \"cv:references\",\n" +
				"    \"skills\": \"cv:skills\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"volunteer\": \"cv:volunteer\",\n" +
				"    \"work\": \"cv:work\"\n" +
				"  }\n" +
				"}",
		},

		// 29
		{
			Name:   "Skill",
			Actual: Skill{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"level\": \"cv:level\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  },\n" +
				"  \"type\": \"Skill\"\n" +
				"}",
		},

		// 30
		{
			Name:   "AnySkill",
			Actual: AnySkill{}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"level\": \"cv:level\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  }\n" +
				"}",
		},
	}

	for testNumber, test := range tests {
		if test.Expected != test.Actual {
			t.Errorf("For test #%d (%s), the actual value is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED:\n%s", test.Expected)
			t.Logf("ACTUAL:\n%s", test.Actual)
		}
	}
}

func TestString_full(t *testing.T) {

	tests := []struct {
		Name     string
		Actual   string
		Expected string
	}{
		// 0
		{
			Name: "Award",
			Actual: Award{
				ID: jsonld.SomeID("http://example.com/award/1"),
				CoreAward: CoreAward{
					Awarder: nul.Something("ACME Corp"),
					Date:    nul.Something("2025-01-15"),
					Summary: nul.Something("For outstanding work"),
					Title:   nul.Something("Best Employee"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"awarder\": \"cv:awarder\",\n" +
				"    \"date\": \"cv:date\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"title\": \"cv:title\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"summary\": \"as:summary\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/award/1\",\n" +
				"  \"type\": \"Award\",\n" +
				"  \"awarder\": \"ACME Corp\",\n" +
				"  \"date\": \"2025-01-15\",\n" +
				"  \"summary\": \"For outstanding work\",\n" +
				"  \"title\": \"Best Employee\"\n" +
				"}",
		},

		// 1
		{
			Name: "AnyAward",
			Actual: AnyAward{
				ID:   jsonld.SomeID("http://example.com/award/1"),
				Type: jsonld.SomeType("Award"),
				CoreAward: CoreAward{
					Awarder: nul.Something("ACME Corp"),
					Date:    nul.Something("2025-01-15"),
					Summary: nul.Something("For outstanding work"),
					Title:   nul.Something("Best Employee"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"awarder\": \"cv:awarder\",\n" +
				"    \"date\": \"cv:date\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"title\": \"cv:title\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"summary\": \"as:summary\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/award/1\",\n" +
				"  \"type\": \"Award\",\n" +
				"  \"awarder\": \"ACME Corp\",\n" +
				"  \"date\": \"2025-01-15\",\n" +
				"  \"summary\": \"For outstanding work\",\n" +
				"  \"title\": \"Best Employee\"\n" +
				"}",
		},

		// 2
		{
			Name: "Basics",
			Actual: Basics{
				ID: jsonld.SomeID("http://example.com/basics/1"),
				CoreBasics: CoreBasics{
					EMail:   activitypub.SomeString("john@example.com"),
					Image:   []activitypub.ProtoImageOrProtoLink{activitypub.HRef("https://example.com/photo.jpg")},
					Label:   activitypub.SomeString("Developer"),
					Name:    nul.Something("John Doe"),
					Phone:   activitypub.SomeString("+1-555-0100"),
					Summary: nul.Something("A software developer"),
					URL:     []activitypub.ProtoLink{activitypub.HRef("https://johndoe.com")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"email\": \"cv:email\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"label\": \"cv:label\",\n" +
				"    \"location\": \"cv:location\",\n" +
				"    \"phone\": \"cv:phone\",\n" +
				"    \"profiles\": \"cv:profiles\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"image\": \"as:image\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/basics/1\",\n" +
				"  \"type\": \"Basics\",\n" +
				"  \"email\": \"john@example.com\",\n" +
				"  \"image\": \"https://example.com/photo.jpg\",\n" +
				"  \"label\": \"Developer\",\n" +
				"  \"name\": \"John Doe\",\n" +
				"  \"phone\": \"+1-555-0100\",\n" +
				"  \"summary\": \"A software developer\",\n" +
				"  \"url\": \"https://johndoe.com\"\n" +
				"}",
		},

		// 3
		{
			Name: "AnyBasics",
			Actual: AnyBasics{
				ID:   jsonld.SomeID("http://example.com/basics/1"),
				Type: jsonld.SomeType("Basics"),
				CoreBasics: CoreBasics{
					EMail:   activitypub.SomeString("john@example.com"),
					Image:   []activitypub.ProtoImageOrProtoLink{activitypub.HRef("https://example.com/photo.jpg")},
					Label:   activitypub.SomeString("Developer"),
					Name:    nul.Something("John Doe"),
					Phone:   activitypub.SomeString("+1-555-0100"),
					Summary: nul.Something("A software developer"),
					URL:     []activitypub.ProtoLink{activitypub.HRef("https://johndoe.com")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"email\": \"cv:email\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"label\": \"cv:label\",\n" +
				"    \"location\": \"cv:location\",\n" +
				"    \"phone\": \"cv:phone\",\n" +
				"    \"profiles\": \"cv:profiles\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"image\": \"as:image\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/basics/1\",\n" +
				"  \"type\": \"Basics\",\n" +
				"  \"email\": \"john@example.com\",\n" +
				"  \"image\": \"https://example.com/photo.jpg\",\n" +
				"  \"label\": \"Developer\",\n" +
				"  \"name\": \"John Doe\",\n" +
				"  \"phone\": \"+1-555-0100\",\n" +
				"  \"summary\": \"A software developer\",\n" +
				"  \"url\": \"https://johndoe.com\"\n" +
				"}",
		},

		// 4
		{
			Name: "Certificate",
			Actual: Certificate{
				ID: jsonld.SomeID("http://example.com/cert/1"),
				CoreCertificate: CoreCertificate{
					Date:   nul.Something("2024-06-01"),
					Name:   nul.Something("AWS Solutions Architect"),
					Issuer: nul.Something("Amazon"),
					URL:    []activitypub.ProtoLink{activitypub.HRef("https://example.com/cert")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"date\": \"cv:date\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"issuer\": \"cv:issuer\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/cert/1\",\n" +
				"  \"type\": \"Certificate\",\n" +
				"  \"date\": \"2024-06-01\",\n" +
				"  \"name\": \"AWS Solutions Architect\",\n" +
				"  \"issuer\": \"Amazon\",\n" +
				"  \"url\": \"https://example.com/cert\"\n" +
				"}",
		},

		// 5
		{
			Name: "AnyCertificate",
			Actual: AnyCertificate{
				ID:   jsonld.SomeID("http://example.com/cert/1"),
				Type: jsonld.SomeType("Certificate"),
				CoreCertificate: CoreCertificate{
					Date:   nul.Something("2024-06-01"),
					Name:   nul.Something("AWS Solutions Architect"),
					Issuer: nul.Something("Amazon"),
					URL:    []activitypub.ProtoLink{activitypub.HRef("https://example.com/cert")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"date\": \"cv:date\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"issuer\": \"cv:issuer\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/cert/1\",\n" +
				"  \"type\": \"Certificate\",\n" +
				"  \"date\": \"2024-06-01\",\n" +
				"  \"name\": \"AWS Solutions Architect\",\n" +
				"  \"issuer\": \"Amazon\",\n" +
				"  \"url\": \"https://example.com/cert\"\n" +
				"}",
		},

		// 6
		{
			Name: "Education",
			Actual: Education{
				ID: jsonld.SomeID("http://example.com/edu/1"),
				CoreEducation: CoreEducation{
					Area:        activitypub.SomeString("Computer Science"),
					Courses:     activitypub.SomeString("CS101"),
					EndDate:     nul.Something("2020-06-01"),
					Institution: nul.Something("MIT"),
					Score:       nul.Something("4.0"),
					StartDate:   nul.Something("2016-09-01"),
					StudyType:   activitypub.SomeString("Bachelor"),
					URL:         []activitypub.ProtoLink{activitypub.HRef("https://mit.edu")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"area\": \"cv:area\",\n" +
				"    \"courses\": \"cv:courses\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"institution\": \"cv:institution\",\n" +
				"    \"score\": \"cv:score\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"studyType\": \"cv:studyType\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/edu/1\",\n" +
				"  \"type\": \"Education\",\n" +
				"  \"area\": \"Computer Science\",\n" +
				"  \"courses\": \"CS101\",\n" +
				"  \"endDate\": \"2020-06-01\",\n" +
				"  \"institution\": \"MIT\",\n" +
				"  \"score\": \"4.0\",\n" +
				"  \"startDate\": \"2016-09-01\",\n" +
				"  \"studyType\": \"Bachelor\",\n" +
				"  \"url\": \"https://mit.edu\"\n" +
				"}",
		},

		// 7
		{
			Name: "AnyEducation",
			Actual: AnyEducation{
				ID:   jsonld.SomeID("http://example.com/edu/1"),
				Type: jsonld.SomeType("Education"),
				CoreEducation: CoreEducation{
					Area:        activitypub.SomeString("Computer Science"),
					Courses:     activitypub.SomeString("CS101"),
					EndDate:     nul.Something("2020-06-01"),
					Institution: nul.Something("MIT"),
					Score:       nul.Something("4.0"),
					StartDate:   nul.Something("2016-09-01"),
					StudyType:   activitypub.SomeString("Bachelor"),
					URL:         []activitypub.ProtoLink{activitypub.HRef("https://mit.edu")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"area\": \"cv:area\",\n" +
				"    \"courses\": \"cv:courses\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"institution\": \"cv:institution\",\n" +
				"    \"score\": \"cv:score\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"studyType\": \"cv:studyType\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/edu/1\",\n" +
				"  \"type\": \"Education\",\n" +
				"  \"area\": \"Computer Science\",\n" +
				"  \"courses\": \"CS101\",\n" +
				"  \"endDate\": \"2020-06-01\",\n" +
				"  \"institution\": \"MIT\",\n" +
				"  \"score\": \"4.0\",\n" +
				"  \"startDate\": \"2016-09-01\",\n" +
				"  \"studyType\": \"Bachelor\",\n" +
				"  \"url\": \"https://mit.edu\"\n" +
				"}",
		},

		// 8
		{
			Name: "Experience",
			Actual: Experience{
				ID: jsonld.SomeID("http://example.com/exp/1"),
				CoreExperience: CoreExperience{
					Description:  nul.Something("Full-stack development"),
					EndDate:      nul.Something("2025-01-01"),
					Highlights:   activitypub.SomeString("Led team of 5"),
					Location:     nul.Something("San Francisco"),
					Name:         nul.Something("ACME Corp"),
					Organization: nul.Something("ACME Corp"),
					Position:     activitypub.SomeString("Engineer"),
					StartDate:    nul.Something("2020-06-15"),
					Summary:      nul.Something("Worked on core platform"),
					URL:          []activitypub.ProtoLink{activitypub.HRef("https://acme.com")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"description\": \"cv:description\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"highlights\": \"cv:highlights\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"location\": \"cv:location\",\n" +
				"    \"organization\": \"cv:organization\",\n" +
				"    \"position\": \"cv:position\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/exp/1\",\n" +
				"  \"type\": \"Experience\",\n" +
				"  \"description\": \"Full-stack development\",\n" +
				"  \"endDate\": \"2025-01-01\",\n" +
				"  \"highlights\": \"Led team of 5\",\n" +
				"  \"location\": \"San Francisco\",\n" +
				"  \"name\": \"ACME Corp\",\n" +
				"  \"organization\": \"ACME Corp\",\n" +
				"  \"position\": \"Engineer\",\n" +
				"  \"startDate\": \"2020-06-15\",\n" +
				"  \"summary\": \"Worked on core platform\",\n" +
				"  \"url\": \"https://acme.com\"\n" +
				"}",
		},

		// 9
		{
			Name: "AnyExperience",
			Actual: AnyExperience{
				ID:   jsonld.SomeID("http://example.com/exp/1"),
				Type: jsonld.SomeType("Experience"),
				CoreExperience: CoreExperience{
					Description:  nul.Something("Full-stack development"),
					EndDate:      nul.Something("2025-01-01"),
					Highlights:   activitypub.SomeString("Led team of 5"),
					Location:     nul.Something("San Francisco"),
					Name:         nul.Something("ACME Corp"),
					Organization: nul.Something("ACME Corp"),
					Position:     activitypub.SomeString("Engineer"),
					StartDate:    nul.Something("2020-06-15"),
					Summary:      nul.Something("Worked on core platform"),
					URL:          []activitypub.ProtoLink{activitypub.HRef("https://acme.com")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"description\": \"cv:description\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"highlights\": \"cv:highlights\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"location\": \"cv:location\",\n" +
				"    \"organization\": \"cv:organization\",\n" +
				"    \"position\": \"cv:position\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/exp/1\",\n" +
				"  \"type\": \"Experience\",\n" +
				"  \"description\": \"Full-stack development\",\n" +
				"  \"endDate\": \"2025-01-01\",\n" +
				"  \"highlights\": \"Led team of 5\",\n" +
				"  \"location\": \"San Francisco\",\n" +
				"  \"name\": \"ACME Corp\",\n" +
				"  \"organization\": \"ACME Corp\",\n" +
				"  \"position\": \"Engineer\",\n" +
				"  \"startDate\": \"2020-06-15\",\n" +
				"  \"summary\": \"Worked on core platform\",\n" +
				"  \"url\": \"https://acme.com\"\n" +
				"}",
		},

		// 10
		{
			Name: "Interest",
			Actual: Interest{
				ID: jsonld.SomeID("http://example.com/interest/1"),
				CoreInterest: CoreInterest{
					Name:     nul.Something("Open Source"),
					Keywords: activitypub.SomeString("golang"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/interest/1\",\n" +
				"  \"type\": \"Interest\",\n" +
				"  \"name\": \"Open Source\",\n" +
				"  \"keywords\": \"golang\"\n" +
				"}",
		},

		// 11
		{
			Name: "AnyInterest",
			Actual: AnyInterest{
				ID:   jsonld.SomeID("http://example.com/interest/1"),
				Type: jsonld.SomeType("Interest"),
				CoreInterest: CoreInterest{
					Name:     nul.Something("Open Source"),
					Keywords: activitypub.SomeString("golang"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/interest/1\",\n" +
				"  \"type\": \"Interest\",\n" +
				"  \"name\": \"Open Source\",\n" +
				"  \"keywords\": \"golang\"\n" +
				"}",
		},

		// 12
		{
			Name: "Language",
			Actual: Language{
				ID: jsonld.SomeID("http://example.com/lang/1"),
				CoreLanguage: CoreLanguage{
					Fluency:  nul.Something("Native"),
					Language: nul.Something("English"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"fluency\": \"cv:fluency\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"language\": \"cv:language\",\n" +
				"    \"type\": \"cv:type\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/lang/1\",\n" +
				"  \"type\": \"Language\",\n" +
				"  \"fluency\": \"Native\",\n" +
				"  \"language\": \"English\"\n" +
				"}",
		},

		// 13
		{
			Name: "AnyLanguage",
			Actual: AnyLanguage{
				ID:   jsonld.SomeID("http://example.com/lang/1"),
				Type: jsonld.SomeType("Language"),
				CoreLanguage: CoreLanguage{
					Fluency:  nul.Something("Native"),
					Language: nul.Something("English"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"fluency\": \"cv:fluency\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"language\": \"cv:language\",\n" +
				"    \"type\": \"cv:type\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/lang/1\",\n" +
				"  \"type\": \"Language\",\n" +
				"  \"fluency\": \"Native\",\n" +
				"  \"language\": \"English\"\n" +
				"}",
		},

		// 14
		{
			Name: "Location",
			Actual: Location{
				ID: jsonld.SomeID("http://example.com/loc/1"),
				CoreLocation: CoreLocation{
					Address:     nul.Something("123 Main St"),
					City:        nul.Something("San Francisco"),
					CountryCode: nul.Something("US"),
					PostalCode:  nul.Something("94105"),
					Region:      nul.Something("California"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"address\": \"cv:address\",\n" +
				"    \"city\": \"cv:city\",\n" +
				"    \"countryCode\": \"cv:countryCode\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"postalCode\": \"cv:postalCode\",\n" +
				"    \"region\": \"cv:region\",\n" +
				"    \"type\": \"cv:type\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/loc/1\",\n" +
				"  \"type\": \"Location\",\n" +
				"  \"address\": \"123 Main St\",\n" +
				"  \"city\": \"San Francisco\",\n" +
				"  \"countryCode\": \"US\",\n" +
				"  \"postalCode\": \"94105\",\n" +
				"  \"region\": \"California\"\n" +
				"}",
		},

		// 15
		{
			Name: "AnyLocation",
			Actual: AnyLocation{
				ID:   jsonld.SomeID("http://example.com/loc/1"),
				Type: jsonld.SomeType("Location"),
				CoreLocation: CoreLocation{
					Address:     nul.Something("123 Main St"),
					City:        nul.Something("San Francisco"),
					CountryCode: nul.Something("US"),
					PostalCode:  nul.Something("94105"),
					Region:      nul.Something("California"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"address\": \"cv:address\",\n" +
				"    \"city\": \"cv:city\",\n" +
				"    \"countryCode\": \"cv:countryCode\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"postalCode\": \"cv:postalCode\",\n" +
				"    \"region\": \"cv:region\",\n" +
				"    \"type\": \"cv:type\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/loc/1\",\n" +
				"  \"type\": \"Location\",\n" +
				"  \"address\": \"123 Main St\",\n" +
				"  \"city\": \"San Francisco\",\n" +
				"  \"countryCode\": \"US\",\n" +
				"  \"postalCode\": \"94105\",\n" +
				"  \"region\": \"California\"\n" +
				"}",
		},

		// 16
		{
			Name: "Meta",
			Actual: Meta{
				ID: jsonld.SomeID("http://example.com/meta/1"),
				CoreMeta: CoreMeta{
					Canonical:    nul.Something("https://example.com/resume"),
					LastModified: nul.Something("2025-05-27"),
					Version:      nul.Something("v1.0.0"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"canonical\": \"cv:canonical\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"lastModified\": \"cv:lastModified\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"version\": \"cv:version\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/meta/1\",\n" +
				"  \"type\": \"Meta\",\n" +
				"  \"canonical\": \"https://example.com/resume\",\n" +
				"  \"lastModified\": \"2025-05-27\",\n" +
				"  \"version\": \"v1.0.0\"\n" +
				"}",
		},

		// 17
		{
			Name: "AnyMeta",
			Actual: AnyMeta{
				ID:   jsonld.SomeID("http://example.com/meta/1"),
				Type: jsonld.SomeType("Meta"),
				CoreMeta: CoreMeta{
					Canonical:    nul.Something("https://example.com/resume"),
					LastModified: nul.Something("2025-05-27"),
					Version:      nul.Something("v1.0.0"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"canonical\": \"cv:canonical\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"lastModified\": \"cv:lastModified\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"version\": \"cv:version\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/meta/1\",\n" +
				"  \"type\": \"Meta\",\n" +
				"  \"canonical\": \"https://example.com/resume\",\n" +
				"  \"lastModified\": \"2025-05-27\",\n" +
				"  \"version\": \"v1.0.0\"\n" +
				"}",
		},

		// 18
		{
			Name: "Profile",
			Actual: Profile{
				ID: jsonld.SomeID("http://example.com/profile/1"),
				CoreProfile: CoreProfile{
					Network:  nul.Something("GitHub"),
					UserName: nul.Something("johndoe"),
					URL:      []activitypub.ProtoLink{activitypub.HRef("https://github.com/johndoe")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"network\": \"cv:network\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"username\": \"cv:username\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/profile/1\",\n" +
				"  \"type\": \"Profile\",\n" +
				"  \"network\": \"GitHub\",\n" +
				"  \"username\": \"johndoe\",\n" +
				"  \"url\": \"https://github.com/johndoe\"\n" +
				"}",
		},

		// 19
		{
			Name: "AnyProfile",
			Actual: AnyProfile{
				ID:   jsonld.SomeID("http://example.com/profile/1"),
				Type: jsonld.SomeType("Profile"),
				CoreProfile: CoreProfile{
					Network:  nul.Something("GitHub"),
					UserName: nul.Something("johndoe"),
					URL:      []activitypub.ProtoLink{activitypub.HRef("https://github.com/johndoe")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"network\": \"cv:network\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"username\": \"cv:username\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/profile/1\",\n" +
				"  \"type\": \"Profile\",\n" +
				"  \"network\": \"GitHub\",\n" +
				"  \"username\": \"johndoe\",\n" +
				"  \"url\": \"https://github.com/johndoe\"\n" +
				"}",
		},

		// 20
		{
			Name: "Project",
			Actual: Project{
				ID: jsonld.SomeID("http://example.com/project/1"),
				CoreProject: CoreProject{
					Description: nul.Something("A web framework"),
					EndDate:     nul.Something("2024-12-31"),
					Entity:      nul.Something("ACME Corp"),
					Highlights:  activitypub.SomeString("1M downloads"),
					Keywords:    activitypub.SomeString("golang"),
					Name:        nul.Something("WebKit"),
					Roles:       activitypub.SomeString("Lead Developer"),
					StartDate:   nul.Something("2022-01-01"),
					ProjectType: nul.Something("application"),
					URL:         []activitypub.ProtoLink{activitypub.HRef("https://example.com/project")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"@type\": \"cv:@type\",\n" +
				"    \"description\": \"cv:description\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"entity\": \"cv:entity\",\n" +
				"    \"highlights\": \"cv:highlights\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"roles\": \"cv:roles\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/project/1\",\n" +
				"  \"@type\": \"Project\",\n" +
				"  \"description\": \"A web framework\",\n" +
				"  \"endDate\": \"2024-12-31\",\n" +
				"  \"entity\": \"ACME Corp\",\n" +
				"  \"highlights\": \"1M downloads\",\n" +
				"  \"keywords\": \"golang\",\n" +
				"  \"name\": \"WebKit\",\n" +
				"  \"roles\": \"Lead Developer\",\n" +
				"  \"startDate\": \"2022-01-01\",\n" +
				"  \"type\": \"application\",\n" +
				"  \"url\": \"https://example.com/project\"\n" +
				"}",
		},

		// 21
		{
			Name: "AnyProject",
			Actual: AnyProject{
				ID:   jsonld.SomeID("http://example.com/project/1"),
				Type: jsonld.SomeType("Project"),
				CoreProject: CoreProject{
					Description: nul.Something("A web framework"),
					EndDate:     nul.Something("2024-12-31"),
					Entity:      nul.Something("ACME Corp"),
					Highlights:  activitypub.SomeString("1M downloads"),
					Keywords:    activitypub.SomeString("golang"),
					Name:        nul.Something("WebKit"),
					Roles:       activitypub.SomeString("Lead Developer"),
					StartDate:   nul.Something("2022-01-01"),
					ProjectType: nul.Something("application"),
					URL:         []activitypub.ProtoLink{activitypub.HRef("https://example.com/project")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"@type\": \"cv:@type\",\n" +
				"    \"description\": \"cv:description\",\n" +
				"    \"endDate\": \"cv:endDate\",\n" +
				"    \"entity\": \"cv:entity\",\n" +
				"    \"highlights\": \"cv:highlights\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"roles\": \"cv:roles\",\n" +
				"    \"startDate\": \"cv:startDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/project/1\",\n" +
				"  \"@type\": \"Project\",\n" +
				"  \"description\": \"A web framework\",\n" +
				"  \"endDate\": \"2024-12-31\",\n" +
				"  \"entity\": \"ACME Corp\",\n" +
				"  \"highlights\": \"1M downloads\",\n" +
				"  \"keywords\": \"golang\",\n" +
				"  \"name\": \"WebKit\",\n" +
				"  \"roles\": \"Lead Developer\",\n" +
				"  \"startDate\": \"2022-01-01\",\n" +
				"  \"type\": \"application\",\n" +
				"  \"url\": \"https://example.com/project\"\n" +
				"}",
		},

		// 22
		{
			Name: "Publication",
			Actual: Publication{
				ID: jsonld.SomeID("http://example.com/pub/1"),
				CorePublication: CorePublication{
					Name:        nul.Something("My Paper"),
					Publisher:   nul.Something("IEEE"),
					ReleaseDate: nul.Something("2023-03-15"),
					Summary:     nul.Something("A research paper"),
					URL:         []activitypub.ProtoLink{activitypub.HRef("https://example.com/paper")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"publisher\": \"cv:publisher\",\n" +
				"    \"releaseDate\": \"cv:releaseDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/pub/1\",\n" +
				"  \"type\": \"Publication\",\n" +
				"  \"name\": \"My Paper\",\n" +
				"  \"publisher\": \"IEEE\",\n" +
				"  \"releaseDate\": \"2023-03-15\",\n" +
				"  \"summary\": \"A research paper\",\n" +
				"  \"url\": \"https://example.com/paper\"\n" +
				"}",
		},

		// 23
		{
			Name: "AnyPublication",
			Actual: AnyPublication{
				ID:   jsonld.SomeID("http://example.com/pub/1"),
				Type: jsonld.SomeType("Publication"),
				CorePublication: CorePublication{
					Name:        nul.Something("My Paper"),
					Publisher:   nul.Something("IEEE"),
					ReleaseDate: nul.Something("2023-03-15"),
					Summary:     nul.Something("A research paper"),
					URL:         []activitypub.ProtoLink{activitypub.HRef("https://example.com/paper")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"publisher\": \"cv:publisher\",\n" +
				"    \"releaseDate\": \"cv:releaseDate\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\",\n" +
				"    \"summary\": \"as:summary\",\n" +
				"    \"url\": \"as:url\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/pub/1\",\n" +
				"  \"type\": \"Publication\",\n" +
				"  \"name\": \"My Paper\",\n" +
				"  \"publisher\": \"IEEE\",\n" +
				"  \"releaseDate\": \"2023-03-15\",\n" +
				"  \"summary\": \"A research paper\",\n" +
				"  \"url\": \"https://example.com/paper\"\n" +
				"}",
		},

		// 24
		{
			Name: "Reference",
			Actual: Reference{
				ID: jsonld.SomeID("http://example.com/ref/1"),
				CoreReference: CoreReference{
					Name:      nul.Something("Jane Smith"),
					Reference: nul.Something("Excellent colleague"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"reference\": \"cv:reference\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/ref/1\",\n" +
				"  \"type\": \"Reference\",\n" +
				"  \"name\": \"Jane Smith\",\n" +
				"  \"reference\": \"Excellent colleague\"\n" +
				"}",
		},

		// 25
		{
			Name: "AnyReference",
			Actual: AnyReference{
				ID:   jsonld.SomeID("http://example.com/ref/1"),
				Type: jsonld.SomeType("Reference"),
				CoreReference: CoreReference{
					Name:      nul.Something("Jane Smith"),
					Reference: nul.Something("Excellent colleague"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"reference\": \"cv:reference\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/ref/1\",\n" +
				"  \"type\": \"Reference\",\n" +
				"  \"name\": \"Jane Smith\",\n" +
				"  \"reference\": \"Excellent colleague\"\n" +
				"}",
		},

		// 26
		{
			Name: "Resume",
			Actual: Resume{
				ID: jsonld.SomeID("http://example.com/resume/1"),
				CoreResume: CoreResume{
					Schema:       nul.Something("https://raw.githubusercontent.com/jsonresume/resume-schema/v1.0.0/schema.json"),
					Awards:       []ProtoAward{SomeAwardID("http://example.com/award/1")},
					Basics:       SomeBasicsID("http://example.com/basics/1"),
					Certificates: []ProtoCertificate{SomeCertificateID("http://example.com/cert/1")},
					Education:    []ProtoEducation{SomeEducationID("http://example.com/edu/1")},
					Interests:    []ProtoInterest{SomeInterestID("http://example.com/interest/1")},
					Languages:    []ProtoLanguage{SomeLanguageID("http://example.com/lang/1")},
					Meta:         SomeMetaID("http://example.com/meta/1"),
					Projects:     []ProtoProject{SomeProjectID("http://example.com/project/1")},
					Publications: []ProtoPublication{SomePublicationID("http://example.com/pub/1")},
					References:   []ProtoReference{SomeReferenceID("http://example.com/ref/1")},
					Skills:       []ProtoSkill{SomeSkillID("http://example.com/skill/1")},
					Volunteer:    []ProtoExperience{SomeExperienceID("http://example.com/vol/1")},
					Work:         []ProtoExperience{SomeExperienceID("http://example.com/work/1")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"$schema\": \"cv:$schema\",\n" +
				"    \"awards\": \"cv:awards\",\n" +
				"    \"basics\": \"cv:basics\",\n" +
				"    \"certificates\": \"cv:certificates\",\n" +
				"    \"education\": \"cv:education\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"interests\": \"cv:interests\",\n" +
				"    \"languages\": \"cv:languages\",\n" +
				"    \"meta\": \"cv:meta\",\n" +
				"    \"projects\": \"cv:projects\",\n" +
				"    \"publications\": \"cv:publications\",\n" +
				"    \"references\": \"cv:references\",\n" +
				"    \"skills\": \"cv:skills\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"volunteer\": \"cv:volunteer\",\n" +
				"    \"work\": \"cv:work\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/resume/1\",\n" +
				"  \"type\": \"Resume\",\n" +
				"  \"$schema\": \"https://raw.githubusercontent.com/jsonresume/resume-schema/v1.0.0/schema.json\",\n" +
				"  \"awards\": [\n" +
				"    \"http://example.com/award/1\"\n" +
				"  ],\n" +
				"  \"basics\": \"http://example.com/basics/1\",\n" +
				"  \"certificates\": [\n" +
				"    \"http://example.com/cert/1\"\n" +
				"  ],\n" +
				"  \"education\": [\n" +
				"    \"http://example.com/edu/1\"\n" +
				"  ],\n" +
				"  \"interests\": [\n" +
				"    \"http://example.com/interest/1\"\n" +
				"  ],\n" +
				"  \"languages\": [\n" +
				"    \"http://example.com/lang/1\"\n" +
				"  ],\n" +
				"  \"meta\": \"http://example.com/meta/1\",\n" +
				"  \"projects\": [\n" +
				"    \"http://example.com/project/1\"\n" +
				"  ],\n" +
				"  \"publications\": [\n" +
				"    \"http://example.com/pub/1\"\n" +
				"  ],\n" +
				"  \"references\": [\n" +
				"    \"http://example.com/ref/1\"\n" +
				"  ],\n" +
				"  \"skills\": [\n" +
				"    \"http://example.com/skill/1\"\n" +
				"  ],\n" +
				"  \"volunteer\": [\n" +
				"    \"http://example.com/vol/1\"\n" +
				"  ],\n" +
				"  \"work\": [\n" +
				"    \"http://example.com/work/1\"\n" +
				"  ]\n" +
				"}",
		},

		// 27
		{
			Name: "AnyResume",
			Actual: AnyResume{
				ID:   jsonld.SomeID("http://example.com/resume/1"),
				Type: jsonld.SomeType("Resume"),
				CoreResume: CoreResume{
					Schema:       nul.Something("https://raw.githubusercontent.com/jsonresume/resume-schema/v1.0.0/schema.json"),
					Awards:       []ProtoAward{SomeAwardID("http://example.com/award/1")},
					Basics:       SomeBasicsID("http://example.com/basics/1"),
					Certificates: []ProtoCertificate{SomeCertificateID("http://example.com/cert/1")},
					Education:    []ProtoEducation{SomeEducationID("http://example.com/edu/1")},
					Interests:    []ProtoInterest{SomeInterestID("http://example.com/interest/1")},
					Languages:    []ProtoLanguage{SomeLanguageID("http://example.com/lang/1")},
					Meta:         SomeMetaID("http://example.com/meta/1"),
					Projects:     []ProtoProject{SomeProjectID("http://example.com/project/1")},
					Publications: []ProtoPublication{SomePublicationID("http://example.com/pub/1")},
					References:   []ProtoReference{SomeReferenceID("http://example.com/ref/1")},
					Skills:       []ProtoSkill{SomeSkillID("http://example.com/skill/1")},
					Volunteer:    []ProtoExperience{SomeExperienceID("http://example.com/vol/1")},
					Work:         []ProtoExperience{SomeExperienceID("http://example.com/work/1")},
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"$schema\": \"cv:$schema\",\n" +
				"    \"awards\": \"cv:awards\",\n" +
				"    \"basics\": \"cv:basics\",\n" +
				"    \"certificates\": \"cv:certificates\",\n" +
				"    \"education\": \"cv:education\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"interests\": \"cv:interests\",\n" +
				"    \"languages\": \"cv:languages\",\n" +
				"    \"meta\": \"cv:meta\",\n" +
				"    \"projects\": \"cv:projects\",\n" +
				"    \"publications\": \"cv:publications\",\n" +
				"    \"references\": \"cv:references\",\n" +
				"    \"skills\": \"cv:skills\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"volunteer\": \"cv:volunteer\",\n" +
				"    \"work\": \"cv:work\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/resume/1\",\n" +
				"  \"type\": \"Resume\",\n" +
				"  \"$schema\": \"https://raw.githubusercontent.com/jsonresume/resume-schema/v1.0.0/schema.json\",\n" +
				"  \"awards\": [\n" +
				"    \"http://example.com/award/1\"\n" +
				"  ],\n" +
				"  \"basics\": \"http://example.com/basics/1\",\n" +
				"  \"certificates\": [\n" +
				"    \"http://example.com/cert/1\"\n" +
				"  ],\n" +
				"  \"education\": [\n" +
				"    \"http://example.com/edu/1\"\n" +
				"  ],\n" +
				"  \"interests\": [\n" +
				"    \"http://example.com/interest/1\"\n" +
				"  ],\n" +
				"  \"languages\": [\n" +
				"    \"http://example.com/lang/1\"\n" +
				"  ],\n" +
				"  \"meta\": \"http://example.com/meta/1\",\n" +
				"  \"projects\": [\n" +
				"    \"http://example.com/project/1\"\n" +
				"  ],\n" +
				"  \"publications\": [\n" +
				"    \"http://example.com/pub/1\"\n" +
				"  ],\n" +
				"  \"references\": [\n" +
				"    \"http://example.com/ref/1\"\n" +
				"  ],\n" +
				"  \"skills\": [\n" +
				"    \"http://example.com/skill/1\"\n" +
				"  ],\n" +
				"  \"volunteer\": [\n" +
				"    \"http://example.com/vol/1\"\n" +
				"  ],\n" +
				"  \"work\": [\n" +
				"    \"http://example.com/work/1\"\n" +
				"  ]\n" +
				"}",
		},

		// 28
		{
			Name: "Skill",
			Actual: Skill{
				ID: jsonld.SomeID("http://example.com/skill/1"),
				CoreSkill: CoreSkill{
					Keywords: activitypub.SomeString("REST"),
					Level:    nul.Something("Advanced"),
					Name:     nul.Something("Go"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"level\": \"cv:level\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/skill/1\",\n" +
				"  \"type\": \"Skill\",\n" +
				"  \"keywords\": \"REST\",\n" +
				"  \"level\": \"Advanced\",\n" +
				"  \"name\": \"Go\"\n" +
				"}",
		},

		// 29
		{
			Name: "AnySkill",
			Actual: AnySkill{
				ID:   jsonld.SomeID("http://example.com/skill/1"),
				Type: jsonld.SomeType("Skill"),
				CoreSkill: CoreSkill{
					Keywords: activitypub.SomeString("REST"),
					Level:    nul.Something("Advanced"),
					Name:     nul.Something("Go"),
				},
			}.String(),
			Expected: "{\n" +
				"  \"@context\": {\n" +
				"    \"cv\": \"https://w3id.org/fep/6158\",\n" +
				"    \"as\": \"http://www.w3.org/ns/activitystreams\",\n" +
				"    \"id\": \"cv:id\",\n" +
				"    \"keywords\": \"cv:keywords\",\n" +
				"    \"level\": \"cv:level\",\n" +
				"    \"type\": \"cv:type\",\n" +
				"    \"name\": \"as:name\"\n" +
				"  },\n" +
				"  \"id\": \"http://example.com/skill/1\",\n" +
				"  \"type\": \"Skill\",\n" +
				"  \"keywords\": \"REST\",\n" +
				"  \"level\": \"Advanced\",\n" +
				"  \"name\": \"Go\"\n" +
				"}",
		},
	}

	for testNumber, test := range tests {
		if test.Expected != test.Actual {
			t.Errorf("For test #%d (%s), the actual value is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED:\n%s", test.Expected)
			t.Logf("ACTUAL:\n%s", test.Actual)
		}
	}
}

// TestStringJSON_marshalFailure verifies that stringJSON gracefully returns "{}"
// when given a value that jsonld.Marshal cannot serialize, rather than panicking
// or returning corrupt output. jsonld.Marshal rejects non-struct/non-map values
// that don't implement json.Marshaler.
func TestStringJSON_marshalFailure(t *testing.T) {

	tests := []struct {
		Name  string
		Value any
	}{
		{Name: "channel",  Value: make(chan int)},
		{Name: "function", Value: func() {}},
		{Name: "integer",  Value: 42},
		{Name: "string",   Value: "hello"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := stringJSON(test.Value)
			if result != "{}" {
				t.Errorf("Expected \"{}\" for unmarshalable %s but got: %s", test.Name, result)
			}
		})
	}
}
