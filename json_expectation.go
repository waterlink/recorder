package main

import "errors"

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
	return errors.New("JSONPath based expectation is not yet implemented")
}
