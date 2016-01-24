package main

import "fmt"

type EqExpectation struct {
	data string
}

func NewEqExpectation(data string) Expectation {
	return &EqExpectation{data: data}
}

func (e *EqExpectation) Verify(r *SerializedRecord) error {
	if e.data != r.Data {
		return fmt.Errorf("Expectation failed on record's data:\n\tExpected: %s\n\tActual:   %s\n\n", e.data, r.Data)
	}

	return nil
}
