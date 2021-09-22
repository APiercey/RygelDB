package commands

import (
	"errors"

	"example.com/rygel/common"
	"example.com/rygel/core"
	comp "example.com/rygel/comparisons"
)

type Command interface {
  Execute(store *core.Store) (result string, store_was_updated bool)
  Valid() bool
}

type CommandParameters struct {
	Operation string `json:"operation"`
	CollectionName  string `json:"collection_name"`
	Limit int `json:"limit"`
	Predicates []struct{
		Path []string `json:"path"`
		Operator string `json:"operator"`
		Value interface{} `json:"value"`
	} `json:"where"`
	Data map[string]interface{} `json:"data"`
}

func New(cmdData CommandParameters) (Command, error) {
	switch cmdData.Operation {
	case "DEFINE COLLECTION":
		return DefineCollectionCommand{collectionName: cmdData.CollectionName}, nil
	case "REMOVE COLLECTION":
		return RemoveCollectionCommand{collectionName: cmdData.CollectionName}, nil
	case "REMOVE ITEMS":
		return RemoveItemCommand{collectionName: cmdData.CollectionName, limit: cmdData.Limit, predicates: comp.BuildPredicateCollection()}, nil
	case "STORE":
		return InsertCommand{collectionName: cmdData.CollectionName, data: cmdData.Data}, nil
	case "FETCH":
		return FetchCommand{collectionName: cmdData.CollectionName, limit: cmdData.Limit, predicates: extractPredicateCollection(cmdData)}, nil
	default:
		return nil, errors.New("Error, do not understand input")
	}
}

func extractPredicateCollection(cmdParameters CommandParameters) comp.PredicateCollection {
  predicates := comp.BuildPredicateCollection()

	for _, wp := range cmdParameters.Predicates {
		predicates.AddPredicate(comp.Predicate{
			Path: common.DataPath{RealPath: wp.Path},
			Operator: wp.Operator,
			Value: wp.Value,
		})
  }

  return predicates
}


