package jsonresume

import (
	"github.com/reiver/go-nul"
)

type CoreCertificate struct {
	Date   nul.Nullable[string] `json:"date"`
	Issuer nul.Nullable[string] `json:"issuer"`
}
