package colored

import (
	"bytes"
	"fmt"
)

const (
	black int = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

// Type alias for colored strings.
type Colored struct {
	fg, bg  *string
	message string
}

// Generates a new `Colored` type value.
func NewColored(message string) *Colored {
	return &Colored{
		fg: nil, bg: nil, message: message,
	}
}

// Func wrapper for Colored field overriding (foreground). Returns itself.
func (c *Colored) overrideFGByFunc(f func(int) string, value int) *Colored {
	temp := f(value)
	c.fg = &temp
	return c
}

// Func wrapper for Colored field overriding (background). Returns itself.
func (c *Colored) overrideBGByFunc(f func(int) string, value int) *Colored {
	temp := f(value)
	c.bg = &temp
	return c
}

func (c *Colored) WithGreenFG() *Colored {
	return c.overrideFGByFunc(nfg, green)
}

func (c *Colored) WithCyanFG() *Colored {
	return c.overrideFGByFunc(nfg, cyan)
}

func (c *Colored) WithBrightGreenFG() *Colored {
	return c.overrideFGByFunc(bfg, green)
}

func (c *Colored) WithBrightCyanFG() *Colored {
	return c.overrideFGByFunc(bfg, cyan)
}

func (c *Colored) WithBlackBG() *Colored {
	return c.overrideBGByFunc(nbg, black)
}

func (c *Colored) WithRedBG() *Colored {
	return c.overrideBGByFunc(nbg, red)
}

func (c *Colored) WithYellowBG() *Colored {
	return c.overrideBGByFunc(nbg, yellow)
}

func (c *Colored) WithBlueBG() *Colored {
	return c.overrideBGByFunc(nbg, blue)
}

func (c *Colored) WithMagentaBG() *Colored {
	return c.overrideBGByFunc(nbg, magenta)
}

func (c *Colored) WithCyanBG() *Colored {
	return c.overrideBGByFunc(nbg, cyan)
}

func (c *Colored) WithWhiteBG() *Colored {
	return c.overrideBGByFunc(nbg, white)
}

func (c *Colored) WithBrightBlackBG() *Colored {
	return c.overrideBGByFunc(bbg, black)
}

func (c *Colored) WithBrightRedBG() *Colored {
	return c.overrideBGByFunc(bbg, red)
}

func (c *Colored) WithBrightYellowBG() *Colored {
	return c.overrideBGByFunc(bbg, yellow)
}

func (c *Colored) WithBrightBlueBG() *Colored {
	return c.overrideBGByFunc(bbg, blue)
}

func (c *Colored) WithBrightMagentaBG() *Colored {
	return c.overrideBGByFunc(bbg, magenta)
}

func (c *Colored) WithBrightCyanBG() *Colored {
	return c.overrideBGByFunc(bbg, cyan)
}

func (c *Colored) WithBrightWhiteBG() *Colored {
	return c.overrideBGByFunc(bbg, white)
}

// Converts the `Colored` value into a string.
func (c *Colored) String() string {
	var buf bytes.Buffer
	var temp *string
	if temp = c.bg; temp != nil {
		buf.WriteString(*temp)
	}
	if temp = c.fg; temp != nil {
		buf.WriteString(*temp)
	}
	return fmt.Sprintf("%s%s%s", buf.String(), c.message, escapeReset)
}
