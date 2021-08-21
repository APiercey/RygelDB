package main

import (
	"testing"

  "fmt"
  "encoding/json"
	"example.com/rygel/store"
)

func buildItem(key string, dataString string) store.Item {
  var data map[string]interface{};
  json.Unmarshal([]byte(dataString), &data)
  item, _ := store.BuildItem(key, data)

  return item
}

func setupStore() store.Store {
  testStore := store.BuildStore("/tmp")

  return testStore
}

func TestDefineCollection(t *testing.T) {
  testStore := setupStore()

  result, _ := ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)

  if result != "OK" { t.Error("Command does not work") }
}

func TestRemoveCollection(t *testing.T) {
  testStore := setupStore()

  ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)
  result, _ := ExecuteStatementAgainstStore(&testStore, `REMOVE COLLECTION test_collection`)

  if result != "OK" { t.Error("Command does not work") }
}

func TestInsertItem(t *testing.T) {
  testStore := setupStore()

  ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)
  result, _ := ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection xyz {"foo":"bar"}`)

  if result != "OK" { t.Error("Command does not work") }
}

func TestRemoveItem(t *testing.T) {
  testStore := setupStore()

  ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)
  ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection xyz {"foo":"bar"}`)
  result, _ := ExecuteStatementAgainstStore(&testStore, "REMOVE ITEM xyz IN test_collection")

  if result != "OK" { t.Error("Command does not work") }
}

func TestLookupItem(t *testing.T) {
  testStore := setupStore()

  ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)
  ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection test_item {"key":"test_item","foo":"bar"}`)
  result, _ := ExecuteStatementAgainstStore(&testStore, "LOOKUP test_item IN test_collection")

  fmt.Println(result)

  if result != `{"foo":"bar","key":"test_item"}` { t.Error("Command does not work") }
}

func TestFetchItemsSingleWhereClause(t *testing.T) {
  testStore := setupStore()

  ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)
  ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection test_item {"key":"test_item","foo":"bar"}`)
  result, _ := ExecuteStatementAgainstStore(&testStore, "FETCH 1 FROM test_collection WHERE foo IS bar")

  if result != `[{"foo":"bar","key":"test_item"}]` { t.Error("Command does not work") }
}

func TestFetchItemsMultipleWhereClause(t *testing.T) {
  testStore := setupStore()

  ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)
  ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection test_item_one {"key":"test_item_one","foo":"bar"}`)
  ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection test_item_two {"key":"test_item_two","foo":"bar"}`)
  result, _ := ExecuteStatementAgainstStore(&testStore, "FETCH 1 FROM test_collection WHERE foo IS bar AND key IS test_item_one")

  if result != `[{"foo":"bar","key":"test_item_one"}]` { t.Error("Command does not work") }
}

func TestFetchAllItems(t *testing.T) {
  testStore := setupStore()

  ExecuteStatementAgainstStore(&testStore, `DEFINE COLLECTION test_collection`)
  ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection test_item_one {"key":"test_item_one","foo":"bar"}`)
  ExecuteStatementAgainstStore(&testStore, `STORE INTO test_collection test_item_two {"key":"test_item_two","foo":"bar"}`)
  result, _ := ExecuteStatementAgainstStore(&testStore, "FETCH all FROM test_collection WHERE foo IS bar")

  if result != `[{"foo":"bar","key":"test_item_one"},{"foo":"bar","key":"test_item_two"}]` { t.Error("Command does not work") }
}
