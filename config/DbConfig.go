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

func New() (dbConfig *DBConfig) {
	// 获取全局环境变量
	contextConfig := NewContextConfig()

	fpath, err := filepath.Abs("resources/db-" + contextConfig.Env + ".yaml")
	checkErr(err)
	configFile, err := ioutil.ReadFile(fpath)
	checkErr(err)
	err = yaml.Unmarshal(configFile, &dbConfig)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
