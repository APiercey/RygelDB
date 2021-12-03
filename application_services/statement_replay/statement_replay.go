package statement_replay

import (
  cs "rygel/core/store"
	"rygel/services/command_executor"
  "rygel/infrastructure/input_parser"
	"rygel/infrastructure/ledger"
  "rygel/services/command_builder"
  "rygel/context"
  "rygel/common"
)

type StatementReplay struct {
  Ledger ledger.Ledger
  CommandExecutor command_executor.CommandExecutor
  CommandBuilder command_builder.CommandBuilder
  StoreRepo cs.StoreRepo
}

func (service StatementReplay) Replay(ctx context.Context) {
  fn := func(line string) {
    params := input_parser.Parse(line)
    store, err := service.StoreRepo.FindByName(ctx.SelectedStore)
    common.HandleErr(err)
    command := service.CommandBuilder.Build(store, params)
    service.CommandExecutor.Enqueue(command)
  }

  service.Ledger.ReplayRecords(fn)
}

