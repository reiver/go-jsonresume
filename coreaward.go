package jsonresume

import (
	"github.com/reiver/go-nul"
)

type CoreAward struct {
	Awarder nul.Nullable[string] `json:"awarder"`
	Date    nul.Nullable[string] `json:"date"`
	Title   nul.Nullable[string] `json:"title"`
}
