package jsonresume

import (
	"slices"

	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-json"
)

func validateTypeForUnmarshalJSON(bytes []byte, entity string, validTypes ...string) error {
	if 0 == len(bytes) {
		return nil
	}

	var typeValue string
	err := json.Unmarshal(bytes, &typeValue)
	if nil != err {
		return erorr.Wrap(err, "failed to json-unmarshal "+entity+" type")
	}

	if !slices.Contains(validTypes, typeValue) {
		return erorr.Errorf("jsonresume: unexpected type for %s: %q", entity, typeValue)
	}

	return nil
}
