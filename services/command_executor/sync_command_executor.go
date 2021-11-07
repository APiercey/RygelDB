package command_executor

import (
  "rygel/commands" 

  "rygel/core" 
  "rygel/services/command_executor/job" 
  "rygel/services/command_executor/command_result" 
)

type SyncCommandExecutor struct {
  Store *core.Store
}

func (service *SyncCommandExecutor) Enqueue(command commands.Command) job.Job {
  job := job.New(command)

  if !command.Valid() {
    job.ResultChan <- command_result.New(false, "Command not valid")
  } else {
    data, storeUpdated := job.Command.Execute(service.Store)
    job.ResultChan <- command_result.New(storeUpdated, data)
  }

  return job
}

func (service *SyncCommandExecutor) Process() bool {
  return false
}


