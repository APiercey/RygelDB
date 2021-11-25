package statement_replay

import (
  "rygel/core"
	"rygel/services/command_executor"
  "rygel/infrastructure/input_parser"
	"rygel/infrastructure/ledger"
  "rygel/services/command_builder"
  "rygel/context"
)

type StatementReplay struct {
  Ledger ledger.Ledger
  CommandExecutor command_executor.CommandExecutor
  CommandBuilder command_builder.CommandBuilder
  StoreRepo core.StoreRepo
}

func (service StatementReplay) Replay(ctx context.Context) {
  fn := func(line string) {
    params := input_parser.Parse(line)
    store := service.StoreRepo.FindByName(ctx.SelectedStore)
    command := service.CommandBuilder.Build(store, params)
    service.CommandExecutor.Enqueue(command)
  }

  service.Ledger.ReplayRecords(fn)
}

