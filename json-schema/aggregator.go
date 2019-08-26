package jsonSchema

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tidwall/sjson"
)

const REF_PREF_TYPE_FILE = "file:///"

/*
GetSchemaWithAggregatedReferences returns schema with the aggregated schema references
*/
func GetSchemaWithAggregatedReferences(filePath string) (map[string]interface{}, error) {
	schema, err := GetSchema(filePath)
	if err != nil {
		return schema, err
	}
	return GetDataWithAggregatedReferences(schema, 0)
}

/*
GetSchema  returns schema without the aggregated schema references
*/
func GetSchema(filePath string) (map[string]interface{}, error) {
	var result map[string]interface{}

	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return result, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal([]byte(byteValue), &result)

	return result, err
}

/*
GetDataWithAggregatedReferences returns schema with the aggregated schema references
*/
func GetDataWithAggregatedReferences(jsonMap map[string]interface{}, level int64) (map[string]interface{}, error) {
	aggregatedValue, _ := json.Marshal(jsonMap)
	aggregatedValueStr := string(aggregatedValue)
	res := DeepSearchParent("$ref", jsonMap)

	// avoid infinite recursive loop
	if level < 10 {
		for k, v := range res {
			refJSONMap := v.(map[string]interface{})

			uri, ok := (refJSONMap["$ref"]).(string)
			if !ok {
				break
			}
			if strings.HasPrefix(uri, REF_PREF_TYPE_FILE) {
				filePath := strings.TrimPrefix(uri, REF_PREF_TYPE_FILE)
				newValueMap, err := GetSchema(filePath)
				if err != nil {
					return nil, err
				}
				delete(newValueMap, "$schema")

				attrs := refJSONMap["attrs"]
				if attrs != nil {
					newValueMap["attrs"] = attrs
				}

				// recursive
				childJSONMap, err := GetDataWithAggregatedReferences(newValueMap, level+1)
				if err != nil {
					return nil, err
				}
				aggregatedValueStr, _ = sjson.Set(aggregatedValueStr, k, childJSONMap)
			}
		}
	}

	json.Unmarshal([]byte(aggregatedValueStr), &jsonMap)

	return jsonMap, nil
}
