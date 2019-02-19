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

func longestArg(args []string) int {
	longest := 0
	for _, arg := range args {
		if !isFlag(arg) {
			l := len(arg)
			if l > longest {
				longest = l
			}
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

// Targ is a typed wrapper around an argument
type Targ struct {
	Arg         string
	name        string
	description string
	position    int
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

// Parse all the args in the container
func (c *Container) Parse() {

	args := c.getArgs()
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

	l := longestArg(c.getArgs())

	for _, arg := range c.Targs {
		if !isFlag(arg.Arg) {
			name := padToSize(fmt.Sprintf("<%s>", arg.name), l+2)
			txt = fmt.Sprintf("%s    %s    %s\n", txt, name, arg.description)
		}
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
