package main

import (
	"testing"

	"example.com/rygel/core"
)

func setupStore() core.Store {
  testStore := core.BuildStore()

  return testStore
}

func TestDefineCollection(t *testing.T) {
  testStore := setupStore()
  expected := "OK"

  result, _ := ExecuteStatementAgainstStore(&testStore, `
    { "operation": "DEFINE COLLECTION", "collection_name": "test_collection" }
  `)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
}

func TestRemoveCollection(t *testing.T) {
  testStore := setupStore()
  expected := "OK"

  ExecuteStatementAgainstStore(&testStore, `{ "operation": "DEFINE COLLECTION", "collection_name": "test_collection" } `)
  result, _ := ExecuteStatementAgainstStore(&testStore, `{ "operation": "REMOVE COLLECTION", "collection_name": "test_collection" } `)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
}

func TestInsertItem(t *testing.T) {
  testStore := setupStore()
  expected := "OK"

  ExecuteStatementAgainstStore(&testStore, `{ "operation": "DEFINE COLLECTION", "collection_name": "test_collection" } `)
  result, _ := ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
}

func TestRemoveSingleItem(t *testing.T) {
  testStore := setupStore()

  ExecuteStatementAgainstStore(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar", "key": "test_item"}
  }`)
  result, _ := ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "REMOVE ITEMS",
    "collection_name": "test_collection",
    "limit": 1
  }`)

  if result != "Removed 1 items" { t.Error("Command does not work", result) }
}

func TestRemovedItemsNotRetreivable(t *testing.T) {
  testStore := setupStore()
  expected := `[{"foo":"bar"}]`

  ExecuteStatementAgainstStore(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "REMOVE ITEMS",
    "collection_name": "test_collection",
    "limit": 1
  }`)
  result, _ := ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "FETCH",
    "collection_name": "test_collection",
    "limit": 1
  }`)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
}

func TestFetchSingleItem(t *testing.T) {
  testStore := setupStore()
  expected := `[{"foo":"bar"}]`

  ExecuteStatementAgainstStore(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"pow": "blam"}
  }`)
  result, _ := ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "FETCH",
    "collection_name": "test_collection",
    "limit": 1
  }`)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
}

func TestFetchItemWithEqualsWhereClause(t *testing.T) {
  testStore := setupStore()
  expected := `[{"foo":"bar","some":{"nested":{"path":"Hello World"}}}]`

  ExecuteStatementAgainstStore(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"should not": "fetch"}
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar", "some": {"nested": {"path": "Hello World"}}}
  }`)
  result, _ := ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "FETCH",
    "collection_name": "test_collection",
    "limit": 1,
    "where": [
      { "path": ["some", "nested", "path"], "operator": "=", "value": "Hello World" }
    ]
  }`)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
}

func TestFetchItemWithGreaterThanWhereClause(t *testing.T) {
  testStore := setupStore()
  expected := `[{"foo":"bar","my_num":23}]`

  ExecuteStatementAgainstStore(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"should not": "fetch"}
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar", "my_num": 23}
  }`)
  result, _ := ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "FETCH",
    "collection_name": "test_collection",
    "limit": 1,
    "where": [
      { "path": ["my_num"], "operator": ">", "value": 22 }
    ]
  }`)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
}

func TestFetchItemsMultipleWhereClause(t *testing.T) {
  testStore := setupStore()
  expected := `[{"health":50,"match":"me"},{"health":99,"match":"me"}]`

  ExecuteStatementAgainstStore(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"health": 50, "match":"me"}
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"health": 99, "match":"me"}
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "choo", "match":"me"}
  }`)
  result, _ := ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "FETCH",
    "collection_name": "test_collection",
    "where": [
      { "path": ["match"], "operator": "=", "value": "me" },
      { "path": ["health"], "operator": ">=", "value": 50 }
    ]
  }`)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
}

func TestUpdateItem(t *testing.T) {
  testStore := setupStore()

  ExecuteStatementAgainstStore(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  updateResult, _ := ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "UPDATE ITEM",
    "collection_name": "test_collection",
    "data": {"foo": "new value"}
  }`)

  if updateResult != "Updated 1 items" { t.Error("Command does not work", updateResult) }

  fetchResult, _ := ExecuteStatementAgainstStore(&testStore, `{ 
    "operation": "FETCH",
    "collection_name": "test_collection",
    "limit": 1,
    "where": [
      { "path": ["foo"], "operator": "=", "value": "new value" }
    ]
  }`)

  expected := `[{"foo":"new value"}]`
  if fetchResult != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", fetchResult)
    t.Fail()
  }
}
