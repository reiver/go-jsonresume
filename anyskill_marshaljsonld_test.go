package jsonresume_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestAnySkill_marshalJSONLD(t *testing.T) {

	const context string = `"@context":{` +
		`"as":"https://www.w3.org/ns/activitystreams"` +
		`,` +
		`"cv":"https://w3id.org/fep/6158"` +

		`,` +

		`"name":"as:name"` +

		`,` +

		`"id":"cv:id"` +
		`,` +
		`"keywords":"cv:keywords"` +
		`,` +
		`"level":"cv:level"` +
		`,` +
		`"type":"cv:type"` +
	`}`

	tests := []struct {
		Value    jsonresume.AnySkill
		Name     activitypub.Name
		Expected string
	}{
		// 0
		{
			Value: jsonresume.AnySkill{
				Type: jsonld.SomeType("Skill"),
				CoreSkill: jsonresume.CoreSkill{
					Level: nul.Null[string](),
				},
			},
			Name: activitypub.SomeName("Go"),
			Expected: `{` +
				context +
				`,` +
				`"name":"Go"` +
				`,` +
				`"type":"Skill"` +
				`,` +
				`"keywords":null` +
				`,` +
				`"level":null` +
				`}`,
		},

		// 1
		{
			Value: jsonresume.AnySkill{
				ID:   jsonld.SomeID("http://example.com/resume/skill/go"),
				Type: jsonld.SomeType("Skill"),
				CoreSkill: jsonresume.CoreSkill{
					Keywords: activitypub.SomeStrings("backend", "concurrency"),
					Level:    nul.Something("Advanced"),
				},
			},
			Name: activitypub.SomeName("Go"),
			Expected: `{` +
				context +
				`,` +
				`"name":"Go"` +
				`,` +
				`"id":"http://example.com/resume/skill/go"` +
				`,` +
				`"type":"Skill"` +
				`,` +
				`"keywords":["backend","concurrency"]` +
				`,` +
				`"level":"Advanced"` +
				`}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := jsonld.Marshal(test.Name, test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		actual := string(actualBytes)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
