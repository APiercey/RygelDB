package input_parser

import (
	"encoding/json"

	"example.com/rygel/commands"
)

func Parse(input string) commands.CommandParameters {
	cmdParameters := commands.CommandParameters{
		Limit: -1,
	}

  json.Unmarshal([]byte(input), &cmdParameters)

	return cmdParameters
}
