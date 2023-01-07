package config

import (
	"encoding/json"
	"fmt"
	"github.com/qingconglaixueit/wechatbot/pkg/logger"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// Configuration 项目配置
type Configuration struct {
	// gpt apikey
	ApiKey string `json:"api_key"`
	// 自动通过好友
	AutoPass bool `json:"auto_pass"`
	// 会话超时时间
	SessionTimeout time.Duration `json:"session_timeout"`
	// GPT请求最大字符数
	MaxTokens uint `json:"max_tokens"`
	// GPT模型
	Model string `json:"model"`
	// 热度
	Temperature float64 `json:"temperature"`
	// 回复前缀
	ReplyPrefix string `json:"reply_prefix"`
	// 清空会话口令
	SessionClearToken string `json:"session_clear_token"`
	// 每天工作起始时间
	StartTime int `json:"start_time"`
	// 每天工作结束时间
	EndTime int `json:"end_time"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 给配置赋默认值
		config = &Configuration{
			AutoPass:          false,
			SessionTimeout:    60,
			MaxTokens:         512,
			Model:             "text-davinci-003",
			Temperature:       0.9,
			SessionClearToken: "下一个问题",
			StartTime:         9,
			EndTime:           21,
		}

		// 判断配置文件是否存在，存在直接JSON读取
		_, err := os.Stat("config.json")
		if err == nil {
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
		}
		// 有环境变量使用环境变量
		ApiKey := os.Getenv("APIKEY")
		AutoPass := os.Getenv("AUTO_PASS")
		SessionTimeout := os.Getenv("SESSION_TIMEOUT")
		Model := os.Getenv("MODEL")
		MaxTokens := os.Getenv("MAX_TOKENS")
		Temperature := os.Getenv("TEMPREATURE")
		ReplyPrefix := os.Getenv("REPLY_PREFIX")
		SessionClearToken := os.Getenv("SESSION_CLEAR_TOKEN")
		StartTime := os.Getenv("START_TIME")
		EndTime := os.Getenv("END_TIME")
		if ApiKey != "" {
			config.ApiKey = ApiKey
		}
		if AutoPass == "true" {
			config.AutoPass = true
		}
		if SessionTimeout != "" {
			duration, err := time.ParseDuration(SessionTimeout)
			if err != nil {
				logger.Danger(fmt.Sprintf("config session timeout err: %v ,get is %v", err, SessionTimeout))
				return
			}
			config.SessionTimeout = duration
		}
		if Model != "" {
			config.Model = Model
		}
		if MaxTokens != "" {
			max, err := strconv.Atoi(MaxTokens)
			if err != nil {
				logger.Danger(fmt.Sprintf("config MaxTokens err: %v ,get is %v", err, MaxTokens))
				return
			}
			config.MaxTokens = uint(max)
		}
		if Temperature != "" {
			temp, err := strconv.ParseFloat(Temperature, 64)
			if err != nil {
				logger.Danger(fmt.Sprintf("config Temperature err: %v ,get is %v", err, Temperature))
				return
			}
			config.Temperature = temp
		}
		if ReplyPrefix != "" {
			config.ReplyPrefix = ReplyPrefix
		}
		if SessionClearToken != "" {
			config.SessionClearToken = SessionClearToken
		}

		if StartTime != "" {
			sTime, err := strconv.Atoi(StartTime)
			if err != nil {
				logger.Warning("StartTime=%s  strconv.Atoi error:%+v",StartTime,err)
			}else{
				config.StartTime = sTime
			}
		}

		if EndTime != "" {
			eTime, err := strconv.Atoi(EndTime)
			if err != nil {
				logger.Warning("EndTime=%s  strconv.Atoi error:%+v",EndTime,err)
			}else{
				config.EndTime = eTime
			}
		}

	})
	if config.ApiKey == "" {
		logger.Danger("config err: api key required")
	}

	return config
}
