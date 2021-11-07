package command_result

type CommandResult struct {
  storeUpdated bool
  data string
}

func (r CommandResult) StoreWasUpdated() bool {
  return r.storeUpdated
}

func (r CommandResult) CommandResult() string {
  return r.data
}

func New(storeUpdated bool, data string) CommandResult {
  return CommandResult{
    storeUpdated: storeUpdated,
    data: data,
  }
}
