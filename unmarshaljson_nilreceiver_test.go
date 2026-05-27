package jsonresume

import (
	"errors"
	"testing"
)

func TestUnmarshalJSON_nilReceiver(t *testing.T) {

	tests := []struct {
		Name string
		Fn   func([]byte) error
	}{
		// 0
		{
			Name: "Award",
			Fn:   (*Award)(nil).UnmarshalJSON,
		},

		// 1
		{
			Name: "Basics",
			Fn:   (*Basics)(nil).UnmarshalJSON,
		},

		// 2
		{
			Name: "Certificate",
			Fn:   (*Certificate)(nil).UnmarshalJSON,
		},

		// 3
		{
			Name: "Education",
			Fn:   (*Education)(nil).UnmarshalJSON,
		},

		// 4
		{
			Name: "Experience",
			Fn:   (*Experience)(nil).UnmarshalJSON,
		},

		// 5
		{
			Name: "Interest",
			Fn:   (*Interest)(nil).UnmarshalJSON,
		},

		// 6
		{
			Name: "Language",
			Fn:   (*Language)(nil).UnmarshalJSON,
		},

		// 7
		{
			Name: "Location",
			Fn:   (*Location)(nil).UnmarshalJSON,
		},

		// 8
		{
			Name: "Meta",
			Fn:   (*Meta)(nil).UnmarshalJSON,
		},

		// 9
		{
			Name: "Profile",
			Fn:   (*Profile)(nil).UnmarshalJSON,
		},

		// 10
		{
			Name: "Project",
			Fn:   (*Project)(nil).UnmarshalJSON,
		},

		// 11
		{
			Name: "Publication",
			Fn:   (*Publication)(nil).UnmarshalJSON,
		},

		// 12
		{
			Name: "Reference",
			Fn:   (*Reference)(nil).UnmarshalJSON,
		},

		// 13
		{
			Name: "Resume",
			Fn:   (*Resume)(nil).UnmarshalJSON,
		},

		// 14
		{
			Name: "Skill",
			Fn:   (*Skill)(nil).UnmarshalJSON,
		},
	}

	for testNumber, test := range tests {

		err := test.Fn([]byte(`{}`))

		if nil == err {
			t.Errorf("For test #%d (%s), expected an error but did not actually get one.", testNumber, test.Name)
			continue
		}

		if !errors.Is(err, ErrReceiverNil) {
			t.Errorf("For test #%d (%s), expected ErrReceiverNil but got a different error.", testNumber, test.Name)
			t.Logf("ACTUAL-ERROR: %s", err)
			continue
		}
	}
}
