package application

import (
	"example.com/rygel/core"
	"example.com/rygel/services"
	"example.com/rygel/services/job"
	"flag"
)

type Application struct {
  Store core.Store
  BasicAuthService services.BasicAuthService
  StatementExecutionService services.StatementExecutionService
  CommandExecutor services.CommandExecutor
  StorePersistenceService services.StorePersistenceService
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

  commandExecutor := services.CommandExecutor{
    Store: &store,
    JobQueue: make(chan job.Job),
  }

  statementExectionService := services.StatementExecutionService{
    CommandExecutor: commandExecutor,
  }

  storePersistenceService := services.StorePersistenceService{
    DiskLocation: "./store.db",
    Store: &store,
  }

  return Application{
    Store: store,
    BasicAuthService: basicAuthService,
    StatementExecutionService: statementExectionService,
    CommandExecutor: commandExecutor,
    StorePersistenceService: storePersistenceService,
  }
}
