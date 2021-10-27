package result

type Result struct {
  storeUpdated bool
  data string
}

func (r Result) StoreWasUpdated() bool {
  return r.storeUpdated
}

func (r Result) CommandResult() string {
  return r.data
}

func New(storeUpdated bool, data string) Result {
  return Result{
    storeUpdated: storeUpdated,
    data: data,
  }
}
