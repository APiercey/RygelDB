package services

import (
  "example.com/rygel/commands" 
  "example.com/rygel/input_parser" 
)

type StatementExecutionService struct {
  CommandExecutor CommandExecutor
}

func (service StatementExecutionService) Execute(statement string) (payload string, store_was_updated bool) {
  cmdParameters := input_parser.Parse(statement)
  command := commands.New(cmdParameters)
  job := service.CommandExecutor.Enqueue(command)
  result := <- job.ResultChan

  return result.CommandResult(), result.StoreWasUpdated()
}

