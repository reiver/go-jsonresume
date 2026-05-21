package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// LanguageID implements a referenced item in the JSON Resume "languages" fields array.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"languages": [
//		"http://example.com/resume/language/algonquian",
//		"http://example.com/resume/language/english",
//		"http://example.com/resume/language/korean",
//		"http://example.com/resume/language/persian",
//		"http://example.com/resume/language/scots",
//	],
//
// LanguageID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Languages = append(cv.Languages, jsonresume.LanguageID("http://example.com/resume/language/algonquian"))
//	cv.Languages = append(cv.Languages, jsonresume.LanguageID("http://example.com/resume/language/english"))
//	cv.Languages = append(cv.Languages, jsonresume.LanguageID("http://example.com/resume/language/korean"))
//	cv.Languages = append(cv.Languages, jsonresume.LanguageID("http://example.com/resume/language/persian"))
//	cv.Languages = append(cv.Languages, jsonresume.LanguageID("http://example.com/resume/language/scots"))
type LanguageID string

func (receiver LanguageID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeLanguage

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver LanguageID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeLanguage

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver LanguageID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeLanguage

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver LanguageID) ProtoLanguage() AnyLanguage {
	const _type string = TypeLanguage

	return AnyLanguage{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

