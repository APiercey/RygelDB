package command_executor

import (
  "rygel/commands" 
  result "rygel/command_result" 

  "rygel/core" 
  "rygel/services/job" 
)

type SyncCommandExecutor struct {
  Store *core.Store
}

func (service *SyncCommandExecutor) Enqueue(command commands.Command) job.Job {
  job := job.New(command)

  if !command.Valid() {
    job.ResultChan <- result.New(false, "Command not valid")
  } else {
    data, storeUpdated := job.Command.Execute(service.Store)
    job.ResultChan <- result.New(storeUpdated, data)
  }

  return job
}

func (service *SyncCommandExecutor) Process() bool {
  return false
}


