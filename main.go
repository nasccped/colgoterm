package main

import (
	"github.com/nasccped/colgoterm/internals/app"
	"github.com/nasccped/colgoterm/internals/utils"
)

var (
	w string = "-w"
	h string = "-h"
	g string = "-g"
)

// Short hand FlagIdentifier unwraper for surely non-error values.
func flagUnwrapper(long string, short *string, requireValue bool, description string) *utils.FlagIdentifier {
	fi, err := utils.NewFlagIdentifier(long, short, requireValue, description)
	if err != nil {
		panic(err)
	}
	return fi
}

func main() {
	var (
		helpFlag   *utils.FlagIdentifier = flagUnwrapper("--help", nil, false, "Prints the help panel")
		widthFlag  *utils.FlagIdentifier = flagUnwrapper("--width", &w, true, "Set the square width (min = 4)")
		heightFlag *utils.FlagIdentifier = flagUnwrapper("--height", &h, true, "Set the square height (min = 4)")
		gapFlag    *utils.FlagIdentifier = flagUnwrapper("--gap", &g, true, "Set the gap between squares (min = 4)")
	)
	flags := []*utils.FlagIdentifier{
		helpFlag, widthFlag, heightFlag, gapFlag,
	}
	app := app.NewApp(flags)
	app.ShowHelp()
}
