package config

import (
	"os"
	"strconv"

	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

var keyMap map[KeyName]string

type Config struct {
	Server Server
}

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// init config from config file
func init() {
	var config Config
	yamlFile, err := os.ReadFile("./.gin-client-go.yaml")
	if err != nil {
		klog.Fatal(err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		klog.Fatal(err)
		return
	}
	keyMap = make(map[KeyName]string)
	keyMap[ServerName] = config.Server.Name
	keyMap[ServerHost] = config.Server.Host
	keyMap[ServerPort] = config.Server.Port
}

// get string value
func GetString(keyName KeyName) string {
	return keyMap[keyName]
}

// get int value
func GetInt(keyName KeyName) int {
	intStr := keyMap[keyName]
	if intStr == "" {
		klog.Fatal("GetInt not read config: ", keyName)
		return -1
	}
	// 将变量转为 int 类型
	v, err := strconv.Atoi(intStr)
	if intStr == "" {
		klog.Fatal(err)
		return -1
	}
	return v
}
