package commands

import (
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
	WhereClauses []struct{
		Path []string `json:"path"`
		Operator string `json:"operator"`
		Value interface{} `json:"value"`
	} `json:"where"`
	Data map[string]interface{} `json:"data"`
}

func New(cp CommandParameters) Command {
	switch cp.Operation {
	case "DEFINE COLLECTION":
		return defineCollectionCommand{collectionName: cp.CollectionName}
	case "REMOVE COLLECTION":
		return removeCollectionCommand{collectionName: cp.CollectionName}
	case "REMOVE ITEMS":
		return removeItemCommand{collectionName: cp.CollectionName, limit: cp.Limit, predicates: comp.BuildPredicateCollection()}
	case "STORE":
		return insertCommand{collectionName: cp.CollectionName, data: cp.Data}
	case "FETCH":
		return fetchCommand{collectionName: cp.CollectionName, limit: cp.Limit, predicates: extractPredicateCollection(cp)}
	default:
		return noopErrorCommand{err: "Command was not understood. Nothing has been executed."}
	}
}

func extractPredicateCollection(cp CommandParameters) comp.PredicateCollection {
  predicates := comp.BuildPredicateCollection()

	for _, wp := range cp.WhereClauses {
		predicates.AddPredicate(comp.Predicate{
			Path: common.DataPath{RealPath: wp.Path},
			Operator: wp.Operator,
			Value: wp.Value,
		})
  }

  return predicates
}


