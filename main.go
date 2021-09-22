package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var inputFileName = os.Args[1]
	var mapping = make(map[string]string)

	data, err := os.ReadFile(inputFileName)
	check(err)

	err = json.Unmarshal([]byte(data), &mapping)
	check(err)

	fmt.Println("[")

	for key, value := range mapping {
		fmt.Println(fmt.Sprintf(`{
  name = "%s",
  value = "%s"
},`, key, value))
	}
	fmt.Print("]")
}
