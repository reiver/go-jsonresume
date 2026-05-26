package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Certificate implements an embedded item in the JSON Resume "certificates" fields array.
//
// In a JSON Resume document this might look like:
//
//	"certificates": [
//		{
//			"name": "CP+",
//			"date": "2018-06-01",
//			"issuer": "CPEdu",
//			"url": "http://cp.example"
//		},
//		{
//			"name": "Event Sourcing Core",
//			"date": "2017-03-12",
//			"issuer": "EngWorks",
//			"url": "http://engworks.example"
//		}
//	],
//
// Certificate is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Certificates = append(cv.Certificates, jsonresume.Certificate{
//		Name:   nul.Something("CP+"),
//		Date:   nul.Something("2018-06-01"),
//		Issuer: nul.Something("CPEdu"),
//		URL:    activitypub.HRef("http://cp.example"),
//	})
//
//	cv.Certificates = append(cv.Certificates, jsonresume.Certificate{
//		Name:   nul.Something("Event Sourcing Core"),
//		Date:   nul.Something("2017-03-12"),
//		Issuer: nul.Something("EngWorks"),
//		URL:    activitypub.HRef("http://engworks.example"),
//	})
type Certificate struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Certificate"`

	CoreCertificate
}

func (receiver Certificate) ProtoNode() activitypub.AnyNode {
	const _type string = TypeCertificate

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Certificate) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeCertificate

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver Certificate) ProtoObject() activitypub.AnyObject {
	const _type string = TypeCertificate

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
		CoreObject: activitypub.CoreObject{
			URL: receiver.URL,
		},
	}
}

func (receiver Certificate) ProtoCertificate() AnyCertificate {
	const _type string = TypeCertificate

	return AnyCertificate{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreCertificate: receiver.CoreCertificate,
	}
}
