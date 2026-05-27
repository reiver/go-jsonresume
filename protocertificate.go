package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoCertificate represents something that is any type 'Certificate', including sub-types.
//
// See also:
//
//	• [AnyCertificate]
//	• [Certificate]
//	• [CertificateID]
type ProtoCertificate interface {
	activitypub.ProtoObject
	ProtoCertificate() AnyCertificate
}

var (
	_ ProtoCertificate = AnyCertificate{}
	_ ProtoCertificate = Certificate{}
	_ ProtoCertificate = CertificateID{}
)
