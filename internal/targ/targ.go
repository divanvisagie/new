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

func isFlag(s string) bool {
	match, _ := regexp.MatchString("^--.*|^-.*", s)
	return match
}

// Targ is a typed wrapper around an argument
type Targ struct {
	Arg         string
	name        string
	description string
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

// Arg gets an arg by position, ignoring flags
func (c *Container) Arg(position int) *Targ {
	if len(c.Args) == 0 {
		c.Err = fmt.Errorf("Unable to parse arguments becuase there were none")
		t := &Targ{
			Arg: "",
		}
		c.Targs = append(c.Targs, t)
		return t
	}
	for i := 0; i < len(c.Args); i++ {
		f := isFlag(c.Args[i])
		if f {
			position++
		}
		if i == position {
			break
		}
	}
	t := &Targ{
		Arg: c.Args[position],
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

	for _, arg := range c.Targs {
		if !isFlag(arg.Arg) {
			txt = fmt.Sprintf("%s    <%s>    %s\n", txt, arg.name, arg.description)
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
