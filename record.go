package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path"
)

const (
	DirPerm        = 0775
	FilePerm       = 0664
	RecordsBaseDir = ".recorder"
)

type SerializedRecord struct {
	Method string
	Path   string
	Data   string
}

type Record struct {
	method string
	url    *url.URL
	body   io.ReadCloser

	baseDir string
	lazyErr error
}

func NewRecord(method string, url *url.URL, body io.ReadCloser) *Record {
	wd, err := os.Getwd()

	return &Record{
		method: method,
		url:    url,
		body:   body,

		baseDir: path.Join(wd, RecordsBaseDir),
		lazyErr: err,
	}
}

// NOTE: this is probably pretty racy
func (r *Record) Save() error {
	if err := r.lazyErr; err != nil {
		return err
	}

	dir := path.Join(r.baseDir, r.method, r.url.Path)
	if err := os.MkdirAll(dir, DirPerm); err != nil {
		return err
	}

	records, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	index := len(records)
	name := fmt.Sprintf("%d.json", index)
	filename := path.Join(dir, name)

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
