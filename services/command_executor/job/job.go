package job

import (
  "rygel/commands" 
  "rygel/services/command_executor/command_result" 
)

type Job struct {
  ResultChan chan command_result.CommandResult
  Command commands.Command
}

func New(command commands.Command) Job {
  return Job{ResultChan: make(chan command_result.CommandResult, 1), Command: command}
}
