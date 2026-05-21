package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-nul"
)

type CoreSkill struct {
	Keywords activitypub.Strings  `json:"keywords"`
	Level    nul.Nullable[string] `json:"level"`
}
