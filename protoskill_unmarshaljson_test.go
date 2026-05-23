package jsonresume_test

import (
	"errors"
	"reflect"
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonresume"
	"github.com/reiver/go-nul"
)

func TestProtoSkillUnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON          []byte
		ExpectNil     bool
		ExpectError   bool
		ExpectedError error
		ExpectSkillID bool
		ExpectSkill   bool
		Expected      jsonresume.ProtoSkill
	}{
		// 0: empty bytes
		{
			JSON:          []byte{},
			ExpectError:   true,
			ExpectedError: jsonresume.ErrBytesEmpty,
		},

		// 1: null
		{
			JSON:      []byte("null"),
			ExpectNil: true,
		},

		// 2: JSON string (IRI) → SkillID
		{
			JSON:          []byte(`"http://example.com/resume/skill/go"`),
			ExpectSkillID: true,
			Expected:      jsonresume.SomeSkillID("http://example.com/resume/skill/go"),
		},

		// 3: JSON object → Skill with level
		{
			JSON:        []byte(`{"type":"Skill","level":"advanced"}`),
			ExpectSkill: true,
			Expected: jsonresume.Skill{
				CoreSkill: jsonresume.CoreSkill{
					Level: nul.Something("advanced"),
				},
			},
		},

		// 4: JSON object → Skill with keywords
		{
			JSON:        []byte(`{"type":"Skill","keywords":["goroutines","channels","testing"]}`),
			ExpectSkill: true,
			Expected: jsonresume.Skill{
				CoreSkill: jsonresume.CoreSkill{
					Keywords: activitypub.SomeStrings("goroutines", "channels", "testing"),
				},
			},
		},

		// 5: JSON object → Skill with level and keywords
		{
			JSON:        []byte(`{"type":"Skill","level":"expert","keywords":["REST","GraphQL"]}`),
			ExpectSkill: true,
			Expected: jsonresume.Skill{
				CoreSkill: jsonresume.CoreSkill{
					Level:    nul.Something("expert"),
					Keywords: activitypub.SomeStrings("REST", "GraphQL"),
				},
			},
		},

		// 6: unsupported type (array)
		{
			JSON:          []byte(`[1,2,3]`),
			ExpectError:   true,
			ExpectedError: jsonresume.ErrTypeUnsupported,
		},

		// 7: unsupported type (number)
		{
			JSON:          []byte(`42`),
			ExpectError:   true,
			ExpectedError: jsonresume.ErrTypeUnsupported,
		},

		// 8: another JSON string (IRI)
		{
			JSON:          []byte(`"http://example.com/resume/skill/rust"`),
			ExpectSkillID: true,
			Expected:      jsonresume.SomeSkillID("http://example.com/resume/skill/rust"),
		},

		// 9: minimal JSON object
		{
			JSON:        []byte(`{}`),
			ExpectSkill: true,
			Expected:    jsonresume.Skill{},
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonresume.ProtoSkillUnmarshalJSON(test.JSON)

		if test.ExpectError {
			if nil == err {
				t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}

			if nil != test.ExpectedError {
				if !errors.Is(err, test.ExpectedError) {
					t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
					t.Logf("EXPECTED-ERROR: %s", test.ExpectedError)
					t.Logf("ACTUAL-ERROR:   %s", err)
					t.Logf("JSON:\n%s", test.JSON)
					continue
				}
			}

			continue
		}

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		if test.ExpectNil {
			if nil != actual {
				t.Errorf("For test #%d, expected nil but actually got non-nil.", testNumber)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
			}
			continue
		}

		if nil == actual {
			t.Errorf("For test #%d, did not expect nil but actually got nil.", testNumber)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		if test.ExpectSkillID {
			if _, ok := actual.(jsonresume.SkillID); !ok {
				t.Errorf("For test #%d, expected type jsonresume.SkillID but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if test.ExpectSkill {
			if _, ok := actual.(jsonresume.Skill); !ok {
				t.Errorf("For test #%d, expected type jsonresume.Skill but actually got %T.", testNumber, actual)
				t.Logf("ACTUAL: %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}

		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
