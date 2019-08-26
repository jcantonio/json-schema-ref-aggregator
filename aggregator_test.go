package main

import (
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
