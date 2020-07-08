package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var LISEN_HOST = ""
var STORAGE_IP = ""
var STATIC_URL = ""
var STATIC_DIR = ""
var DB_ADDRESS = ""

type conf struct {
	Lisen_host string `yaml:"lisen_host"`
	Storage_ip string `yaml:"storage_ip"`
	Static_url string `yaml:"static_url"`
	Static_dir string `yaml:"static_dir"`
	Db_address string `yaml:"db_address"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("./config/file_config.yaml")
	if err != nil {
		fmt.Println("config init ... err:", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println("config init ... err:", err)
	}

	return c
}

func Init() {
	var c conf
	c.getConf()

	LISEN_HOST = c.Lisen_host
	STORAGE_IP = c.Storage_ip
	STATIC_URL = c.Static_url
	STATIC_DIR = c.Static_dir
	DB_ADDRESS = c.Db_address

}
