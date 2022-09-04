package commontypes

type Command struct {
	Identifier  string
	Description string

	CommandString string            // This is the full command sent to the terminal ex. "ls -a". Use ${Arg1} for variable substitution
	Variables     map[string]string // This is a list of variables for substitution
}

type CommandList struct {
	Commands []Command
}
