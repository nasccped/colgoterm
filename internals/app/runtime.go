package app

import (
	"bytes"
	"fmt"

	"github.com/nasccped/colgoterm/internals/colored"
)

type col = colored.Colored

// Fields loaded at runtime.
type runtimeFields struct {
	width, height, gap int
}

// Create a new `runtimeFields` object.
func newRuntimeFields(width, height, gap int) *runtimeFields {
	return &runtimeFields{
		width:  width,
		height: height,
		gap:    gap,
	}
}

// Does the job by the runtime vars.
func (rf *runtimeFields) run() {
	funcs := []*functionPointer{
		newFunctionPointer((*col).WithBlackBG, (*col).WithBrightBlackBG),
		newFunctionPointer((*col).WithRedBG, (*col).WithBrightRedBG),
		newFunctionPointer((*col).WithYellowBG, (*col).WithBrightYellowBG),
		newFunctionPointer((*col).WithBlueBG, (*col).WithBrightBlueBG),
		newFunctionPointer((*col).WithMagentaBG, (*col).WithBrightMagentaBG),
		newFunctionPointer((*col).WithCyanBG, (*col).WithBrightCyanBG),
		newFunctionPointer((*col).WithWhiteBG, (*col).WithBrightWhiteBG),
	}
	fmt.Println()
	for y := range rf.height {
		fmt.Print("  ")
		for _, fp := range funcs {
			gapAdd := 0
			coloredItem := colored.NewColored(fmt.Sprintf("%*s", rf.width, " "))
			bright := fp.bright(coloredItem).String()
			shadow := fp.normal(colored.NewColored("  ")).String()
			fmt.Printf("%s", bright)
			if y > 0 {
				gapAdd -= 2
				fmt.Printf("%s", shadow)
			}
			fmt.Printf("%*s", rf.gap+gapAdd, " ")
		}
		fmt.Println()
	}
	fmt.Print("    ")
	for _, fp := range funcs {
		shadow := fp.normal(colored.NewColored(string(bytes.Repeat([]byte{' '}, rf.width)))).String()
		fmt.Printf("%s%*s", shadow, rf.gap, " ")
	}
	fmt.Print("\n\n")
}
