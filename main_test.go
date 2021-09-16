package main

import (
	"testing"
	"example.com/rygel/store"
)

func setupStore() store.Store {
  testStore := store.BuildStore()

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

// func TestRemoveItemsSingleWhereClause(t *testing.T) {
//   testStore := setupStore()

//   ExecuteStatementAgainstStore(&testStore, `{
//     "operation": "DEFINE COLLECTION",
//     "collection_name": "test_collection"
//   }`)
//   ExecuteStatementAgainstStore(&testStore, `{ 
//     "operation": "STORE",
//     "collection_name": "test_collection",
//     "data": {"foo": "bar"}
//   }`)
//   ExecuteStatementAgainstStore(&testStore, `{ 
//     "operation": "STORE",
//     "collection_name": "test_collection",
//     "data": {"foo": "bar", "key": "test_item"}
//   }`)
//   result, _ := ExecuteStatementAgainstStore(&testStore, "REMOVE 1 ITEM FROM test_collection WHERE foo IS bar")
//   result, _ := ExecuteStatementAgainstStore(&testStore, `{ 
//     "operation": "REMOVE ITEMS",
//     "collection_name": "test_collection",
//     "limit": 1,
//     "where": []
//   }`)

//   if result != "Removed 1 items" { t.Error("Command does not work", result) }
// }

// func TestRemoveAllItemsWhereClause(t *testing.T) {
//   testStore := setupStore()

//   ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)
//   ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection {"foo":"bar"}`)
//   ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection {"key":"test_item","foo":"bar"}`)
//   result, _ := ExecuteStatementAgainstStore(&testStore, "REMOVE all ITEM FROM test_collection WHERE foo IS bar")

//   if result != "Removed 2 items" { t.Error("Command does not work", result) }
// }

// func TestFetchItemsSingleWhereClause(t *testing.T) {
//   testStore := setupStore()
//   expected := `[{"foo":"bar","key":"test_item"}]`

//   ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)
//   ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection {"key":"test_item","foo":"bar"}`)
//   result, _ := ExecuteStatementAgainstStore(&testStore, "FETCH 1 FROM test_collection WHERE foo IS bar")

//   if result != expected {
//     t.Log("Expected: ", expected)
//     t.Log("Received: ", result)
//     t.Fail()
//   }
// }
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

// func TestFetchAllItems(t *testing.T) {
//   testStore := setupStore()
//   expected := `[{"foo":"bar","key":"test_item_one"},{"foo":"bar","key":"test_item_two"}]`

//   ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)
//   ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection {"key":"test_item_one","foo":"bar"}`)
//   ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection {"key":"test_item_two","foo":"bar"}`)
//   result, _ := ExecuteStatementAgainstStore(&testStore, "FETCH all FROM test_collection WHERE foo IS bar")

//   if result != expected {
//     t.Log("Expected: ", expected)
//     t.Log("Received: ", result)
//     t.Fail()
//   }
// }
