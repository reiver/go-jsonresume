package jsonresume

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
)

func TestSetURL(t *testing.T) {

	const testURL = "https://example.com/something"
	expected := activitypub.HRef(testURL)

	check := func(t *testing.T, name string, url []activitypub.ProtoLink) {
		t.Helper()

		if nil == url {
			t.Errorf("For %s, expected URL to be set but it was nil.", name)
			return
		}

		if 1 != len(url) {
			t.Errorf("For %s, expected URL to have 1 element but actually had %d.", name, len(url))
			return
		}

		if url[0] != expected {
			t.Errorf("For %s, the URL value is not what was expected.", name)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", url[0])
		}
	}

	{
		var core CoreBasics
		core.SetURL(testURL)
		check(t, "CoreBasics", core.URL)
	}

	{
		var core CoreCertificate
		core.SetURL(testURL)
		check(t, "CoreCertificate", core.URL)
	}

	{
		var core CoreEducation
		core.SetURL(testURL)
		check(t, "CoreEducation", core.URL)
	}

	{
		var core CoreExperience
		core.SetURL(testURL)
		check(t, "CoreExperience", core.URL)
	}

	{
		var core CoreProfile
		core.SetURL(testURL)
		check(t, "CoreProfile", core.URL)
	}

	{
		var core CoreProject
		core.SetURL(testURL)
		check(t, "CoreProject", core.URL)
	}

	{
		var core CorePublication
		core.SetURL(testURL)
		check(t, "CorePublication", core.URL)
	}
}

func TestSetURL_nilReceiver(t *testing.T) {

	// These should not panic.
	(*CoreBasics)(nil).SetURL("https://example.com")
	(*CoreCertificate)(nil).SetURL("https://example.com")
	(*CoreEducation)(nil).SetURL("https://example.com")
	(*CoreExperience)(nil).SetURL("https://example.com")
	(*CoreProfile)(nil).SetURL("https://example.com")
	(*CoreProject)(nil).SetURL("https://example.com")
	(*CorePublication)(nil).SetURL("https://example.com")
}
