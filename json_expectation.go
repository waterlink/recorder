package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	EPS = 1e-9
)

type JSONExpectation struct {
	data     string
	jsonPath string
}

func NewJSONExpectation(data string, jsonPath string) Expectation {
	return &JSONExpectation{
		data:     data,
		jsonPath: jsonPath,
	}
}

func (e *JSONExpectation) Verify(r *SerializedRecord) error {
	parts := strings.Split(e.jsonPath, "/")

	var err error

	value, err := e.parseTraverse([]byte(r.Data), map[string]interface{}{}, parts)
	if err != nil {
		value, err = e.parseTraverse([]byte(r.Data), []interface{}{}, parts)
	}

	if err != nil {
		return fmt.Errorf("Unable to parse JSON nor as an Object, nor as an Array: %s", err)
	}

	if asNum, ok := value.(float64); ok {
		expected, err := strconv.ParseFloat(e.data, 64)
		if err != nil {
			return err
		}

		if math.Abs(expected-asNum) >= EPS {
			return fmt.Errorf("JSON expectation failed:\n\tPath:     %s\n\tExpected: %f\n\tActual:   %f\n\n", e.jsonPath, expected, asNum)
		}

		return nil
	}

	if asStr, ok := value.(string); ok {
		expected := e.data

		if expected != asStr {
			return fmt.Errorf("JSON expectation failed:\n\tPath:     %s\n\tExpected: %s\n\tActual:   %s\n\n", e.jsonPath, expected, asStr)
		}

		return nil
	}

	return fmt.Errorf("Unexpected value '%#v' at JSON path %s", value, e.jsonPath)
}

func (e *JSONExpectation) parseTraverse(bytes []byte, data interface{}, parts []string) (interface{}, error) {
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Printf("Unable to unmarshal: %s", err)
		return nil, err
	}

	return e.traverse(data, parts)
}

func (e *JSONExpectation) traverse(data interface{}, parts []string) (interface{}, error) {
	if len(parts) == 0 {
		return data, nil
	}

	key := parts[0]

	idx, err := strconv.Atoi(key)
	if err != nil {
		dataAsMap, ok := data.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Not an Object at suffix %v", parts)
		}

		nextData, ok := dataAsMap[key]
		if !ok {
			return nil, fmt.Errorf("Key '%s' is not found at suffix %v", key, parts)
		}

		return e.traverse(nextData, parts[1:])
	}

	dataAsSlice, ok := data.([]interface{})
	if !ok {
		log.Printf("%#v", data)
		return nil, fmt.Errorf("Not an Array at suffix %v", parts)
	}

	if len(dataAsSlice) <= idx {
		return nil, fmt.Errorf("Index '%d' is out of bounds at suffix %v", idx, parts)
	}

	return e.traverse(dataAsSlice[idx], parts[1:])
}
