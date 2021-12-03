package input_parser

// NOTE: Should this be apart of command_parameters?
import (
	"encoding/json"
	cp "rygel/services/command_builder/command_parameters"
)

func Parse(input string) cp.CommandParameters {
	cmdParameters := cp.New()

	err := json.Unmarshal([]byte(input), &cmdParameters)

	if err != nil {
		cmdParameters.Error = err.Error()
	}

	return cmdParameters
}
