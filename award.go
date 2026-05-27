package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Award implements an embedded item in the JSON Resume "awards" fields array.
//
// In a JSON Resume document this might look like:
//
//	"awards": [
//		{
//			"title": "Best Employee (2024)",
//			"date": "2024-05-21",
//			"awarder": "SuperCo",
//			"summary": "He did good work."
//		},
//		{
//			"title": "Acme Excellence (2021)",
//			"date": "2021-01-17",
//			"awarder": "Acme",
//			"summary": "For Joe Blow's excellent work."
//		}
//	],
//
// Award is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Awards = append(cv.Awards, jsonresume.Award{
//		Title:   nul.Something("Best Employee (2024)"),
//		Date:    nul.Something("2024-05-21"),
//		Awarder: nul.Something("SuperCo"),
//		Summary: nul.Something("He did good work."),
//	})
//
//	cv.Awards = append(cv.Awards, jsonresume.Award{
//		Title:   nul.Something("Acme Excellence (2021)"),
//		Date:    nul.Something("2021-01-17"),
//		Awarder: nul.Something("Acme"),
//		Summary: nul.Something("For Joe Blow's excellent work."),
//	})
//
// Award is for marshaling with a fixed type of "Award".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Award", "cv:Award", or "https://w3id.org/fep/6158#Award".
//
// For unmarshaling that accepts any type value, use [AnyAward] instead.
//
// See also:
//
//	• [AnyAward]
//	• [AwardID]
//	• [CoreAward]
//	• [ProtoAward]
//	• [TypeAward]
type Award struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Award"`

	CoreAward
}

func (receiver *Award) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawAward
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal award")
		return err
	}

	if err := validateTypeForUnmarshalJSON([]byte(raw.Type), "award", TypeAward, CompactTypeAward, ExpandedTypeAward); nil != err {
		return err
	}

	return receiver.CoreAward.unmarshalRawAward(raw, &receiver.ID)
}

func (receiver Award) ProtoNode() activitypub.AnyNode {
	const _type string = TypeAward

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Award) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeAward

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Award) ProtoObject() activitypub.AnyObject {
	const _type string = TypeAward

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreObject: activitypub.CoreObject{
			Summary: receiver.Summary,
		},
	}
}

func (receiver Award) ProtoAward() AnyAward {
	const _type string = TypeAward

	return AnyAward{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreAward: receiver.CoreAward,
	}
}
