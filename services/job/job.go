package job

import (
  "rygel/commands" 
  result "rygel/command_result" 
)

type Job struct {
  ResultChan chan result.CommandResult
  Command commands.Command
}

func New(command commands.Command) Job {
  return Job{ResultChan: make(chan result.CommandResult, 1), Command: command}
}
