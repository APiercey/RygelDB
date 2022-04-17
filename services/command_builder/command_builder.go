package command_builder

import (
  "rygel/commands"
	cs "rygel/core/store"
	cr "rygel/core/store_repo"
	cp "rygel/services/command_builder/command_parameters"
)

type CommandBuilder struct {
	StoreRepo cr.StoreRepo
}

func (cb CommandBuilder) Build(store *cs.Store, params cp.CommandParameters) commands.Command {
	if params.Error != "" {
		return commands.NoopErrorCommand{Err: params.Error}
	}

	switch params.Operation {
	case "DEFINE STORE":
		return commands.DefineStoreCommand{
			StoreName: params.StoreName,
			StoreRepo: cb.StoreRepo,
		}
	case "DEFINE COLLECTION":
		return commands.DefineCollectionCommand{
			Store: store,
			CollectionName: params.CollectionName,
		}
	case "REMOVE COLLECTION":
		return commands.RemoveCollectionCommand{
			Store: store,
			CollectionName: params.CollectionName,
		}
	case "REMOVE ITEMS":
		return commands.RemoveItemCommand{
			Store: store,
			CollectionName: params.CollectionName,
			Limit: params.Limit,
			Predicates: params.ExtractPredicateCollection(),
		}
	case "STORE":
		return commands.InsertCommand{
			Store: store,
			CollectionName: params.CollectionName,
			Data: params.Data,
		}
	case "UPDATE ITEM":
		return commands.UpdateItemCommand{
			Store: store,
			CollectionName: params.CollectionName,
			Limit: params.Limit,
			Predicates: params.ExtractPredicateCollection(),
			Data: params.Data,
		}
	case "FETCH":
		return commands.FetchCommand{
			Store: store,
			CollectionName: params.CollectionName,
			Limit: params.Limit,
			Predicates: params.ExtractPredicateCollection(),
		}
	default:
		return commands.NoopErrorCommand{
			Err: "Command was not understood. Nothing has been executed.",
		}
	}
}

