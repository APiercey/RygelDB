package command_parameters
import (
	"rygel/common"
	comp "rygel/comparisons"
)

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
	RawStatement string
}

func (params CommandParameters) HasError() bool {
	return params.Error != ""
}

func (params CommandParameters) ExtractPredicateCollection() comp.PredicateCollection {
  predicates := comp.BuildPredicateCollection()

	for _, wp := range params.WhereClauses {
		predicates.AddPredicate(comp.Predicate{
			Path: common.DataPath{RealPath: wp.Path},
			Operator: wp.Operator,
			Value: wp.Value,
		})
  }

  return predicates
}

func New() CommandParameters {
	return CommandParameters{
		Limit: -1,
		Error: "",
	}
}
