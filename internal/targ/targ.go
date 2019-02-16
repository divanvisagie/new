// Package targ means Typed Arg
package targ

import "regexp"

// Targ is a typed wrapper around an argument
type Targ struct {
	Arg string
}

// Container is the main storage container for all of this
type Container struct {
	Args []string
}

// Arg gets an arg by position, ignoring flags
func (c *Container) Arg(position int) Targ {
	i := 0
	for i < position {
		match, _ := regexp.Match("^--|^-", []byte(c.Args[i]))
		if !match {
			i++
		}
	}
	return Targ{
		Arg: c.Args[i],
	}
}

// NewContainer creates a new instance of Container from os args
func NewContainer(args []string) *Container {
	return &Container{
		Args: args,
	}
}
