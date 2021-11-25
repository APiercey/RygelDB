package common

import (
	"encoding/json"
)

type Data map[string]interface{}

func EncodeData(data Data) []byte  {
	out, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return out
}

func DecodeData(rawData []byte) Data {
  var data Data

	err := json.Unmarshal(rawData, &data)

	if err != nil {
		panic(err)
	}

	return data
}
