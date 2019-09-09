package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Conf struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
	RPC struct {
		Port     string `yaml:"port"`
		KeyFile  string `yaml:"keyFile"`
		CertFile string `yaml:"certFile"`
	}
}

func (c *Conf) LoadConf() *Conf {
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/configs/application.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Print("loadConf successful")
	return c
}
