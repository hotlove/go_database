package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type ContextConfig struct {
	Env  string // 系统环境 dev test prod
	Name string // 系统名字根目录
}

func NewContextConfig() (context *ContextConfig) {
	fpath, err := filepath.Abs("env.yaml")
	fmt.Println(fpath)
	checkErr(err)
	configFile, err := ioutil.ReadFile(fpath)
	checkErr(err)
	err = yaml.Unmarshal(configFile, &context)
	return
}
