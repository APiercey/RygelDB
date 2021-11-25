package statement_executor

import (
  "rygel/services/command_builder"
	"rygel/services/command_executor"
	"rygel/infrastructure/input_parser"
	"rygel/infrastructure/ledger"
  "rygel/core"
)

type StatementExecutor struct {
  CommandExecutor command_executor.CommandExecutor
  CommandBuilder command_builder.CommandBuilder
  Ledger ledger.Ledger
  StoreRepo core.StoreRepo
}

func (service StatementExecutor) Execute(statement string) string {
  params := input_parser.Parse(statement)
  store := service.StoreRepo.FindByName("make dynamic later")

  command := service.CommandBuilder.Build(store, params)

  job := service.CommandExecutor.Enqueue(command)

  result := <- job.ResultChan

  if result.StoreWasUpdated() {
    service.Ledger.AppendRecord(statement)
  }

  return result.CommandResult()
}

