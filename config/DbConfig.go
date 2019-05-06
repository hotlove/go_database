package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func New()(dbConfig * DBConfig)  {
	fpath, err := filepath.Abs("resources/db.yaml")
	checkErr(err)
	configFile, err := ioutil.ReadFile(fpath)
	checkErr(err)
	err = yaml.Unmarshal(configFile, &dbConfig)
	return
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}