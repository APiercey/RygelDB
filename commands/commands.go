package commands

import (
	"rygel/core"
	"rygel/commands/input_parser"
	comp "rygel/comparisons"
)

type Command interface {
  Execute(store *core.Store) (result string, store_was_updated bool)
  Valid() bool
}

func New(input string) Command {
	params := input_parser.Parse(input)

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
		return fetchCommand{collectionName: params.CollectionName, limit: params.Limit, predicates: params.ExtractPredicateCollection()}
	default:
		return noopErrorCommand{err: "Command was not understood. Nothing has been executed."}
	}
}

