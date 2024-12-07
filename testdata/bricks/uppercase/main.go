package main

import (
	"encoding/json"
	"strings"
)

//export BeforeInsert
func BeforeInsert(in []byte) []byte {
	var data map[string]any
	err := json.Unmarshal(in, &data)
	if err != nil {
		return []byte{}
	}
	name, ok := data["name"].(string)
	if !ok {
		return []byte{}
	}
	data["name"] = strings.ToUpper(name)
	out, err := json.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return out
}
