package command_executor

import (
  "rygel/commands" 

  "rygel/services/command_executor/job" 
  "rygel/services/command_executor/command_result" 
)

type AsyncCommandExecutor struct {
  JobQueue chan job.Job
}

func (service *AsyncCommandExecutor) Enqueue(command commands.Command) job.Job {
  job := job.New(command)

  if !command.Valid() {
    job.ResultChan <- command_result.New(false, "Command not valid")
  } else {
    service.JobQueue <- job
  }

  return job
}

func (service *AsyncCommandExecutor) Process() bool {
  select {
  case job := <- service.JobQueue:
    data, storeUpdated := job.Command.Execute()
    job.ResultChan <- command_result.New(storeUpdated, data)

    return true
  default:
    return false
  }
}
