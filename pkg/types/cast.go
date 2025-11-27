package types

import (
	"encoding/json"
	"fmt"
)

func CastTo[T any](data any) (*T, error) {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := new(T)
	if err := json.Unmarshal(jsonbody, &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, err
}
