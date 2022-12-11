package handlers

import (
	"fmt"
	"github.com/869413421/wechatbot/config"
	"github.com/869413421/wechatbot/pkg/logger"
	"github.com/eatmoreapple/openwechat"
	"github.com/patrickmn/go-cache"
	"github.com/skip2/go-qrcode"
	"log"
	"runtime"
	"strings"
	"time"
)

var c = cache.New(config.LoadConfig().SessionTimeout, time.Minute*5)

// MessageHandlerInterface 消息处理接口
type MessageHandlerInterface interface {
	handle() error
	ReplyText() error
}

// QrCodeCallBack 登录扫码回调，
func QrCodeCallBack(uuid string) {
	if runtime.GOOS == "windows" {
		// 运行在Windows系统上
		openwechat.PrintlnQrcodeUrl(uuid)
	} else {
		log.Println("login in linux")
		url := "https://login.weixin.qq.com/l/" + uuid
		log.Printf("如果二维码无法扫描，请缩小控制台尺寸，或更换命令行工具，缩小二维码像素")
		q, _ := qrcode.New(url, qrcode.High)
		fmt.Println(q.ToSmallString(true))
	}
}

// Handler 全局处理入口
func Handler(msg *openwechat.Message) {
	defer func() {
		err := recover()
		if err != nil {
			logger.Warning(fmt.Sprintf("handler recover error: %v", err))
		}
	}()

	// 清空会话
	if strings.Contains(msg.Content, config.LoadConfig().SessionClearToken) {
		// 获取口令消息处理器
		handler, err := NewTokenMessageHandler(msg)
		if err != nil {
			logger.Warning(fmt.Sprintf("init token message handler error: %s", err))
		}

		// 获取口令消息处理器
		err = handler.handle()
		if err != nil {
			logger.Warning(fmt.Sprintf("handle token message error: %s", err))
		}
		return
	}

	// 处理群消息
	if msg.IsSendByGroup() {
		// 获取用户消息处理器
		handler, err := NewGroupMessageHandler(msg)
		if err != nil {
			logger.Warning(fmt.Sprintf("init group message handler error: %s", err))
			return
		}

		// 处理用户消息
		err = handler.handle()
		if err != nil {
			logger.Warning(fmt.Sprintf("handle group message error: %s", err))
		}
		return
	}

	// 好友申请
	if msg.IsFriendAdd() {
		if config.LoadConfig().AutoPass {
			_, err := msg.Agree("")
			if err != nil {
				logger.Warning(fmt.Sprintf("add friend agree error : %v", err))
				return
			}
		}
	}

	// 私聊
	// 获取用户消息处理器
	handler, err := NewUserMessageHandler(msg)
	if err != nil {
		logger.Warning(fmt.Sprintf("init user message handler error: %s", err))
	}

	// 处理用户消息
	err = handler.handle()
	if err != nil {
		logger.Warning(fmt.Sprintf("handle user message error: %s", err))
	}
	return
}
