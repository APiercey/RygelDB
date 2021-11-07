package input_parser

import (
	"encoding/json"
	cp "rygel/commands/command_parameters"
)

func Parse(input string) cp.CommandParameters {
	cmdParameters := cp.New()

	err := json.Unmarshal([]byte(input), &cmdParameters)

	if err != nil {
		cmdParameters.Error = err.Error()
	}

	return cmdParameters
}
