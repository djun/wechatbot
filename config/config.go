package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// Configuration 项目配置
type Configuration struct {
	// gtp apikey
	ApiKey string `json:"api_key"`

	//允许回复群列表
	AutoReplyGroups string `json:"auto_reply_groups"`

	//允许回复好友列表
	AutoReplyFriends string `json:"auto_reply_friends"`
	
	// 自动通过好友
	AutoPass bool `json:"auto_pass"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{}
		f, err := os.Open("config.json")
		if err != nil {
			log.Fatalf("open config err: %v", err)
			return
		}
		defer f.Close()
		encoder := json.NewDecoder(f)
		err = encoder.Decode(config)
		if err != nil {
			log.Fatalf("decode config err: %v", err)
			return
		}

		// 如果环境变量有配置，读取环境变量
		ApiKey := os.Getenv("ApiKey")
		AutoPass := os.Getenv("AutoPass")
		AutoReplyGroups := os.Getenv("AutoReplyGroups")
		AutoReplyFriends := os.Getenv("AutoReplyFriends")

		if ApiKey != "" {
			config.ApiKey = ApiKey
		}
		if AutoReplyGroups != "" {
			config.AutoReplyGroups = AutoReplyGroups
		}
		if AutoReplyFriends != "" {
			config.AutoReplyFriends = AutoReplyFriends
		}
		if AutoPass == "true" {
			config.AutoPass = true
		}
	})
	return config
}
