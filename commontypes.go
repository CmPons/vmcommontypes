package commontypes

type Command struct {
	Identifier    string
	Description   string
	CommandString string    // This is the full command sent to the terminal ex. "ls -a". Use ${Arg1} for variable substitution
	Variables     []Command // This is a list of variables for substitution
}

func (c *Command) Usage() string {
	ret := c.Identifier + "\t\t\t\t" + c.Description + "\n"

	for _, command := range c.Variables {
		ret += "\t\t\t\t" + command.Usage()
	}

	return ret
}

type CommandList []Command

func (cl CommandList) Usage() string {
	var ret string
	for _, command := range cl {
		ret += command.Usage() + "\n"
	}

	return ret
}
