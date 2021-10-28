package commands

import (
	"testing"
	"github.com/stretchr/testify/assert"

	cp "rygel/command_parameters"
)

func TestReturnsDefineCollection(t *testing.T) {
	params := cp.New()
	params.Operation = "DEFINE COLLECTION"
	cmd := New(params)

  assert.IsType(t, defineCollectionCommand{}, cmd)
}

func TestReturnsRemoveCollection(t *testing.T) {
	params := cp.New()
	params.Operation = "REMOVE COLLECTION"
	cmd := New(params)

  assert.IsType(t, removeCollectionCommand{}, cmd)
}

func TestReturnsRemoveItems(t *testing.T) {
	params := cp.New()
	params.Operation = "REMOVE ITEMS"
	cmd := New(params)

  assert.IsType(t, removeItemCommand{}, cmd)
}

func TestReturnsInsert(t *testing.T) {
	params := cp.New()
	params.Operation = "STORE"
	cmd := New(params)

  assert.IsType(t, insertCommand{}, cmd)
}

func TestReturnsFetch(t *testing.T) {
	params := cp.New()
	params.Operation = "FETCH"
	cmd := New(params)

  assert.IsType(t, fetchCommand{}, cmd)
}

func TestReturnsUpdate(t *testing.T) {
	params := cp.New()
	params.Operation = "UPDATE ITEM"
	cmd := New(params)

  assert.IsType(t, updateItemCommand{}, cmd)
}

func TestReturnsNoopError(t *testing.T) {
	params := cp.New()
	params.Operation = "DOES NOT EXIST"
	cmd := New(params)

  assert.IsType(t, noopErrorCommand{}, cmd)
}

func TestWithError(t *testing.T) {
	params := cp.New()
	params.Operation = "DEFINE COLLECTION"
	params.Error = "Something wrong with input"
	cmd := New(params)

  assert.IsType(t, noopErrorCommand{}, cmd)
}

