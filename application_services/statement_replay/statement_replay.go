package statement_replay

import (
	"rygel/commands"
	"rygel/services/command_executor"
  "rygel/infrastructure/input_parser"
	"rygel/infrastructure/ledger"
)

type StatementReplay struct {
  Ledger ledger.Ledger
  CommandExecutor command_executor.CommandExecutor
}

func (service StatementReplay) Replay() {
  fn := func(line string) {
    params := input_parser.Parse(line)
    command := commands.New(params)
    service.CommandExecutor.Enqueue(command)
  }

  service.Ledger.ReplayRecords(fn)
}

