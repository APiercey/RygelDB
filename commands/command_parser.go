package commands

import (
  "encoding/json"
  "errors"
  "strconv"
  "strings"
)

func buildDefineCommand(commandStrings []string) (Command, error) {
  if len(commandStrings) > 0 {
    switch commandStrings[0] {
    case "COLLECTION":
      return DefineCollectionCommand{collectionName: commandStrings[1]}, nil
    }
  }

  return nil, errors.New("Error parsing DEFINE COLLECTION statement")
}

func buildInsertCommand(commandStrings []string) (Command, error) {
  if len(commandStrings) > 0 {
    switch commandStrings[0] {
    case "INTO":
      var data map[string]interface{}

      err := json.Unmarshal([]byte(commandStrings[3]), &data)

      if err != nil {
        return nil, errors.New("Could not parse data")
      }

      return InsertCommand{collectionName: commandStrings[1], key: commandStrings[2], data: data}, nil
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

  wps := extractWhereStatements(commandStrings[3:])

  if commandStrings[0] == "all" {
    return FetchCommand{collectionName: commandStrings[2], limit: -1, wherePredicates: wps}, nil
  } else {
    limit, err := strconv.Atoi(commandStrings[0])

    if err != nil {
      return nil, errors.New("Error parsing FETCH statement. Limit not understood") 
    }

    return FetchCommand{collectionName: commandStrings[2], limit: limit, wherePredicates: wps}, nil
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

func buildRemoveCommand(commandStrings []string) (Command, error) {
  if len(commandStrings) > 0 {
    switch commandStrings[0] {
    case "COLLECTION":
      return buildRemoveCollectionCommand(commandStrings)
    case "ITEM":
      return buildRemoveItemCommand(commandStrings)
    }
  }

  return nil, errors.New("Error parsing REMOVE statement.")
}

func buildRemoveItemCommand(commandStrings []string) (Command, error) {
  if len(commandStrings) < 4 {
    return nil, errors.New("Error parsing REMOVE statement")
  }

  if commandStrings[2] != "IN" {
    return nil, errors.New("Error parsing REMOVE statement. Unknown IN.")
  }

  return RemoveItemCommand{collectionName: commandStrings[3], key: commandStrings[1]}, nil
}

func buildRemoveCollectionCommand(commandStrings []string) (Command, error) {
  return RemoveCollectionCommand{collectionName: commandStrings[1]}, nil
}

func extractWhereStatements(commandStrings []string) []wherePredicate {
  wherePredicates := []wherePredicate{}

  for i := 0; i < len(commandStrings) - 3; i += 4 {
    word := commandStrings[i]

    if word == "WHERE" || word == "AND" {
      wp := wherePredicate{path: commandStrings[i + 1], value: commandStrings[i + 3]}

      wherePredicates = append(wherePredicates, wp)
    } else {
      break
    }
  }

  return wherePredicates
}

func CommandParser(rawCommand string) (Command, error) {
  seperatedCommandStrings := strings.Split(rawCommand, " ")

  switch seperatedCommandStrings[0] {
  case "DEFINE":
    return buildDefineCommand(seperatedCommandStrings[1:])
  case "STORE":
    return buildInsertCommand(seperatedCommandStrings[1:])
  case "REMOVE":
    return buildRemoveCommand(seperatedCommandStrings[1:])
  case "FETCH":
    return buildFetchCommand(seperatedCommandStrings[1:])
  case "LOOKUP":
    return buildLookupCommand(seperatedCommandStrings[1:])
  default:
    return nil, errors.New("Unknown instruction.")
  }
}

