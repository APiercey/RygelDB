package services

import (
  "example.com/rygel/commands" 
  "example.com/rygel/input_parser" 

  "example.com/rygel/core" 
)

type StatementExecutionService struct {}

func (service *StatementExecutionService) Execute(store *core.Store, statement string) (result string, store_was_updated bool) {
  cmdParameters := input_parser.Parse(statement)
  command := commands.New(cmdParameters)

  if !command.Valid() {
    return "Command not valid", false
  }

  return command.Execute(store)
}

