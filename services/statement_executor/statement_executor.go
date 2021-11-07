package statement_executor

import (
  "rygel/commands" 
  "rygel/services/command_executor"
  "rygel/services/ledger"
)

type StatementExecutor struct {
  CommandExecutor command_executor.CommandExecutor
  Ledger ledger.Ledger
}

func (service StatementExecutor) Execute(statement string) string {
  command := commands.New(statement)
  job := service.CommandExecutor.Enqueue(command)

  result := <- job.ResultChan

  if result.StoreWasUpdated() {
    service.Ledger.AppendRecord(statement)
  }

  return result.CommandResult()
}

