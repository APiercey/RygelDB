package services

import (
  "rygel/commands" 
  "rygel/input_parser" 
  "rygel/services/command_executor"
)

type StatementExecutionService struct {
  CommandExecutor command_executor.CommandExecutor
}

func (service StatementExecutionService) Execute(statement string) (payload string, store_was_updated bool) {
  cmdParameters := input_parser.Parse(statement)
  command := commands.New(cmdParameters)
  job := service.CommandExecutor.Enqueue(command)
  result := <- job.ResultChan

  return result.CommandResult(), result.StoreWasUpdated()
}

