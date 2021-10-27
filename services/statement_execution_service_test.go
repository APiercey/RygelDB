package services

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
  serv := StatementExecutionService{}
  expected := "OK"

  result, _ := serv.Execute(&testStore, `
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
  serv := StatementExecutionService{}
  expected := "OK"

  serv.Execute(&testStore, `{ "operation": "DEFINE COLLECTION", "collection_name": "test_collection" } `)
  result, _ := serv.Execute(&testStore, `{ "operation": "REMOVE COLLECTION", "collection_name": "test_collection" } `)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
}

func TestInsertItem(t *testing.T) {
  testStore := setupStore()
  serv := StatementExecutionService{}
  expected := "OK"

  serv.Execute(&testStore, `{ "operation": "DEFINE COLLECTION", "collection_name": "test_collection" } `)
  result, _ := serv.Execute(&testStore, `{ 
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
  serv := StatementExecutionService{}

  serv.Execute(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar", "key": "test_item"}
  }`)
  result, _ := serv.Execute(&testStore, `{ 
    "operation": "REMOVE ITEMS",
    "collection_name": "test_collection",
    "limit": 1
  }`)

  if result != "Removed 1 items" { t.Error("Command does not work", result) }
}

func TestRemovedItemsNotRetreivable(t *testing.T) {
  testStore := setupStore()
  serv := StatementExecutionService{}
  expected := `[{"foo":"bar"}]`

  serv.Execute(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "REMOVE ITEMS",
    "collection_name": "test_collection",
    "limit": 1
  }`)
  result, _ := serv.Execute(&testStore, `{ 
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
  serv := StatementExecutionService{}
  expected := `[{"foo":"bar"}]`

  serv.Execute(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"pow": "blam"}
  }`)
  result, _ := serv.Execute(&testStore, `{ 
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
  serv := StatementExecutionService{}
  expected := `[{"foo":"bar","some":{"nested":{"path":"Hello World"}}}]`

  serv.Execute(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"should not": "fetch"}
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar", "some": {"nested": {"path": "Hello World"}}}
  }`)
  result, _ := serv.Execute(&testStore, `{ 
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
  serv := StatementExecutionService{}
  expected := `[{"foo":"bar","my_num":23}]`

  serv.Execute(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"should not": "fetch"}
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar", "my_num": 23}
  }`)
  result, _ := serv.Execute(&testStore, `{ 
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
  serv := StatementExecutionService{}
  expected := `[{"health":50,"match":"me"},{"health":99,"match":"me"}]`

  serv.Execute(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"health": 50, "match":"me"}
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"health": 99, "match":"me"}
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "choo", "match":"me"}
  }`)
  result, _ := serv.Execute(&testStore, `{ 
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
  serv := StatementExecutionService{}

  serv.Execute(&testStore, `{
    "operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(&testStore, `{ 
    "operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  updateResult, _ := serv.Execute(&testStore, `{ 
    "operation": "UPDATE ITEM",
    "collection_name": "test_collection",
    "data": {"foo": "new value"}
  }`)

  if updateResult != "Updated 1 items" { t.Error("Command does not work", updateResult) }

  fetchResult, _ := serv.Execute(&testStore, `{ 
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
