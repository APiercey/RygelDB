package commands

import (
	"encoding/json"
)

func CommandParser(input string) (Command, error) {
	cmdParameters := CommandParameters{
		Limit: -1,
	}

  json.Unmarshal([]byte(input), &cmdParameters)

	return New(cmdParameters)
}
