package command_parameters

type WhereClause struct {
	Path []string `json:"path"`
	Operator string `json:"operator"`
	Value interface{} `json:"value"`
}

type CommandParameters struct {
	Operation string `json:"operation"`
	CollectionName  string `json:"collection_name"`
	Limit int `json:"limit"`
	WhereClauses []struct{
		Path []string `json:"path"`
		Operator string `json:"operator"`
		Value interface{} `json:"value"`
	} `json:"where"`
	Data map[string]interface{} `json:"data"`
	Error string
}

func (cp CommandParameters) HasError() bool {
	return cp.Error != ""
}

func New() CommandParameters {
	return CommandParameters{
		Limit: -1,
		Error: "",
	}
}
