package main

import (
	"errors"
	"strconv"
	"strings"
)

func buildDefineCommand(commandStrings []string) (Command, error) {
  if len(commandStrings) > 0 {
    switch commandStrings[0] {
    case "COLLECTION":
      return CreateCollectionCommand{collectionName: commandStrings[1]}, nil
    }
  }

  return nil, errors.New("Error parsing DEFINE COLLECTION statement")
}

func buildInsertCommand(commandStrings []string) (Command, error) {
  if len(commandStrings) > 0 {
    switch commandStrings[0] {
    case "INTO":
      return InsertCommand{collectionName: commandStrings[1], key: commandStrings[2], data: commandStrings[3]}, nil
    }
  }

  return nil, errors.New("Error parsing STORE INTO statement")
}

func buildFetchCommand(commandStrings []string) (Command, error) {
  if len(commandStrings) < 3 {
    return nil, errors.New("Error parsing FETCH statement")
  }

  if commandStrings[1] != "FROM" {
    return nil, errors.New("Error parsing FETCH statement. Unknown FROM.")
  }

  if commandStrings[0] == "all" {
    return FetchCommand{collectionName: commandStrings[2], limit: -1}, nil
  } else {
    limit, err := strconv.Atoi(commandStrings[0])

    if err != nil {
      return nil, errors.New("Error parsing FETCH statement. Limit not understood") 
    }

    return FetchCommand{collectionName: commandStrings[2], limit: limit}, nil
  }
}

func buildLookupCommand(commandStrings []string) (Command, error) {
  if len(commandStrings) < 3 {
    return nil, errors.New("Error parsing LOOKUP statement")
  }

  if commandStrings[1] != "IN" {
    return nil, errors.New("Error parsing LOOKUP statement. Unknown IN.")
  }

  return LookupCommand{collectionName: commandStrings[2], key: commandStrings[0]}, nil
}

func CommandParser(rawCommand string) (Command, error) {
  seperatedCommandStrings := strings.Split(rawCommand, " ")

  switch seperatedCommandStrings[0] {
  case "DEFINE":
    return buildDefineCommand(seperatedCommandStrings[1:])
  case "STORE":
    return buildInsertCommand(seperatedCommandStrings[1:])
  case "FETCH":
    return buildFetchCommand(seperatedCommandStrings[1:])
  case "LOOKUP":
    return buildLookupCommand(seperatedCommandStrings[1:])
  default:
    return nil, errors.New("Unknown instruction.")
  }
}

