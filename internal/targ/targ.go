// Package targ means Typed Arg
package targ

/*
usage: new [<flags>] <project name> <repository>

generate projects from git repositories

Flags:
      --help     Show context-sensitive help (also try --help-long and --help-man).
  -v, --verbose

Args:
  <project name>  Name of the new project
  <repository>    Custom git repo URL or GitHub <username>/<project>
*/

import (
	"fmt"
	"regexp"
)

func getFlag(flags []string, tflag *Tflag) (string, error) {
	for _, flag := range flags {
		if flag == tflag.name || flag == tflag.short {
			return flag, nil
		}
	}
	return "", nil
}

func getArgAtPosition(args []string, pos int) (string, error) {
	for i := 0; i < len(args); i++ {
		f := isFlag(args[i])
		if f {
			pos++
		}
		if i == pos {
			break
		}
	}
	if pos >= len(args) {
		return "", fmt.Errorf("There were not enough arguments")
	}
	return args[pos], nil
}

func isFlag(s string) bool {
	match, _ := regexp.MatchString("^--.*|^-.*", s)
	return match
}

func longestArg(args []*Targ) int {
	longest := 0
	for _, arg := range args {
		l := len(arg.getName())
		if l > longest {
			longest = l
		}
	}
	return longest
}
func longestFlag(args []*Tflag) int {
	longest := 0
	for _, arg := range args {
		l := len(arg.getName())
		if l > longest {
			longest = l
		}
	}
	return longest
}

func padToSize(s string, size int) string {
	for i := len(s); i < size; i++ {
		s = fmt.Sprintf("%s ", s)
	}
	return s
}

// Tflag is a typed wrapper around a flag
type Tflag struct {
	Arg         string //original flag is still an argument actually
	name        string
	short       string //short flag
	description string
}

func (t *Tflag) getName() string {
	return t.name
}

// Short is to supply a short flag
func (t *Tflag) Short(s string) *Tflag {
	t.short = s
	return t
}

// Name give the Tflag a name for help printing
func (t *Tflag) Name(s string) *Tflag {
	t.name = s
	return t
}

// Description sets the description for help printing
func (t *Tflag) Description(s string) *Tflag {
	t.description = s
	return t
}

// Bool gets the boolean value of the flag
func (t *Tflag) Bool() bool {
	if t.Arg == t.name || t.Arg == t.short {
		return true
	}
	return false
}

// Targ is a typed wrapper around an argument
type Targ struct {
	Arg         string
	name        string
	description string
	position    int
}

func (t *Targ) getName() string {
	return t.name
}

// Name gives the targ a name for help printing
func (t *Targ) Name(s string) *Targ {
	t.name = s
	return t
}

// Description gives the targ a description for help printing
func (t *Targ) Description(s string) *Targ {
	t.description = s
	return t
}

func (t *Targ) String() string {
	return t.Arg
}

// Container is the main storage container for all of this
type Container struct {
	Args        []string
	Targs       []*Targ
	Tflags      []*Tflag
	Err         error
	description string
	name        string
}

func (c *Container) getArgs() []string {
	var args []string
	for _, x := range c.Args {
		if !isFlag(x) {
			args = append(args, x)
		}
	}
	return args
}

func (c *Container) getFlags() []string {
	var flags []string
	for _, x := range c.Args {
		if isFlag(x) {
			flags = append(flags, x)
		}
	}
	return flags
}

// Parse all the args in the container
func (c *Container) Parse() {

	args := c.getArgs()
	flags := c.getFlags()
	if len(args) < len(c.Targs) {
		c.Err = fmt.Errorf("There were not enough arguments")
	}

	for i := 0; i < len(c.Targs); i++ {
		p := c.Targs[i].position
		arg, err := getArgAtPosition(args, p)
		if err != nil {
			c.Err = err
		}
		c.Targs[i].Arg = arg
	}

	for _, tflag := range c.Tflags {
		flag, err := getFlag(flags, tflag)
		if err != nil {
			c.Err = err
		}
		tflag.Arg = flag
	}

}

// Flag gets an argument by its flag
func (c *Container) Flag() *Tflag {
	t := &Tflag{}
	c.Tflags = append(c.Tflags, t)
	return t
}

// Arg gets an arg by position, ignoring flags
func (c *Container) Arg(position int) *Targ {
	t := &Targ{
		position: position,
	}
	c.Targs = append(c.Targs, t)
	return t
}

// Help get the help text
func (c *Container) Help() string {
	txt := fmt.Sprintf("usage: %s [<flags>]", c.name)
	for _, arg := range c.Targs {
		if !isFlag(arg.Arg) {
			txt = fmt.Sprintf("%s <%s>", txt, arg.name)
		}
	}

	txt = fmt.Sprintf("%s\n\n%s\n\nArgs:\n", txt, c.description)

	l := longestArg(c.Targs)

	for _, arg := range c.Targs {
		name := padToSize(fmt.Sprintf("<%s>", arg.name), l+2)
		txt = fmt.Sprintf("%s    %s    %s\n", txt, name, arg.description)
	}

	txt = fmt.Sprintf("%s\n\nFlags:\n", txt)
	l = longestFlag(c.Tflags)

	for _, flag := range c.Tflags {
		name := padToSize(fmt.Sprintf("%s, %s", flag.short, flag.name), l+4)
		txt = fmt.Sprintf("%s    %s    %s\n", txt, name, flag.description)
	}
	txt = fmt.Sprintf("%s\n", txt)
	return txt
}

// Name adds the container name for help printing
func (c *Container) Name(s string) *Container {
	c.name = s
	return c
}

// Description adds description to container for help printing
func (c *Container) Description(s string) *Container {
	c.description = s
	return c
}

// NewContainer creates a new instance of Container from os args
func NewContainer(args []string) *Container {
	c := &Container{
		Args: args,
	}
	return c
}
