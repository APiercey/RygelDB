package application

import (
	"rygel/core"
	"rygel/services"
	cx "rygel/services/command_executor"
	sx "rygel/services/statement_executor"
	sr "rygel/services/statement_replay"
  "rygel/services/ledger"
	"rygel/services/job"
	"flag"
  "os"
)

type Application struct {
  Store core.Store
  BasicAuthService services.BasicAuthService
  StatementExecutor sx.StatementExecutor
  CommandExecutor cx.CommandExecutor
  StatementReplay sr.StatementReplay
}

func New() Application {
  configuredUsername := flag.String("username", "root", "Username, defaults to root")
  configuredPassword := flag.String("password", "password", "Password, defaults to password")
  flag.Parse()

  store := core.BuildStore()

  basicAuthService := services.BasicAuthService{
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
    BasicAuthService: basicAuthService,
    StatementExecutor: statementExecutor,
    CommandExecutor: &commandExecutor,
    StatementReplay: statementReplay,
  }
}
