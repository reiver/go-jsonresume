package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// CertificateID implements a referenced item in the JSON Resume "certificates" fields array.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"certificates": [
//		"http://example.com/resume/certificate/8",
//		"http://example.com/resume/certificate/7"
//	],
//
// CertificateID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//
//	// ...
//
//	cv.Certificates = append(cv.Certificates, jsonresume.CertificateID("http://example.com/resume/certificate/8"))
//	cv.Certificates = append(cv.Certificates, jsonresume.CertificateID("http://example.com/resume/certificate/7"))
type CertificateID jsonld.ID

func SomeCertificateID(value string) CertificateID {
	return CertificateID(jsonld.SomeID(value))
}

func (receiver CertificateID) MarshalJSON() ([]byte, error) {
	return jsonld.ID(receiver).MarshalJSON()
}

func (receiver *CertificateID) UnmarshalJSON(data []byte) error {
	return (*jsonld.ID)(receiver).UnmarshalJSON(data)
}

func (receiver CertificateID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeCertificate

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver CertificateID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeCertificate

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver CertificateID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeCertificate

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver CertificateID) ProtoCertificate() AnyCertificate {
	const _type string = TypeCertificate

	return AnyCertificate{
		ID:   jsonld.SomeID(jsonld.ID(receiver).GetElse("")),
		Type: jsonld.SomeType(_type),
	}
}

// String makes [CertificateID] fit the [fmt.Stringer] interface.
func (receiver CertificateID) String() string {
	return jsonld.ID(receiver).GetElse("")
}
