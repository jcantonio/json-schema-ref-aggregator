package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	jsonSchema "github.com/jcantonio/json-schema-ref-aggregator/json-schema"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) != 3 {
		fmt.Println("Requires two args: path to json schema file and path to output file.")
		os.Exit(2)
	}
	jsonMap, err := jsonSchema.GetSchemaWithAggregatedReferences(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	jsonBytes, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = ioutil.WriteFile(os.Args[2], jsonBytes, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
