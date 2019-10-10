package jsonSchema

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestGetSchema(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{ // TODO: Add test cases.
		{
			name: "",
			args: args{
				filePath: "weight.json",
			},
			want: map[string]interface{}{
				"attrs": map[string]interface{}{},
			},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetSchema(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSchemaWithAggregatedReferences(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{ // TODO: Add test cases.
		{
			name: "",
			args: args{
				filePath: "weight.json",
			},
			want: map[string]interface{}{
				"attrs": map[string]interface{}{},
			},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetSchemaWithAggregatedReferences(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSchemaWithAggregatedReferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeepSearchParent(t *testing.T) {
	data, err := ioutil.ReadFile("../weight-full.json")
	if err != nil {
		t.Error(err)
		return
	}
	jsonMap := make(map[string]interface{})
	json.Unmarshal(data, &jsonMap)
	res := DeepSearchParent("description", jsonMap)
	println(res)
}

func TestDeepValidate(t *testing.T) {
	data, err := ioutil.ReadFile("../weight-full.json")
	if err != nil {
		t.Error(err)
	}
	jsonMap := make(map[string]interface{})
	json.Unmarshal([]byte(data), &jsonMap)
	println(DeepValidate(jsonMap))
}
