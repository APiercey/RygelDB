package commands

import (
	"encoding/json"
	"errors"

	comp "example.com/rygel/comparisons"
)

type commandData struct {
	Operation string `json:"operation"`
	CollectionName  string `json:"collection_name"`
	Limit int `json:"limit"`
	Predicates []comp.Predicate `json:"where"`
	Data map[string]interface{} `json:"data"`
}

func extractPredicateCollection(cmdData commandData) comp.PredicateCollection {
  predicates := comp.BuildPredicateCollection()

	for _, wp := range cmdData.Predicates {
			predicates.AddPredicate(wp)
  }

  return predicates
}

func buildDefineCollectionCommand(cmdData commandData) (Command, error) {
	return DefineCollectionCommand{collectionName: cmdData.CollectionName}, nil
}

func buildRemoveCollectionCommand(cmdData commandData) (Command, error) {
	return RemoveCollectionCommand{collectionName: cmdData.CollectionName}, nil
}

func buildRemoveItemsCommand(cmdData commandData) (Command, error) {
	return RemoveItemCommand{
      collectionName: cmdData.CollectionName,
      limit: cmdData.Limit,
      predicates: comp.BuildPredicateCollection(),
	}, nil
}

func buildInsertItemCommand(cmdData commandData) (Command, error) {
	return InsertCommand{
		collectionName: cmdData.CollectionName,
		data: cmdData.Data,
	}, nil
}

func buildFetchCommand(cmdData commandData) (Command, error) {
	return FetchCommand{
		collectionName: cmdData.CollectionName,
		limit: cmdData.Limit,
		predicates: extractPredicateCollection(cmdData),
	}, nil
}

func CommandParser(input string) (Command, error) {
	cmdData := commandData{
		Limit: -1,
		Predicates: []comp.Predicate{},
	}

  json.Unmarshal([]byte(input), &cmdData)

	switch cmdData.Operation {
	case "DEFINE COLLECTION":
		return buildDefineCollectionCommand(cmdData)
	case "REMOVE COLLECTION":
		return buildRemoveCollectionCommand(cmdData)
	case "REMOVE ITEMS":
		return buildRemoveItemsCommand(cmdData)
	case "STORE":
		return buildInsertItemCommand(cmdData)
	case "FETCH":
		return buildFetchCommand(cmdData)
	default:
		return nil, errors.New("Error, do not understand input")
	}
}
