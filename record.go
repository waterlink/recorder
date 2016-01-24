package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"
	"os"
)

const (
	DirPerm        = 0775
	FilePerm       = 0664
	RecordsBaseDir = ".recorder"
)

type Record struct {
	method string
	url    *url.URL
	body   io.ReadCloser

	baseDir string
	lazyErr error
}

func NewRecord(method string, url *url.URL, body io.ReadCloser) *Record {
	baseDir, err := BaseDir()

	return &Record{
		method: method,
		url:    url,
		body:   body,

		baseDir: baseDir,
		lazyErr: err,
	}
}

// NOTE: this is probably pretty racy
func (r *Record) Save() error {
	if err := r.lazyErr; err != nil {
		return err
	}

	filename, err := NextRecordFileName(r.method, r.url.Path)
	if err != nil {
		return err
	}

	data, err := r.data()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, FilePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)

	s := SerializedRecord{
		Method: r.method,
		Path:   r.url.Path,
		Data:   data,
	}

	if err := encoder.Encode(s); err != nil {
		return err
	}

	return f.Close()
}

func (r *Record) data() (string, error) {
	if r.method == "GET" {
		return r.url.RawQuery, nil
	}

	bytes, err := ioutil.ReadAll(r.body)
	return string(bytes), err
}
