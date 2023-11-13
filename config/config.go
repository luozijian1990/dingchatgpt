package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type AppConfig struct {
	DBIp          string `yaml:"db_ip"`
	DBPort        int    `yaml:"db_port"`
	DBUser        string `yaml:"db_user"`
	DBPasswd      string `yaml:"db_password"`
	DBName        string `yaml:"db_name"`
	DingAppKey    string `yaml:"ding_app_key"`
	DingAppSecret string `yaml:"ding_app_secret"`
	GptBaseUrl    string `yaml:"gpt_baseurl"`
	GptToken      string `yaml:"gpt_token"`
}

var (
	AppConfigFile string
)

func init()  {

}

func InitFlat() {
	flag.StringVar(&AppConfigFile, "conf-file", "D:\\GoWork\\src\\gostudy\\new-chat-ding\\config.yaml", "config-file")
	//GetConfig()
	//AppConfig()
}

func GetConfig() *AppConfig {
	var config *AppConfig
	yamlFile, err := ioutil.ReadFile(AppConfigFile)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &config)

	return config

}
