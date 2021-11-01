package statement_executor

import (
  "rygel/commands" 
  "rygel/input_parser" 
  "rygel/services/command_executor"
  "rygel/services"
)

type StatementExecutor struct {
  CommandExecutor command_executor.CommandExecutor
  StorePersistenceService services.StorePersistenceService
}

func (service StatementExecutor) Execute(statement string) string {
  cmdParameters := input_parser.Parse(statement)
  command := commands.New(cmdParameters)
  job := service.CommandExecutor.Enqueue(command)
  result := <- job.ResultChan

  if result.StoreWasUpdated() {
    service.StorePersistenceService.LogCommand(statement)
  }

  return result.CommandResult()
}

