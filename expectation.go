package main

type Expectation interface {
	Verify(r *SerializedRecord) error
}

func NewExpectation(data string, jsonPath string) Expectation {
	if jsonPath != "" {
		return NewJSONExpectation(data, jsonPath)
	}

	return NewEqExpectation(data)
}
