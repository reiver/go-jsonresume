package jsonresume_test

import (
	"reflect"
	"testing"

	"github.com/reiver/go-jsonresume"
)

func TestJSONResume_SetResumeID(t *testing.T) {

	// SetResumeID on an empty JSONResume sets a single-element slice.
	{
		var jr jsonresume.JSONResume

		jr.SetResumeID("http://example.com/resume/1")

		expected := []jsonresume.ProtoResume{
			jsonresume.SomeResumeID("http://example.com/resume/1"),
		}

		if !reflect.DeepEqual(expected, jr.Resume) {
			t.Errorf("After SetResumeID on empty JSONResume, Resume is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", jr.Resume)
		}
	}

	// SetResumeID replaces an existing Resume slice (does not append).
	{
		var jr jsonresume.JSONResume

		jr.AppendResumeID("http://example.com/resume/old")
		jr.SetResumeID("http://example.com/resume/new")

		expected := []jsonresume.ProtoResume{
			jsonresume.SomeResumeID("http://example.com/resume/new"),
		}

		if !reflect.DeepEqual(expected, jr.Resume) {
			t.Errorf("After SetResumeID on non-empty JSONResume, Resume is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", jr.Resume)
		}

		if len(jr.Resume) != 1 {
			t.Errorf("After SetResumeID, expected exactly 1 element but got %d.", len(jr.Resume))
		}
	}

	// SetResumeID on nil receiver does not panic.
	{
		var jr *jsonresume.JSONResume

		jr.SetResumeID("http://example.com/resume/1") // must not panic
	}
}

func TestJSONResume_AppendResumeID(t *testing.T) {

	// AppendResumeID accumulates entries rather than replacing.
	{
		var jr jsonresume.JSONResume

		jr.AppendResumeID("http://example.com/resume/1")
		jr.AppendResumeID("http://example.com/resume/2")

		expected := []jsonresume.ProtoResume{
			jsonresume.SomeResumeID("http://example.com/resume/1"),
			jsonresume.SomeResumeID("http://example.com/resume/2"),
		}

		if !reflect.DeepEqual(expected, jr.Resume) {
			t.Errorf("After two AppendResumeID calls, Resume is not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", jr.Resume)
		}

		if len(jr.Resume) != 2 {
			t.Errorf("After two AppendResumeID calls, expected 2 elements but got %d.", len(jr.Resume))
		}
	}

	// AppendResumeID on nil receiver does not panic.
	{
		var jr *jsonresume.JSONResume

		jr.AppendResumeID("http://example.com/resume/1") // must not panic
	}
}
