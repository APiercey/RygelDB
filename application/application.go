package application

import (
	"rygel/core/store_repo"
	sx "rygel/application_services/statement_executor"
	sr "rygel/application_services/statement_replay"
	ba "rygel/services/basic_auth"
	cx "rygel/services/command_executor"
	cb "rygel/services/command_builder"
  "rygel/infrastructure/ledger"
  "rygel/services/command_executor/job" 
	"flag"
  "os"
)

type Application struct {
  StoreRepo store_repo.StoreRepo
  BasicAuth ba.BasicAuth
  StatementExecutor sx.StatementExecutor
  CommandExecutor cx.CommandExecutor
  StatementReplay sr.StatementReplay
  CommandBuilder cb.CommandBuilder
}

func New() Application {
  configuredUsername := flag.String("username", "root", "Username, defaults to root")
  configuredPassword := flag.String("password", "password", "Password, defaults to password")
  flag.Parse()

  basicAuth := ba.BasicAuth{
    ConfiguredUsername: *configuredUsername,
    ConfiguredPassword: *configuredPassword,
  }

  commandExecutor := cx.AsyncCommandExecutor{
    JobQueue: make(chan job.Job),
  }

  f, _ := os.OpenFile("/tmp/rygel-store/store.ledger", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

  ledger := ledger.OnDiskLedger{
    LedgerFile: f,
  }

  storeRepo := store_repo.InitializeFromDir("/tmp/rygel-store")

  statementExecutor := sx.StatementExecutor{
    CommandExecutor: &commandExecutor,
    Ledger: &ledger,
    StoreRepo: storeRepo,
  }

  statementReplay := sr.StatementReplay{
    CommandExecutor: &commandExecutor,
    Ledger: &ledger,
    StoreRepo: storeRepo,
  }

  commandBuilder := cb.CommandBuilder{}

  return Application{
    StoreRepo: storeRepo,
    BasicAuth: basicAuth,
    StatementExecutor: statementExecutor,
    CommandExecutor: &commandExecutor,
    StatementReplay: statementReplay,
    CommandBuilder: commandBuilder,
  }
}
