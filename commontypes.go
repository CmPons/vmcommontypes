package commontypes

import "fmt"

type Command struct {
	Identifier    string
	Description   string
	CommandString string    // This is the full command sent to the terminal ex. "ls -a". Use ${Arg1} for variable substitution
	Variables     []Command // This is a list of variables for substitution
}

func (c *Command) Usage() string {
	ret := c.Identifier + "\t" + c.Description

	for _, command := range c.Variables {
		ret += "\t" + command.Usage()
	}

	return ret
}
