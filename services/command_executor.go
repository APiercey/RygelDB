package services

import (
  "rygel/commands" 
  result "rygel/command_result" 

  "rygel/core" 
  "rygel/services/job" 
)

type CommandExecutor struct {
  JobQueue chan job.Job
  Store *core.Store
}

func (service *CommandExecutor) Enqueue(command commands.Command) job.Job {
  job := job.New(command)

  if !command.Valid() {
    job.ResultChan <- result.New(false, "Command not valid")
  } else {
    service.JobQueue <- job
  }

  return job
}

func (service *CommandExecutor) Process() bool {
  select {
  case job := <- service.JobQueue:
      data, storeUpdated := job.Command.Execute(service.Store)
      job.ResultChan <- result.New(storeUpdated, data)

      return true
    default:
      return false
  }
}
