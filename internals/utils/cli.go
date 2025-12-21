package utils

import (
	"bytes"
	"fmt"
	"strings"
)

// Flag indentifier for value unwraping.
type FlagIdentifier struct {
	// Long name for this flag (includes "--").
	Long string
	// Short name (alias) for this flag (includes "-").
	Short *string
	// If this flag requires a values.
	RequireValue bool
	// Description of what this flag does.
	Description string
}

// Checks if a string is a valid long flag.
func validLong(s string) bool {
	return len(s) >= 3 && strings.HasPrefix(s, "--") && s[2] != '-'
}

// Checks if a string is a valid short flag.
func validShort(s string) bool {
	return len(s) >= 2 && strings.HasPrefix(s, "-") && s[1] != '-'
}

// Returns an error for invalid flag identifiers.
func invalidFlagIds(long string, short *string) error {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Invalid flag identifiers `%s` ", long))
	if short != nil {
		buf.WriteString(fmt.Sprintf("| `%s` ", *short))
	}
	return fmt.Errorf("%s.", buf.String())
}

// Create a new flag identifier with the provided fields.
func NewFlagIdentifier(long string, short *string, requireValue bool, description string) (*FlagIdentifier, error) {
	if !validLong(long) || (short != nil && !validShort(*short)) {
		return nil, invalidFlagIds(long, short)
	}
	return &FlagIdentifier{
		Long:         long,
		Short:        short,
		RequireValue: requireValue,
		Description:  description,
	}, nil
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

// Unwraps the value passed through args based on the flag identifier. If the flag wasn't called,
// returns (nil, nil), but if the flag was called with no value, (nil, error) is returned.
//
// If value isn't required, returns (nil, nil) if everythin ok.
func (fi *FlagIdentifier) Unwrap(args []string) (*string, error) {
	var (
		value        *string = nil
		err          error   = nil
		catchValue   bool    = false
		alreadyCatch bool    = false
	)
	for _, a := range args {
		if catchValue && strings.HasPrefix(a, "-") {
			value, err = nil, fi.followedByAnotherFlag(a)
			break
		} else if catchValue {
			value = &a
			alreadyCatch = true
			catchValue = false
		} else if temp := fi.Short; a == fi.Long || (temp != nil && *temp == a) {
			if !fi.RequireValue {
				alreadyCatch = true
				continue
			} else if alreadyCatch {
				value, err = nil, fi.flagCalledMoreThanOnce()
				break
			}
			catchValue = true
		}
	}
	if catchValue {
		err = fi.missingValue()
	}
	return value, err
}
