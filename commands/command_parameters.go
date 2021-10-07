package commands

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
}

func NewCommandParameters() CommandParameters {
	return CommandParameters{
		Limit: -1,
	}
}
