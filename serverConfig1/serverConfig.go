package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// serverConfig struct 对应配置文件的架构
type serverConfig struct {
	AppConfig struct {
		Name    string `yaml:"Name"`
		Address string `yaml:"Address"`
		Env     string `yaml:"Env"`
	} `yaml:"AppConfig"`

	Mysql struct {
		Dsn     string `yaml:"Dsn"`
		MaxIdle int    `yaml:"MaxIdle"`
		MaxOpen int    `yaml:"MaxOpen"`
		Name    string `yaml:"Name"`
		Debug   bool   `yaml:"Debug"`
	} `yaml:"Mysql"`
}

func newserverConfig() *serverConfig {
	return &serverConfig{
		AppConfig: struct {
			Name    string `yaml:"Name"`
			Address string `yaml:"Address"`
			Env     string `yaml:"Env"`
		}{
			Name:    "default-name",
			Address: "default-address",
			Env:     "default-env",
		},
		Mysql: struct {
			Dsn     string `yaml:"Dsn"`
			MaxIdle int    `yaml:"MaxIdle"`
			MaxOpen int    `yaml:"MaxOpen"`
			Name    string `yaml:"Name"`
			Debug   bool   `yaml:"Debug"`
		}{
			Dsn:     "default-dsn",
			MaxIdle: 0,
			MaxOpen: 0,
			Name:    "default-name",
			Debug:   false,
		},
	}
}

// readConfig 读取 yaml 配置文件并返回它的数据
func (c *serverConfig) readConfig(filename string) (*serverConfig, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("readConfig failed: %v", err)
		return nil, err
	}

	c = &serverConfig{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *serverConfig) getMysqlDsn() (Dsn string, err error) {
	conf, err := c.readConfig("././deployment/serverConfig1.yml")
	if err != nil {
		log.Fatalf("getMysqlDsn failed: %v", err)
	}
	Dsn = conf.Mysql.Dsn
	return Dsn, err
}

//func main() {
//	c := newserverConfig()
//	dsn, err := c.getMysqlDsn()
//	if err != nil {
//		log.Fatalf("getMysqlDsn failed: %v", err)
//	}
//
//	fmt.Printf("DSN: %v\n", dsn)
//}
