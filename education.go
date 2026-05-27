package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Education implements an embedded item in the JSON Resume "education" fields array.
//
// In a JSON Resume document this might look like:
//
//	"education": [
//		{
//			"institution": "SFU",
//			"url": "https://www.sfu.ca",
//			"area": "computer science",
//			"studyType": "Bachelor of Science (B.Sc.)",
//			"startDate": "1995",
//			"endDate": "2000",
//			"score": "4.0",
//			"courses": [
//				"MACM 101 — Discrete Mathematics I",
//				"MATH 151 — Calculus I",
//				"MATH 152 — Calculus II",
//				"MACM 201 — Discrete Mathematics II",
//				"CMPT 201 — Systems Programming",
//				"CMPT 225 — Data Structures and Programming",
//				"MATH 251 — Calculus III",
//				"CMPT 303 - Operating Systems",
//				"CMPT 307 - Data Structures and Algorithms",
//				"CMPT 310 — Introduction to Artificial Intelligence",
//				"MACM 316 — Numerical Analysis I",
//				"CMPT 379 - Principles of Compiler Design",
//				"CMPT 404 - Cryptography and Cryptographic Protocols",
//				"CMPT 407 - Computational Complexity",
//				"CMPT 410 - Machine Learning",
//				"MACM 442 - Cryptography"
//			]
//		},
//		{
//			"institution": "KPU",
//			"url": "https://www.kpu.ca",
//			"area": "mathematics",
//			"studyType": "Certificate",
//			"startDate": "1993",
//			"endDate": "1994",
//			"score": "4.0",
//			"courses": [
//				"PHYS 1100 - Introductory Physics",
//				"MATH 1120 - Differential Calculus",
//				"PHYS 1170 - Mechanics I",
//				"MATH 1220 - Integral Calculus",
//				"PHYS 2010 - Modern Physics",
//				"PHYS 2030 - Classical Mechanics",
//				"MATH 2321 - Multivariate Calculus (Calculus III)",
//				"MATH 2232 - Linear Algebra",
//				"PHYS 2330 - Intermediate Mechanics",
//				"PHYS 2420 - Electricity and Magnetism",
//				"MATH 2821 - Multivariate and Vector Calculus"
//			]
//		}
//	],
//
// Education is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	education1 := jsonresume.Education{
//		Institution: nul.Something("SFU"),
//		Area:        activitypub.SomeString("computer science"),
//		StudyType:   activitypub.SomeString("Bachelor of Science (B.Sc.)"),
//		StartDate:   nul.Something("1995"),
//		EndDate:     nul.Something("2000"),
//		Score:       nul.Something("4.0"),
//		Courses:     activitypub.SomeStrings(
//			"MACM 101 — Discrete Mathematics I",
//			"MATH 151 — Calculus I",
//			"MATH 152 — Calculus II",
//			"MACM 201 — Discrete Mathematics II",
//			"CMPT 201 — Systems Programming",
//			"CMPT 225 — Data Structures and Programming",
//			"MATH 251 — Calculus III",
//			"CMPT 303 - Operating Systems",
//			"CMPT 307 - Data Structures and Algorithms",
//			"CMPT 310 — Introduction to Artificial Intelligence",
//			"MACM 316 — Numerical Analysis I",
//			"CMPT 379 - Principles of Compiler Design",
//			"CMPT 404 - Cryptography and Cryptographic Protocols",
//			"CMPT 407 - Computational Complexity",
//			"CMPT 410 - Machine Learning",
//			"MACM 442 - Cryptography",
//		),
//	}
//	education1.SetURL("https://www.sfu.ca")
//
//	cv.Education = append(cv.Education, education1)
//
//	education2 := jsonresume.Education{
//		Institution: nul.Something("KPU"),
//		Area:        activitypub.SomeString("mathematics"),
//		StudyType:   activitypub.SomeString("Certificate"),
//		StartDate:   nul.Something("1993"),
//		EndDate:     nul.Something("1994"),
//		Score:       nul.Something("4.0"),
//		Courses:     activitypub.SomeStrings(
//			"PHYS 1100 - Introductory Physics",
//			"MATH 1120 - Differential Calculus",
//			"PHYS 1170 - Mechanics I",
//			"MATH 1220 - Integral Calculus",
//			"PHYS 2010 - Modern Physics",
//			"PHYS 2030 - Classical Mechanics",
//			"MATH 2321 - Multivariate Calculus (Calculus III)",
//			"MATH 2232 - Linear Algebra",
//			"PHYS 2330 - Intermediate Mechanics",
//			"PHYS 2420 - Electricity and Magnetism",
//			"MATH 2821 - Multivariate and Vector Calculus",
//		),
//	}
//	education2.SetURL("https://www.kpu.ca")
//
//	cv.Education = append(cv.Education, education2)
//
// Education is for marshaling with a fixed type of "Education".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Education", "cv:Education", or "https://w3id.org/fep/6158#Education".
//
// For unmarshaling that accepts any type value, use [AnyEducation] instead.
//
// See also:
//
//	• [AnyEducation]
//	• [CoreEducation]
//	• [EducationID]
//	• [ProtoEducation]
//	• [TypeEducation]
type Education struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Education"`

	CoreEducation
}

func (receiver *Education) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawEducation
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal education")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			var typeValue string
			err := gojson.Unmarshal(bb, &typeValue)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal education type")
				return err
			}

			switch typeValue {
			case TypeEducation, CompactTypeEducation, ExpandedTypeEducation:
				// OK
			default:
				return erorr.Errorf("jsonresume: unexpected type for education: %q", typeValue)
			}
		}
	}

	return receiver.CoreEducation.unmarshalRawEducation(raw, &receiver.ID)
}

func (receiver Education) ProtoNode() activitypub.AnyNode {
	const _type string = TypeEducation

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Education) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeEducation

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Education) ProtoObject() activitypub.AnyObject {
	const _type string = TypeEducation

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreObject: activitypub.CoreObject{
			URL: receiver.URL,
		},
	}
}

func (receiver Education) ProtoEducation() AnyEducation {
	const _type string = TypeEducation

	return AnyEducation{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEducation: receiver.CoreEducation,
	}
}
