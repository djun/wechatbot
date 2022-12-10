package handlers

import (
	"errors"
	"fmt"
	"github.com/869413421/wechatbot/config"
	"github.com/869413421/wechatbot/gtp"
	"github.com/869413421/wechatbot/pkg/logger"
	"github.com/eatmoreapple/openwechat"
	"strings"
)

var _ MessageHandlerInterface = (*UserMessageHandler)(nil)

// UserMessageHandler 私聊消息处理
type UserMessageHandler struct {
}

// handle 处理消息
func (g *UserMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewUserMessageHandler 创建私聊处理器
func NewUserMessageHandler() MessageHandlerInterface {
	return &UserMessageHandler{}
}

// ReplyText 发送文本消息到群
func (g *UserMessageHandler) ReplyText(msg *openwechat.Message) error {
	// 接收私聊消息
	sender, err := msg.Sender()
	logger.Info(fmt.Sprintf("Received User %v Text Msg : %v", sender.NickName, msg.Content))
	if UserService.ClearUserSessionContext(sender.ID(), msg.Content) {
		_, err = msg.ReplyText("上下文已经清空了，你可以问下一个问题啦。")
		if err != nil {
			return errors.New(fmt.Sprintf("response user error: %v", err))
		}
		return nil
	}

	// 获取上下文，向GPT发起请求
	requestText := strings.TrimSpace(msg.Content)
	requestText = strings.Trim(msg.Content, "\n")
	requestText = UserService.GetUserSessionContext(sender.ID()) + requestText
	reply, err := gtp.Completions(requestText)
	if err != nil {
		errMsg := fmt.Sprintf("gtp request error: %v", err)
		_, err = msg.ReplyText(errMsg)
		if err != nil {
			return errors.New(fmt.Sprintf("response user error: %v ", err))
		}
		return err
	}
	if reply == "" {
		return nil
	}

	// 设置上下文，回复用户
	UserService.SetUserSessionContext(sender.ID(), requestText, reply)
	_, err = msg.ReplyText(buildUserReply(reply))
	if err != nil {
		return errors.New(fmt.Sprintf("response user error: %v ", err))
	}
	return err
}

// buildUserReply 构建用户回复
func buildUserReply(reply string) string {
	reply = strings.Trim(strings.Trim(reply, "？"), "\n")
	if reply == "" {
		return "请求得不到任何有意义的回复，请具体提出问题。"
	}
	reply = config.LoadConfig().ReplyPrefix + "\n" + reply
	reply = strings.Trim(reply, "\n")
	return reply
}
