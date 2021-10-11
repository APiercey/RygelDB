package input_parser

import (
	"encoding/json"
	"example.com/rygel/commands"
)

func Parse(input string) commands.CommandParameters {
	cmdParameters := commands.NewCommandParameters()

	err := json.Unmarshal([]byte(input), &cmdParameters)

	if err != nil {
		cmdParameters.Error = err.Error()
	}

	return cmdParameters
}
