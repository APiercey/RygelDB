package statement_executor

import (
  "rygel/services/command_builder"
	"rygel/services/command_executor"
	"rygel/infrastructure/input_parser"
	"rygel/infrastructure/ledger"
  cs "rygel/core/store"
  "rygel/common"
  "rygel/context"
)

type StatementExecutor struct {
  CommandExecutor command_executor.CommandExecutor
  CommandBuilder command_builder.CommandBuilder
  Ledger ledger.Ledger
  StoreRepo cs.StoreRepo
}

func (service StatementExecutor) Execute(ctx context.Context, statement string) string {
  params := input_parser.Parse(statement)
  store, err := service.StoreRepo.FindByName(ctx.SelectedStore)

  common.HandleErr(err)

  command := service.CommandBuilder.Build(store, params)

  job := service.CommandExecutor.Enqueue(command)

  result := <- job.ResultChan

  if result.StoreWasUpdated() {
    service.Ledger.AppendRecord(statement)
  }

  return result.CommandResult()
}

