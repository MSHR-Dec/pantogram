package service

import "encoding/json"

func ToMap(input interface{}) (output map[string]string) {
	marshaled, _ := json.Marshal(input)
	json.Unmarshal(marshaled, &output)

	return
}
