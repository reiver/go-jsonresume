package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Publication implements an embedded item in the JSON Resume "publications" fields array.
//
// In a JSON Resume document this might look like:
//
//	"publications": [
//		{
//			"name": "Distributed Systems for Fun and Profit",
//			"publisher": "Self",
//			"releaseDate": "2013-09-01",
//			"url": "http://book.example/distsys/",
//			"summary": "A short book about distributed systems."
//		},
//		{
//			"name": "Video Compression for Mere Mortals",
//			"publisher": "Whose Press",
//			"releaseDate": "2014-10-01",
//			"url": "http://example.com/vcfmm",
//			"summary": "An introduction to video compression."
//		}
//	],
//
// Publication is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//
//	// ...
//
//	cv.Publications = append(cv.Publications, jsonresume.Publication{
//		Name:        nul.Something("Distributed Systems for Fun and Profit"),
//		Publisher:   nul.Something("Self"),
//		ReleaseDate: nul.Something("2013-09-01"),
//		URL:         activitypub.HRef("http://book.example/distsys/"),
//		Summary:     nul.Something("A short book about distributed systems."),
//	})
//
//	cv.Publications = append(cv.Publications, jsonresume.Publication{
//		Name:        nul.Something("Video Compression for Mere Mortals"),
//		Publisher:   nul.Something("Whose Press"),
//		ReleaseDate: nul.Something("2014-10-01"),
//		URL:         activitypub.HRef("http://example.com/vcfmm"),
//		Summary:     nul.Something("An introduction to video compression."),
//	})
//
// Publication is for marshaling with a fixed type of "Publication".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Publication", "cv:Publication", or "https://w3id.org/fep/6158#Publication".
//
// For unmarshaling that accepts any type value, use [AnyPublication] instead.
//
// See also:
//
//	• [AnyPublication]
//	• [CorePublication]
//	• [ProtoPublication]
//	• [PublicationID]
//	• [TypePublication]
type Publication struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Publication"`

	CorePublication
}

func (receiver *Publication) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawPublication
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal publication")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			var typeValue string
			err := gojson.Unmarshal(bb, &typeValue)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal publication type")
				return err
			}

			switch typeValue {
			case TypePublication, CompactTypePublication, ExpandedTypePublication:
				// OK
			default:
				return erorr.Errorf("jsonresume: unexpected type for publication: %q", typeValue)
			}
		}
	}

	return receiver.CorePublication.unmarshalRawPublication(raw, &receiver.ID)
}

func (receiver Publication) ProtoNode() activitypub.AnyNode {
	const _type string = TypePublication

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Publication) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypePublication

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver Publication) ProtoObject() activitypub.AnyObject {
	const _type string = TypePublication

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
		CoreObject: activitypub.CoreObject{
			Summary: receiver.Summary,
			URL:     receiver.URL,
		},
	}
}

func (receiver Publication) ProtoPublication() AnyPublication {
	const _type string = TypePublication

	return AnyPublication{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CorePublication: receiver.CorePublication,
	}
}
