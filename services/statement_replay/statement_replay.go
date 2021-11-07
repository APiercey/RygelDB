package statement_replay

import (
	"rygel/commands"
	"rygel/input_parser"
	"rygel/services/command_executor"
	"rygel/services/ledger"
)

type StatementReplay struct {
  Ledger ledger.Ledger
  CommandExecutor command_executor.CommandExecutor
}

func (service StatementReplay) Replay() {
  fn := func(line string) {
    cmdParameters := input_parser.Parse(line)
    command := commands.New(cmdParameters)
    service.CommandExecutor.Enqueue(command)
  }

  service.Ledger.ReplayRecords(fn)
}

