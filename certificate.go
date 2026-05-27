package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
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
//	certificate1 := jsonresume.Certificate{
//		Name:   nul.Something("CP+"),
//		Date:   nul.Something("2018-06-01"),
//		Issuer: nul.Something("CPEdu"),
//	}
//	certificate1.SetURL("http://cp.example")
//
//	cv.Certificates = append(cv.Certificates, certificate1)
//
//	certificate2 := jsonresume.Certificate{
//		Name:   nul.Something("Event Sourcing Core"),
//		Date:   nul.Something("2017-03-12"),
//		Issuer: nul.Something("EngWorks"),
//	}
//	certificate2.SetURL("http://engworks.example")
//
//	cv.Certificates = append(cv.Certificates, certificate2)
//
// Certificate is for marshaling with a fixed type of "Certificate".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Certificate", "cv:Certificate", or "https://w3id.org/fep/6158#Certificate".
//
// For unmarshaling that accepts any type value, use [AnyCertificate] instead.
//
// See also:
//
//	• [AnyCertificate]
//	• [CoreCertificate]
//	• [CertificateID]
//	• [ProtoCertificate]
//	• [TypeCertificate]
type Certificate struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Certificate"`

	CoreCertificate
}

func (receiver *Certificate) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawCertificate
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal certificate")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			var typeValue string
			err := gojson.Unmarshal(bb, &typeValue)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal certificate type")
				return err
			}

			switch typeValue {
			case TypeCertificate, CompactTypeCertificate, ExpandedTypeCertificate:
				// OK
			default:
				return erorr.Errorf("jsonresume: unexpected type for certificate: %q", typeValue)
			}
		}
	}

	return receiver.CoreCertificate.unmarshalRawCertificate(raw, &receiver.ID)
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
