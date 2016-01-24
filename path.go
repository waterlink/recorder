package main

import (
	"fmt"
	"io/ioutil"
	"os"
	p "path"
)

func BaseDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return p.Join(wd, RecordsBaseDir), nil
}

func RecordDir(method string, path string) (string, error) {
	base, err := BaseDir()
	if err != nil {
		return "", err
	}

	dir := p.Join(base, method, path)
	if err := os.MkdirAll(dir, DirPerm); err != nil {
		return "", err
	}

	return dir, nil
}

func NextRecordIndex(method string, path string) (int, error) {
	dir, err := RecordDir(method, path)
	if err != nil {
		return 0, err
	}

	records, err := ioutil.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	return len(records), nil
}

func LastRecordIndex(method string, path string) (int, error) {
	index, err := NextRecordIndex(method, path)
	if err != nil {
		return 0, err
	}

	if index == 0 {
		return 0, fmt.Errorf("No records found for %s %s", method, path)
	}

	return index - 1, nil
}

func RecordFileName(method string, path string, index int) (string, error) {
	dir, err := RecordDir(method, path)
	if err != nil {
		return "", err
	}

	name := fmt.Sprintf("%d.json", index)
	return p.Join(dir, name), nil
}

func NextRecordFileName(method string, path string) (string, error) {
	index, err := NextRecordIndex(method, path)
	if err != nil {
		return "", err
	}

	return RecordFileName(method, path, index)
}
