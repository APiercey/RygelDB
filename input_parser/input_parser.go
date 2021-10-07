package input_parser

import (
	"encoding/json"
	"example.com/rygel/commands"
)

func Parse(input string) commands.CommandParameters {
	cmdParameters := commands.NewCommandParameters()

  json.Unmarshal([]byte(input), &cmdParameters)

	return cmdParameters
}
