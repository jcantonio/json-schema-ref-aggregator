package main

import (
	"fmt"
	"regexp"
)

/*
DeepSearchParent returns parent element of an attribute
*/
func DeepSearchParent(attributeName string, jsonMap map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	searchParent("", attributeName, jsonMap, result)
	return result
}

func searchParent(parentPath, attributeName string, jsonMap map[string]interface{}, result map[string]interface{}) {

	for k, v := range jsonMap {
		if k == attributeName {
			result[parentPath] = jsonMap
		}
		switch vv := v.(type) {
		case map[string]interface{}:
			var path string
			if len(parentPath) > 0 {
				path = fmt.Sprintf("%s.%s", parentPath, k)
			} else {
				path = k
			}
			searchParent(path, attributeName, vv, result)
		case []interface{}:
			// TODO
			// for i, u := range vv {
			//
			// }
		}
	}
}

/*
DeepValidate verifies attribute names
*/
func DeepValidate(jsonMap map[string]interface{}) error {
	return validate("", jsonMap)
}

var validAttributeName = regexp.MustCompile(`^[a-zA-Z0-9_]*$`)

func validate(parentPath string, jsonMap map[string]interface{}) error {
	for k, v := range jsonMap {
		if !validAttributeName.MatchString(k) {
			return fmt.Errorf("Invalid attribute attributeName: %s", k)
		}
		switch vv := v.(type) {
		case map[string]interface{}:
			var path string
			if len(parentPath) > 0 {
				path = fmt.Sprintf("%s.%s", parentPath, k)
			} else {
				path = k
			}
			err := validate(path, vv)
			if err != nil {
				return err
			}
		case []interface{}:
			var path string
			if len(parentPath) > 0 {
				path = fmt.Sprintf("%s.%s", parentPath, k)
			} else {
				path = k
			}
			err := validateArray(path, vv)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateArray(parentPath string, jsonArray []interface{}) error {
	for _, v := range jsonArray {
		switch vv := v.(type) {
		case map[string]interface{}:
			err := validate(parentPath, vv)
			if err != nil {
				return err
			}
		case []interface{}:
			err := validateArray(parentPath, vv)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
