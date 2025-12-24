package utils

import (
	"bytes"
	"fmt"
)

// When the value passed isn't valid.
func InvalidValue(flag, value, expects string) error {
	return fmt.Errorf(
		"the \"%s\" isn't valid for the `%s` flag.\n"+
			"\n"+
			"This flag expects `%s`.",
		value,
		flag,
		expects,
	)
}

// When passing an unexpected flag.
func InvalidFlag(f string) error {
	return fmt.Errorf("the `%s` flag isn't valid.", f)
}

// Generates an error message when flag/alias is called more than once.
func (flagId *FlagIdentifier) flagCalledMoreThanOnce() error {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("The `%s` ", flagId.Long))
	if temp := flagId.Short; temp != nil {
		buf.WriteString(fmt.Sprintf("| `%s` ", *temp))
	}
	buf.WriteString("flag was called more than once.")
	return fmt.Errorf("%s", buf.String())
}

// When flag is called but there's no value for it.
func (flagId *FlagIdentifier) missingValue() error {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("The `%s` ", flagId.Long))
	if temp := flagId.Short; temp != nil {
		buf.WriteString(fmt.Sprintf("| `%s` ", *temp))
	}
	buf.WriteString("flag was called but with no value.")
	return fmt.Errorf("%s", buf.String())
}

// When a flag is followed by another one instead of a value.
func (flagId *FlagIdentifier) followedByAnotherFlag(otherFlag string) error {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("A value for `%s` ", flagId.Long))
	if temp := flagId.Short; temp != nil {
		buf.WriteString(fmt.Sprintf("| `%s` ", *temp))
	}
	buf.WriteString(fmt.Sprintf("flag was expected, but received another flag (`%s`).", otherFlag))
	return fmt.Errorf("%s", buf.String())
}
