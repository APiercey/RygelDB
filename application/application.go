package application

import (
	"rygel/core"
	sx "rygel/application_services/statement_executor"
	sr "rygel/application_services/statement_replay"
	ba "rygel/services/basic_auth"
	cx "rygel/services/command_executor"
  "rygel/infrastructure/ledger"
	"rygel/services/job"
	"flag"
  "os"
)

type Application struct {
  Store core.Store
  BasicAuth ba.BasicAuth
  StatementExecutor sx.StatementExecutor
  CommandExecutor cx.CommandExecutor
  StatementReplay sr.StatementReplay
}

func New() Application {
  configuredUsername := flag.String("username", "root", "Username, defaults to root")
  configuredPassword := flag.String("password", "password", "Password, defaults to password")
  flag.Parse()

  store := core.BuildStore()

  basicAuth := ba.BasicAuth{
    ConfiguredUsername: *configuredUsername,
    ConfiguredPassword: *configuredPassword,
  }

  commandExecutor := cx.AsyncCommandExecutor{
    Store: &store,
    JobQueue: make(chan job.Job),
  }

  f, _ := os.OpenFile("./store.ledger", os.O_RDWR, 0644)

  ledger := ledger.OnDiskLedger{
    LedgerFile: f,
  }

  statementExecutor := sx.StatementExecutor{
    CommandExecutor: &commandExecutor,
    Ledger: &ledger,
  }

  statementReplay := sr.StatementReplay{
    CommandExecutor: &commandExecutor,
    Ledger: &ledger,
  }

  return Application{
    Store: store,
    BasicAuth: basicAuth,
    StatementExecutor: statementExecutor,
    CommandExecutor: &commandExecutor,
    StatementReplay: statementReplay,
  }
}
