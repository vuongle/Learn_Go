package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	// convert json string to struct
	var ninja Ninja
	jsonString := `{"name": "james", "age": 18, "score": 100}`
	err := json.Unmarshal([]byte(jsonString), &ninja)
	if err != nil {
		panic(err)
	}

	fmt.Println(ninja)

	// convert struct to json string
	jsonTo, err := json.Marshal(ninja)
	fmt.Println(string(jsonTo))
}
