package application

import (
	"example.com/rygel/services"
	"flag"
)

type Application struct {
  BasicAuthService services.BasicAuthService
  StatementExecutionService services.StatementExecutionService
  StorePersistenceService services.StorePersistenceService
}

func New() Application {
  configuredUsername := flag.String("username", "root", "Username, defaults to root")
  configuredPassword := flag.String("password", "password", "Password, defaults to password")

  flag.Parse()

  return Application{
    BasicAuthService: services.BasicAuthService{
      ConfiguredUsername: *configuredUsername,
      ConfiguredPassword: *configuredPassword,
    },
    StatementExecutionService: services.StatementExecutionService{},
    StorePersistenceService: services.StorePersistenceService{
      DiskLocation: "./store.db",
    },
  }
}
