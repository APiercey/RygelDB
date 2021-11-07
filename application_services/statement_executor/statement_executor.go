package statement_executor

import (
	"rygel/commands"
	"rygel/services/command_executor"
	"rygel/infrastructure/input_parser"
	"rygel/infrastructure/ledger"
)

type StatementExecutor struct {
  CommandExecutor command_executor.CommandExecutor
  Ledger ledger.Ledger
}

func (service StatementExecutor) Execute(statement string) string {
  params := input_parser.Parse(statement)
  command := commands.New(params)
  job := service.CommandExecutor.Enqueue(command)

  result := <- job.ResultChan

  if result.StoreWasUpdated() {
    service.Ledger.AppendRecord(statement)
  }

  return result.CommandResult()
}

