package job

import (
  "example.com/rygel/commands" 
  "example.com/rygel/result" 
)

type Job struct {
  ResultChan chan result.Result
  Command commands.Command
}

func New(command commands.Command) Job {
  return Job{ResultChan: make(chan result.Result, 1), Command: command}
}
