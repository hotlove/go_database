package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func New()(dbConfig * DBConfig)  {
	configFile, err := ioutil.ReadFile("D:/my_code/go_database/resources/db.yaml")

	checkErr(err)

	err = yaml.Unmarshal(configFile, &dbConfig)
	return
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}