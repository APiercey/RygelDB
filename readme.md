# RygelDB
[RygelDB](https://github.com/APiercey/RygelDB) is a NoSQL document store using commands to store and query documents!

<img src="https://raw.githubusercontent.com/APiercey/RygelDB/main/sparky.png">

> What the Yotz?! - Dominar Rygel the XVI

- [Usage](#usage)
  * [Storing and Querying Data](#storing-and-querying-data)
  * [Defining Collections](#defining-collections)
  * [Removing Collections](#removing-collections)
  * [Storing Data](#storing-data)
  * [Querying data](#querying-data)
  * [Remove data](#remove-data)
  * [Update data](#update-data)
  * [Where Predicates](#where-predicates)

## Usage
Run `go run .` or `go build .`, as you like.

```
$ go run .

Starting tcp server on localhost:8080
```

You can connect using a socket once it has started like so:
```
nc localhost 8080
```

### Storing and Querying Data
Rygel stores JSON document items in collections.

All commands can be sent over socket. To test them individually, you may pipe time:
```bash
echo '{ "operation": "DEFINE COLLECTION", "collection_name": "test_collection" }' | nc localhost 8080
```

### Defining Collections
```json
{ 
  "operation": "DEFINE COLLECTION",
  "collection_name": "collection_name"
}
```
Creates a new collection where document items may be stored.

### Removing Collections
```json
{ 
  "operation": "REMOVE COLLECTION",
  "collection_name": "collection_name"
}
```
Removes a collection from the datbase. This removes all stored JSON document items within it.

### Storing Data
```json
{ 
  "operation": "STORE",
  "collection_name": "collection_name",
  "data": {"data": "structure of document"}
}
```
Stores a document item in a collection. Data can be any valid JSON structure.

### Querying data
```json
{ 
  "operation": "FETCH",
  "collection_name": "collection_name",
  "limit": 1,
  "where": [
    { "path": ["foo"], "operator": "=", "value": "new value" }
  ]
}
```
Queries for data in a collection. The `limit` parameter is optional and if provided, will stop looking for items after it has reached it's limit.

Rygel can query for data based on certain criteria, named where predicates. Each where predicate is defined used:
- `path` an array of keys, with each key being a step into a nested structure. The value for the last key is used in the comparison.
- `operator` how the value is compared.
- `value` what the value is expected to be to match.

Given the following data:
```javascript
{
  "operation": "DEFINE COLLECTION",
  "collection_name": "fruits"
}
{ 
  "operation": "STORE",
  "collection_name": "fruits",
  "data": {"key":"apple","color":"red"}
}
{ 
  "operation": "STORE",
  "collection_name": "fruits",
  "data": {"key":"orange","color":"orange"}
}
```

Querying for a single document would look like:
```javascript
{ 
  "operation": "FETCH",
  "collection_name": "fruits",
  "limit": 1
}
```
> [{"color":"orange","key":"orange"}]

Querying for all documents that meet a criteria:
```javascript
{ 
  "operation": "FETCH",
  "collection_name": "fruits",
  "where": [
    { "path": ["color"], "operator": "=", "value": "red" },
  ]
}
```
> [{"color":"red","key":"apple"}]

It's possible to query based on deep properties and multiple WHERE clauses:
```javascript
{ 
  "operation": "STORE",
  "collection_name": "fruits",
  "data": {
    "key": "dragonfruit",
    "color": "red",
    "properties": {
      "spikes": "many",
      "internal_color": "white"
    }
  }
}
{ 
  "operation": "FETCH",
  "collection_name": "fruits",
  "where": [
    { "path": ["color"], "operator": "=", "value": "red" },
    { "path": ["properties", "internal_color"], "operator": "=", "value": "white" },
  ]
}
```
> [{"color":"red","key":"dragonfruit","properties":{"internal_color":"white","spikes":"many"}}]

### Remove data
```javascript
{ 
  "operation": "REMOVE ITEMS",
  "collection_name": "test_collection",
  "limit": 1,
  "where": [
    { "path": ["amount"], "operator": ">", "value": 1000 },
  ]
}
```
Removes JSON document items from a collection. Limit is optional.

### Update data
```json
{ 
  "operation": "UPDATE ITEM",
  "collection_name": "test_collection",
  "limit": 1,
  "where": [
    { "path": ["foo"], "operator": "=", "value": "bar" }
  ],
  "data": {"foo": "YOTZ!"}
}
```

Updates JSON items in a collection. Works with where predicates and limit is optional.

Note, this command _replaces_ the data structure and is not an deep merge of the two data objects.

### Where Predicates
Where predicates can compare strings, integers, and booleans.

They provide a number of different operators useful for these value types:
- `=` asserts if the value on path is equal to the expected value.
- `!=` refutes if value on path is equal to the expected value.
- `>` compares if value on path is greater than expected value.
- `>=` compares if value on path is greater than or equal to the expected value.
- `<` compares if value on path is less than the expected value.
- `<=` compares if value on path is less than or equal to the expected value.

---

> May your afterlife be almost as pleasant as mine. - Dominar Rygel the XVI

