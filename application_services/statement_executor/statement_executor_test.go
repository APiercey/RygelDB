package statement_executor

import (
  "testing"
  "os"

  cx "rygel/services/command_executor"
  sr "rygel/core/store_repo"
  "rygel/infrastructure/ledger"
  "rygel/context"
)

func setupContext() context.Context {
  return context.Context{SelectedStore: "test"}
}

func useTempStoreRepo(name string, f func(sr.StoreRepo)) {
  tmpDir := "./tmp-test-store"

  if err := os.MkdirAll(tmpDir, 0755); err != nil {
    panic(err)
  }

  defer func() {
    if err := os.RemoveAll(tmpDir); err != nil {
      panic(err)
    }
  }()

  storeRepo := sr.InitializeFromDir(tmpDir)
  storeRepo.Create(name)

  f(storeRepo)
}

func setupService(storeRepo sr.StoreRepo) StatementExecutor {
  commandExecutor := cx.SyncCommandExecutor{ }
  ledger := ledger.InMemoryLedger{}

  return StatementExecutor{
    CommandExecutor: &commandExecutor,
    Ledger: &ledger,
    StoreRepo: storeRepo,
  }
}

func TestDefineCollection(t *testing.T) {
  useTempStoreRepo("test", func(storeRepo sr.StoreRepo) {
    serv := setupService(storeRepo)
    ctx := setupContext()
    expected := "OK"

    result := serv.Execute(ctx, `
{ "operation": "DEFINE COLLECTION", "collection_name": "test_collection" }
    `)

    if result != expected {
      t.Log("Expected: ", expected)
      t.Log("Received: ", result)
      t.Fail()
    }
  })
}

func TestRemoveCollection(t *testing.T) {
  useTempStoreRepo("test", func(storeRepo sr.StoreRepo) {
    serv := setupService(storeRepo)
    ctx := setupContext()

    expected := "OK"

    serv.Execute(ctx, `{ "operation": "DEFINE COLLECTION", "collection_name": "test_collection" } `)
    result := serv.Execute(ctx, `{ "operation": "REMOVE COLLECTION", "collection_name": "test_collection" } `)

    if result != expected {
      t.Log("Expected: ", expected)
      t.Log("Received: ", result)
      t.Fail()
    }
  })
}

func TestInsertItem(t *testing.T) {
  useTempStoreRepo("test", func(storeRepo sr.StoreRepo) {
    serv := setupService(storeRepo)
    ctx := setupContext()

    expected := "OK"

    serv.Execute(ctx, `{ "operation": "DEFINE COLLECTION", "collection_name": "test_collection" } `)
    result := serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
})
}

func TestRemoveSingleItem(t *testing.T) {
  useTempStoreRepo("test", func(storeRepo sr.StoreRepo) {
    serv := setupService(storeRepo)
    ctx := setupContext()

    serv.Execute(ctx, `{
"operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar", "key": "test_item"}
  }`)
  result := serv.Execute(ctx, `{ 
"operation": "REMOVE ITEMS",
    "collection_name": "test_collection",
    "limit": 1
  }`)

  if result != "Removed 1 items" { t.Error("Command does not work", result) }
})
}

func TestRemovedItemsNotRetreivable(t *testing.T) {
  useTempStoreRepo("test", func(storeRepo sr.StoreRepo) {
    serv := setupService(storeRepo)
    ctx := setupContext()

    expected := `[{"foo":"bar"}]`

    serv.Execute(ctx, `{
"operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  serv.Execute(ctx, `{ 
"operation": "REMOVE ITEMS",
    "collection_name": "test_collection",
    "limit": 1
  }`)
  result := serv.Execute(ctx, `{ 
"operation": "FETCH",
    "collection_name": "test_collection",
    "limit": 1
  }`)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
})
}

func TestFetchSingleItem(t *testing.T) {
  useTempStoreRepo("test", func(storeRepo sr.StoreRepo) {
    serv := setupService(storeRepo)
    ctx := setupContext()

    expected := `[{"foo":"bar"}]`

    serv.Execute(ctx, `{
"operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"pow": "blam"}
  }`)
  result := serv.Execute(ctx, `{ 
"operation": "FETCH",
    "collection_name": "test_collection",
    "limit": 1
  }`)

  if result != expected {
    t.Log("Expected: ", expected)
    t.Log("Received: ", result)
    t.Fail()
  }
})
}

func TestFetchItemWithEqualsWhereClause(t *testing.T) {
  useTempStoreRepo("test", func(storeRepo sr.StoreRepo) {
    serv := setupService(storeRepo)
    ctx := setupContext()

    expected := `[{"foo":"bar","some":{"nested":{"path":"Hello World"}}}]`

    serv.Execute(ctx, `{
"operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"should not": "fetch"}
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar", "some": {"nested": {"path": "Hello World"}}}
  }`)
  result := serv.Execute(ctx, `{ 
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
})
}

func TestFetchItemWithGreaterThanWhereClause(t *testing.T) {
  useTempStoreRepo("test", func(storeRepo sr.StoreRepo) {
    serv := setupService(storeRepo)
    ctx := setupContext()

    expected := `[{"foo":"bar","my_num":23}]`

    serv.Execute(ctx, `{
"operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"should not": "fetch"}
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar", "my_num": 23}
  }`)
  result := serv.Execute(ctx, `{ 
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
})
}

func TestFetchItemsMultipleWhereClause(t *testing.T) {
  useTempStoreRepo("test", func(storeRepo sr.StoreRepo) {
    serv := setupService(storeRepo)
    ctx := setupContext()

    expected := `[{"health":50,"match":"me"},{"health":99,"match":"me"}]`

    serv.Execute(ctx, `{
"operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"health": 50, "match":"me"}
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"health": 99, "match":"me"}
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "choo", "match":"me"}
  }`)
  result := serv.Execute(ctx, `{ 
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
})
}

func TestUpdateItem(t *testing.T) {
  useTempStoreRepo("test", func(storeRepo sr.StoreRepo) {
    serv := setupService(storeRepo)
    ctx := setupContext()

    serv.Execute(ctx, `{
"operation": "DEFINE COLLECTION",
    "collection_name": "test_collection"
  }`)
  serv.Execute(ctx, `{ 
"operation": "STORE",
    "collection_name": "test_collection",
    "data": {"foo": "bar"}
  }`)
  updateResult := serv.Execute(ctx, `{ 
"operation": "UPDATE ITEM",
    "collection_name": "test_collection",
    "data": {"foo": "new value"}
  }`)

  if updateResult != "Updated 1 items" { t.Error("Command does not work", updateResult) }

  fetchResult := serv.Execute(ctx, `{ 
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
})

}
