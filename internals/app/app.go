package app

import (
	"fmt"
	"github.com/nasccped/colgoterm/internals/utils"
	"slices"
	"strconv"
	"strings"
)

const (
	defaultWidth  = 8
	dwfaultHeight = 4
	defaultGap    = 4
)

var (
	widthAlias  = "-w"
	heightAlias = "-h"
	gapAlias    = "-g"
)

func (app *App) printHelp() {
	const defaultValuePlaceholder = "<VALUE>"
	var (
		maxLong  int = 0
		maxShort int = 0
		maxValue int = 0
	)
	fmt.Print(`A go lang colored terminal printer.

Usage: colgoterm [OPTIONS]

`)
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
	fmt.Println("Options:")
	for _, f := range app.flags {
		fmt.Printf("  %s%*s", f.Long, maxLong-len(f.Long), " ")
		if temp := f.Short; temp != nil {
			fmt.Printf("| %s%*s", *temp, maxShort-len(*temp), " ")
		} else {
			fmt.Printf("%*s", maxShort+3, " ")
		}
		if f.RequireValue {
			fmt.Printf("%s ", defaultValuePlaceholder)
		} else {
			fmt.Printf("%*s ", maxValue, " ")
		}
		fmt.Printf("%s\n", f.Description)
	}
}

// App fields when calling the program.
type App struct {
	flags []*utils.FlagIdentifier
	args  []string
}

// Creates a new app struct instance.
func NewApp(flags []*utils.FlagIdentifier, args []string) *App {
	return &App{flags: flags, args: args}
}

// Runs the app by the provided flag/args.
func (app *App) Run() error {
	if temp := app.checkFlagsAndValues(); temp != nil {
		return temp
	}
	return nil
}

// If the `--help` flag/alias was called.
func (app *App) containsHelp() bool {
	return slices.ContainsFunc(app.args, func(s string) bool {
		return s == "--help"
	})
}

// Check any unexpected flag or missing value.
func (app *App) checkFlagsAndValues() error {
	if app.containsHelp() && len(app.args) > 1 {
		return utils.InvalidValue(app.args[0], app.args[1], "no flag/value")
	} else if app.containsHelp() {
		app.printHelp()
		return nil
	}
	var (
		w = defaultWidth
		h = dwfaultHeight
		g = defaultGap
	)
	for _, a := range app.args {
		if strings.HasPrefix(a, "-") && !slices.ContainsFunc(app.flags, func(fi *utils.FlagIdentifier) bool {
			return fi.Long == a || (fi.Short != nil && *fi.Short == a)
		}) {
			return utils.InvalidFlag(a)
		}
	}
	for _, f := range app.flags {
		value, err := f.Unwrap(app.args)
		if err != nil {
			return err
		} else if value == nil {
			continue
		}
		intVal, err := strconv.Atoi(*value)
		if err != nil {
			return utils.InvalidValue(f.Long, *value, "<integer value>")
		} else if intVal < 3 || intVal > 100 {
			return utils.InvalidValue(f.Long, *value, "<value between 4 and 100>")
		}
		if f.FlagIs("--width", &widthAlias) {
			w = intVal
		} else if f.FlagIs("--height", &heightAlias) {
			h = intVal
		} else if f.FlagIs("--gap", &gapAlias) {
			g = intVal
		}
	}
	rf := newRuntimeFields(w, h, g)
	rf.run()
	return nil
}
