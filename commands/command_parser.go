package commands

import (
	"encoding/json"
	"errors"

	comp "example.com/rygel/comparisons"
)

type Operation struct {
	Operation string `json:"operation"`
	CollectionName  string `json:"collection_name"`
	Limit int `json:"limit"`
	Predicates []comp.Predicate `json:"where"`
	Data map[string]interface{} `json:"data"`
}

func extractPredicateCollection(operation Operation) comp.PredicateCollection {
  predicates := comp.BuildPredicateCollection()

	for _, wp := range operation.Predicates {
			predicates.AddPredicate(wp)
  }

  return predicates
}

func buildDefineCollectionCommand(operation Operation) (Command, error) {
	return DefineCollectionCommand{collectionName: operation.CollectionName}, nil
}

func buildRemoveCollectionCommand(operation Operation) (Command, error) {
	return RemoveCollectionCommand{collectionName: operation.CollectionName}, nil
}

func buildRemoveItemsCommand(operation Operation) (Command, error) {
	return RemoveItemCommand{
      collectionName: operation.CollectionName,
      limit: operation.Limit,
      predicates: comp.BuildPredicateCollection(),
	}, nil
}

func buildInsertItemCommand(operation Operation) (Command, error) {
	return InsertCommand{
		collectionName: operation.CollectionName,
		data: operation.Data,
	}, nil
}

func buildFetchCommand(operation Operation) (Command, error) {
	return FetchCommand{
		collectionName: operation.CollectionName,
		limit: operation.Limit,
		predicates: extractPredicateCollection(operation),
	}, nil
}

func CommandParser(rawCommand string) (Command, error) {
	operation := Operation{
		Limit: -1,
		Predicates: []comp.Predicate{},
	}

  json.Unmarshal([]byte(rawCommand), &operation)

	switch operation.Operation {
	case "DEFINE COLLECTION":
		return buildDefineCollectionCommand(operation)
	case "REMOVE COLLECTION":
		return buildRemoveCollectionCommand(operation)
	case "REMOVE ITEMS":
		return buildRemoveItemsCommand(operation)
	case "STORE":
		return buildInsertItemCommand(operation)
	case "FETCH":
		return buildFetchCommand(operation)
	default:
		return nil, errors.New("Error, do not understand operation")
	}
}
