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
	"regexp"
)

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
	description string
}

func isFlag(s string) bool {
	match, _ := regexp.MatchString("^--.*|^-.*", s)
	return match
}

// Arg gets an arg by position, ignoring flags
func (c *Container) Arg(position int) *Targ {
	for i := 0; i < len(c.Args); i++ {
		f := isFlag(c.Args[i])
		if f {
			position++
		}
		if i == position {
			break
		}
	}
	return &Targ{
		Arg: c.Args[position],
	}
}

// Description adds description to container for help printing
func (c *Container) Description(s string) *Container {
	c.description = s
	return c
}

// NewContainer creates a new instance of Container from os args
func NewContainer(args []string) *Container {
	return &Container{
		Args: args,
	}
}
