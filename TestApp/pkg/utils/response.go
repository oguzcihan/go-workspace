package utils

import (
	"encoding/json"
	"fmt"
)

type APIResponse struct {
	isError bool
	code    int
	data    string
}

type NewAPIResponse struct {
	isError bool
	code    int
	data    string
}

func ResponseAPI(_isError bool, _code int, _data string) []byte {
	newData := &APIResponse{
		isError: _isError,
		code:    _code,
		data:    _data,
	}
	response, _ := json.Marshal(newData)
	fmt.Println("marshal", response, newData)

	var unmarshalData *NewAPIResponse
	json.Unmarshal(response, &unmarshalData)
	fmt.Println("unmarshal", []byte(response), unmarshalData)
	
	return response
}
