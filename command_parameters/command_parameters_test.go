package command_parameters

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDefaultValues(t *testing.T) {
  var params = New()

  assert.Equal(t, params.Limit, -1)
  assert.Equal(t, params.Error, "")
}

func TestWithError(t *testing.T) {
  var params = New()
  params.Error = "some error"

  assert.True(t, params.HasError())
}
