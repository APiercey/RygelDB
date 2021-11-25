package command_builder

import (
	"testing"
	"github.com/stretchr/testify/assert"

	cp "rygel/commands/command_parameters"
	"rygel/commands"
	"rygel/core"
)

func setup() CommandBuilder {
	return CommandBuilder{}
}

func TestReturnsDefineCollection(t *testing.T) {
	cb := setup()
	store := core.Store{}
	params := cp.New()
	params.Operation = "DEFINE COLLECTION"
	cmd := cb.Build(&store, params)

  assert.IsType(t, commands.DefineCollectionCommand{}, cmd)
}

func TestReturnsRemoveCollection(t *testing.T) {
	cb := setup()
	store := core.Store{}
	params := cp.New()
	params.Operation = "REMOVE COLLECTION"
	cmd := cb.Build(&store, params)

  assert.IsType(t, commands.RemoveCollectionCommand{}, cmd)
}

func TestReturnsRemoveItems(t *testing.T) {
	cb := setup()
	store := core.Store{}
	params := cp.New()
	params.Operation = "REMOVE ITEMS"
	cmd := cb.Build(&store, params)

  assert.IsType(t, commands.RemoveItemCommand{}, cmd)
}

func TestReturnsInsert(t *testing.T) {
	cb := setup()
	store := core.Store{}
	params := cp.New()
	params.Operation = "STORE"
	cmd := cb.Build(&store, params)

  assert.IsType(t, commands.InsertCommand{}, cmd)
}

func TestReturnsFetch(t *testing.T) {
	cb := setup()
	store := core.Store{}
	params := cp.New()
	params.Operation = "FETCH"
	cmd := cb.Build(&store, params)

  assert.IsType(t, commands.FetchCommand{}, cmd)
}

func TestReturnsUpdate(t *testing.T) {
	cb := setup()
	store := core.Store{}
	params := cp.New()
	params.Operation = "UPDATE ITEM"
	cmd := cb.Build(&store, params)

  assert.IsType(t, commands.UpdateItemCommand{}, cmd)
}

func TestReturnsNoopError(t *testing.T) {
	cb := setup()
	store := core.Store{}
	params := cp.New()
	params.Operation = "DOES NOT EXIST"
	cmd := cb.Build(&store, params)

  assert.IsType(t, commands.NoopErrorCommand{}, cmd)
}

func TestWithError(t *testing.T) {
	cb := setup()
	store := core.Store{}
	params := cp.New()
	params.Operation = "DEFINE COLLECTION"
	params.Error = "Something wrong with input"
	cmd := cb.Build(&store, params)

  assert.IsType(t, commands.NoopErrorCommand{}, cmd)
}

