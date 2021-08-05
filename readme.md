# RygelDB
Rygel is a NoSQL document store using commands to store and query documents!

<img src="https://raw.githubusercontent.com/APiercey/RygelDB/main/sparky.png">

> What the Yotz?! - Dominar Rygel the XVI

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
Rygel is collection based and uses commands to store and query data.

All commands can be sent over socket. To test them individually, you may pipe time:
```
echo "DEFINE COLLECTION my_new_collection" | nc localhost 8080
```

#### Defining Collections
`DEFINE COLLECTION collection_name` will create a new collection where document items may be stored.

#### Storing Data
`STORE INTO collection_name key {"data": "structure of document"}` will store a document item.

#### Lookup of direct data
`LOOKUP key IN collection_name` retrieves a document by key

#### Querying data
`FETCH [all | 1, ...n] FROM collection_name [WHERE path.of.document.properties IS value AND ...n]` queries data using 0 or many WHERE clauses and enforces either _all_ or a limit.

Given the following data:
```
DEFINE COLLECTION fruits
STORE INTO fruits apple {"key":"apple","color":"red"}
STORE INTO fruits orange {"key":"orange","color":"orange"}
```

Querying for a single document would look like:
```
FETCH 1 FROM fruits
```
> [{"color":"orange","key":"orange"}]

Querying for all documents that meet a criteria:
```
FETCH all FROM fruits WHERE color IS red
```
> [{"color":"red","key":"apple"}]

It's possible to query based on deep properties and multiple WHERE clauses:
```
STORE INTO fruits orange {"key":"dragondruit","color":"red","properties":{"spikes":"many","internal_color":"white"}}
FETCH all FROM fruits WHERE color IS red AND properties.internal_color IS white
```
> [{"color":"red","key":"dragondruit","properties":{"internal_color":"white","spikes":"many"}}]

#### Remove data
`REMOVE [COLLECTION | ITEM] collection_name [key]` removes either a collection or a document item in a colleciton. Key is mandatory when removing a document item.

