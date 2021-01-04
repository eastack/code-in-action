package jsondb

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var dbpath = "jsondb/data"

func init() {
	Clear()
}

func Clear() {
	if _, err := os.Stat(dbpath); !os.IsNotExist(err) {
		err := os.RemoveAll(dbpath)
		if err != nil {
			panic(err)
		}
	}
}

func Write(resource string, id string, v interface{}) {
	// 如果文件夹不存在创建
	if _, err := os.Stat(filepath.Join(dbpath, resource)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Join(dbpath, resource), 0755)
		if err != nil {
			panic(err)
		}
	}

	jsonString, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filepath.Join(dbpath, resource, id), jsonString, 0644)
	if err != nil {
		panic(err)
	}
}

func Delete(resource, id string) {
	if _, err := os.Stat(filepath.Join(dbpath, resource, id)); !os.IsNotExist(err) {
		err := os.Remove(filepath.Join(dbpath, resource, id))
		if err != nil {
			panic(err)
		}
	}
}

func Read(resource string, id string, v interface{}) bool {
	if _, err := os.Stat(filepath.Join(dbpath, resource, id)); os.IsNotExist(err) {

		return false
	} else {
		file, err := ioutil.ReadFile(filepath.Join(dbpath, resource, id))
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(file, v)
		if err != nil {
			panic(err)
		}

		return true
	}
}
