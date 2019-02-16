// Package targ means Typed Arg
package targ

import (
	"regexp"
)

// Targ is a typed wrapper around an argument
type Targ struct {
	Arg string
}

// Container is the main storage container for all of this
type Container struct {
	Args []string
}

func isFlag(s string) bool {
	match, _ := regexp.MatchString("--.*", s)
	return match
}

// Arg gets an arg by position, ignoring flags
func (c *Container) Arg(position int) Targ {
	for i := 0; i < len(c.Args); i++ {
		match := isFlag(c.Args[i])
		if match {
			position++
		}
		if i == position {
			break
		}

	}
	println(position)
	return Targ{
		Arg: c.Args[position],
	}
}

// NewContainer creates a new instance of Container from os args
func NewContainer(args []string) *Container {
	return &Container{
		Args: args,
	}
}
