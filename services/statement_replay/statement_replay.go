package statement_replay

import (
	"rygel/commands"
	"rygel/services/command_executor"
	"rygel/services/ledger"
)

type StatementReplay struct {
  Ledger ledger.Ledger
  CommandExecutor command_executor.CommandExecutor
}

func (service StatementReplay) Replay() {
  fn := func(line string) {
    command := commands.New(line)
    service.CommandExecutor.Enqueue(command)
  }

  service.Ledger.ReplayRecords(fn)
}

