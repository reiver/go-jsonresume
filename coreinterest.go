package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

type CoreInterest struct {
	Keywords activitypub.Strings `json:"keywords"`
}
