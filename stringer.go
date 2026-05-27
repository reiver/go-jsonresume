package jsonresume

import (
	gobytes "bytes"
	gojson "encoding/json"

	"github.com/reiver/go-jsonld"
)

func stringJSON(value any) string {
	var buffer gobytes.Buffer

	bytes, err := jsonld.Marshal(value)
	if nil != err {
		return "{}"
	}

	err = gojson.Indent(&buffer, bytes, "", "  ")
	if nil != err {
		return "{}"
	}

	return buffer.String()
}

func (receiver Award) String() string {
	return stringJSON(receiver)
}

func (receiver AnyAward) String() string {
	return stringJSON(receiver)
}

func (receiver Basics) String() string {
	return stringJSON(receiver)
}

func (receiver AnyBasics) String() string {
	return stringJSON(receiver)
}

func (receiver Certificate) String() string {
	return stringJSON(receiver)
}

func (receiver AnyCertificate) String() string {
	return stringJSON(receiver)
}

func (receiver Education) String() string {
	return stringJSON(receiver)
}

func (receiver AnyEducation) String() string {
	return stringJSON(receiver)
}

func (receiver Experience) String() string {
	return stringJSON(receiver)
}

func (receiver AnyExperience) String() string {
	return stringJSON(receiver)
}

func (receiver Interest) String() string {
	return stringJSON(receiver)
}

func (receiver AnyInterest) String() string {
	return stringJSON(receiver)
}

func (receiver Language) String() string {
	return stringJSON(receiver)
}

func (receiver AnyLanguage) String() string {
	return stringJSON(receiver)
}

func (receiver Location) String() string {
	return stringJSON(receiver)
}

func (receiver AnyLocation) String() string {
	return stringJSON(receiver)
}

func (receiver Meta) String() string {
	return stringJSON(receiver)
}

func (receiver AnyMeta) String() string {
	return stringJSON(receiver)
}

func (receiver Profile) String() string {
	return stringJSON(receiver)
}

func (receiver AnyProfile) String() string {
	return stringJSON(receiver)
}

func (receiver Project) String() string {
	return stringJSON(receiver)
}

func (receiver AnyProject) String() string {
	return stringJSON(receiver)
}

func (receiver Publication) String() string {
	return stringJSON(receiver)
}

func (receiver AnyPublication) String() string {
	return stringJSON(receiver)
}

func (receiver Reference) String() string {
	return stringJSON(receiver)
}

func (receiver AnyReference) String() string {
	return stringJSON(receiver)
}

func (receiver Resume) String() string {
	return stringJSON(receiver)
}

func (receiver AnyResume) String() string {
	return stringJSON(receiver)
}

func (receiver Skill) String() string {
	return stringJSON(receiver)
}

func (receiver AnySkill) String() string {
	return stringJSON(receiver)
}
