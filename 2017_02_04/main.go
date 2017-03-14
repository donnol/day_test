package main

import (
	"encoding/json"
	"fmt"
)

func decodeJsonError() {
	var result interface{}
	data := []byte(`{"test":"json"`)
	// data := []byte(`{"test":"json"}`)
	err := json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func main() {
	decodeJsonError()
}
