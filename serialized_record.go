package main

import (
	"encoding/json"
	"os"
)

type SerializedRecord struct {
	Method string
	Path   string
	Data   string
}

func LoadRecord(method string, path string, index int) (*SerializedRecord, error) {
	var err error

	if index == LastIndex {
		index, err = LastRecordIndex(method, path)
	}

	if err != nil {
		return nil, err
	}

	filename, err := RecordFileName(method, path, index)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	record := &SerializedRecord{}
	if err := decoder.Decode(record); err != nil {
		return nil, err
	}

	return record, nil
}
