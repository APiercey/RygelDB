package commands

import (
	"example.com/rygel/common"
	"example.com/rygel/core"
	cp "example.com/rygel/command_parameters"
	comp "example.com/rygel/comparisons"
)

type Command interface {
  Execute(store *core.Store) (result string, store_was_updated bool)
  Valid() bool
}

// TODO: Should this be a part of the CommandParameters struct?
func extractPredicateCollection(params cp.CommandParameters) comp.PredicateCollection {
  predicates := comp.BuildPredicateCollection()

	for _, wp := range params.WhereClauses {
		predicates.AddPredicate(comp.Predicate{
			Path: common.DataPath{RealPath: wp.Path},
			Operator: wp.Operator,
			Value: wp.Value,
		})
  }

  return predicates
}

func New(params cp.CommandParameters) Command {
	if params.Error != "" {
		return noopErrorCommand{err: params.Error}
	}

	switch params.Operation {
	case "DEFINE COLLECTION":
		return defineCollectionCommand{collectionName: params.CollectionName}
	case "REMOVE COLLECTION":
		return removeCollectionCommand{collectionName: params.CollectionName}
	case "REMOVE ITEMS":
		return removeItemCommand{collectionName: params.CollectionName, limit: params.Limit, predicates: comp.BuildPredicateCollection()}
	case "STORE":
		return insertCommand{collectionName: params.CollectionName, data: params.Data}
	case "UPDATE ITEM":
		return updateItemCommand{collectionName: params.CollectionName, limit: params.Limit, predicates: comp.BuildPredicateCollection(), data: params.Data}
	case "FETCH":
		return fetchCommand{collectionName: params.CollectionName, limit: params.Limit, predicates: extractPredicateCollection(params)}
	default:
		return noopErrorCommand{err: "Command was not understood. Nothing has been executed."}
	}
}

