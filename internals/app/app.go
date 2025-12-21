package app

import (
	"fmt"
	"github.com/nasccped/colgoterm/internals/utils"
	"slices"
)

// App fields when calling the program.
type App struct {
	flags []*utils.FlagIdentifier
}

// Creates a new app struct instance.
func NewApp(flags []*utils.FlagIdentifier) *App {
	return &App{flags: flags}
}

// Returns error when invalid flags are passed.
func (app *App) CheckInvalidFlags(args []string) error {
	for _, a := range args {
		if !slices.ContainsFunc(app.flags, func(fi *utils.FlagIdentifier) bool {
			shortCheck := false
			if temp := fi.Short; temp != nil {
				shortCheck = *temp == a
			}
			return shortCheck || fi.Long == a
		}) {
			return unexpectedFlag(a)
		}
	}
	return nil
}

// When the provided flag wasn't expected.
func unexpectedFlag(flag string) error {
	return fmt.Errorf("The `%s` flag wasn't expected.", flag)
}

func (app *App) ShowHelp() {
	const defaultValuePlaceholder = "<VALUE>"
	var (
		maxLong  int = 0
		maxShort int = 0
		maxValue int = 0
	)
	fmt.Printf("A go lang colored terminal printing.\n\n")
	fmt.Printf("Usage: colgoterm [OPTIONS]\n\n")
	for _, f := range app.flags {
		if temp := len(f.Long); temp > maxLong {
			maxLong = temp
		}
		if temp := f.Short; temp != nil && len(*temp) > maxShort {
			maxShort = len(*temp)
		}
		if defLen := len(defaultValuePlaceholder); f.RequireValue && defLen > maxValue {
			maxValue = defLen
		}
	}
	maxLong++
	fmt.Printf("Options:\n")
	for _, f := range app.flags {
		fmt.Printf("  %s", f.Long)
		fmt.Printf("%*s", maxLong-len(f.Long), " ")
		if temp := f.Short; temp != nil {
			fmt.Printf("| %s", *temp)
			fmt.Printf("%*s", maxShort-len(*temp), " ")
		} else {
			fmt.Printf("%*s", maxShort+3, " ")
		}
		if f.RequireValue {
			fmt.Printf("%s ", defaultValuePlaceholder)
		} else {
			fmt.Printf("%*s ", maxValue, " ")
		}
		fmt.Println(f.Description)
	}
}
