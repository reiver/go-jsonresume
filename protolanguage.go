package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoLanguage represents something that is any type 'Language', including sub-types.
//
// See also:
//
//	• [AnyLanguage]
//	• [Language]
//	• [LanguageID]
type ProtoLanguage interface {
	activitypub.ProtoObject
	ProtoLanguage() AnyLanguage
}

var (
	_ ProtoLanguage = Language{}
	_ ProtoLanguage = LanguageID("")
)
