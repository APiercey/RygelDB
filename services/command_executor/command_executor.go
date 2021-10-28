package command_executor

import (
  "rygel/commands" 
  "rygel/services/job" 
)

type CommandExecutor interface {
  Enqueue(commands.Command) job.Job
  Process() bool
}
