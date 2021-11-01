package application

import (
	"rygel/core"
	"rygel/services"
	cx "rygel/services/command_executor"
	sx "rygel/services/statement_executor"
	"rygel/services/job"
	"flag"
)

type Application struct {
  Store core.Store
  BasicAuthService services.BasicAuthService
  StatementExecutor sx.StatementExecutor
  CommandExecutor cx.CommandExecutor
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

  storePersistenceService := services.StorePersistenceService{
    DiskLocation: "./store.db",
    PersistenceDir: "/tmp",
    Store: &store,
  }

  statementExecutor := sx.StatementExecutor{
    CommandExecutor: &commandExecutor,
    StorePersistenceService: storePersistenceService,
  }

  return Application{
    Store: store,
    BasicAuthService: basicAuthService,
    StatementExecutor: statementExecutor,
    CommandExecutor: &commandExecutor,
  }
}
