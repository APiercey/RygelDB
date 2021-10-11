package input_parser

import (
	"fmt"
	"testing"
	"reflect"

	"github.com/stretchr/testify/assert"
)

func TestParsingDefaultValues(t *testing.T) {
  var result = Parse(`{}`);

  assert.Equal(t, result.Limit, -1)
  assert.Equal(t, result.Error, "")
}

func TestParsingOperation(t *testing.T) {
  var result = Parse(`{
    "operation": "DEFINE COLLECTION"
  }`);

  assert.Equal(t, result.Operation, "DEFINE COLLECTION")
  assert.False(t, result.HasError())
}

func TestParseCollectionName(t *testing.T) {
  var result = Parse(`{
    "collection_name": "test_collection"
  }`);

  assert.Equal(t, result.CollectionName, "test_collection")
  assert.False(t, result.HasError())
}

func TestParseData(t *testing.T) {
  // NOTE: Integers for some reason do not match
  var result = Parse(`{
    "data": {
      "string": "bar",
      "float": 45.67,
      "boolean": true,
      "nested": {
        "data": "works"
      }
    }
  }`);

  eq := reflect.DeepEqual(result.Data, map[string]interface{}{
    "string": "bar",
    "float": 45.67,
    "boolean": true,
    "nested": map[string]interface{}{
      "data": "works",
    },
  })

  assert.True(t, eq)
  assert.False(t, result.HasError())
}

func TestParsing(t *testing.T) {
  var result = Parse(`{
    "where": [
      { "path": ["match"], "operator": "=", "value": "me" }
    ]
  }`);

  fmt.Printf("%#v\n", result)

  assert.NotEmpty(t, result.WhereClauses)
  assert.Equal(t, result.WhereClauses[0].Path, []string{"match"})
  assert.Equal(t, result.WhereClauses[0].Operator, "=")
  assert.Equal(t, result.WhereClauses[0].Value, "me")
  assert.False(t, result.HasError())
}

func TestParsingErrors(t *testing.T) {
  var result = Parse(`{"incorrect json":}`);

  assert.NotEqual(t, result.Error, "")
  assert.True(t, result.HasError())
}
