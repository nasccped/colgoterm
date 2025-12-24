package main

import (
	"fmt"
	"os"
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

// Prints the error message + exit with error code (`1`).
func printErrAndExit(err error) {
	err = fmt.Errorf("\x1b[91merror\x1b[0m: %s", err)
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	var (
		helpFlag   *utils.FlagIdentifier = flagUnwrapper("--help", nil, false, "Prints the help panel")
		widthFlag  *utils.FlagIdentifier = flagUnwrapper("--width", &w, true, "Set the square width (default = 8)")
		heightFlag *utils.FlagIdentifier = flagUnwrapper("--height", &h, true, "Set the square height (default = 4)")
		gapFlag    *utils.FlagIdentifier = flagUnwrapper("--gap", &g, true, "Set the gap between squares (default = 4)")
	)
	flags := []*utils.FlagIdentifier{
		helpFlag, widthFlag, heightFlag, gapFlag,
	}
	app := app.NewApp(flags, os.Args[1:])
	if temp := app.Run(); temp != nil {
		printErrAndExit(temp)
	}
}
