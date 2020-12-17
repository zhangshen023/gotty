package base

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

const configPath = "../config/config.yaml"

var GottyConfig *Config
var once sync.Once

func init() {
	once.Do(func() {
		GottyConfig = &Config{}
	})
	err := ParseConfigYaml(GottyConfig, configPath)
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败:%v", err))
	}
	fmt.Println("配置文件读取成功")
}

type Config struct {
	Server Server `json:"server" yaml:"server"`
	Client Client `json:"client" yaml:"client"`
}

type Server struct {
	Concurrency             uint   `json:"concurrency" yaml:"concurrency"`
	SessionNumPerConnection uint   `json:"sessionNumPerConnection" yaml:"sessionNumPerConnection"`
	Port                    uint   `json:"port" yaml:"port"`
	ListenerName            string `json:"listenerName" yaml:"listenerName"`
}
type Client struct {
}

func ParseConfigYaml(target *Config, filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(content, target); err != nil {
		return err
	}
	return nil
}