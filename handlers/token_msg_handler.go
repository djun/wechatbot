package handlers

import (
	"github.com/869413421/wechatbot/pkg/logger"
	"github.com/869413421/wechatbot/service"
	"github.com/eatmoreapple/openwechat"
)

var _ MessageHandlerInterface = (*TokenMessageHandler)(nil)

// TokenMessageHandler 口令消息处理器
type TokenMessageHandler struct {
	// 接收到消息
	msg *openwechat.Message
	// 发送的用户
	sender *openwechat.User
	// 实现的用户业务
	service service.UserServiceInterface
}

// NewTokenMessageHandler 口令消息处理器
func NewTokenMessageHandler(msg *openwechat.Message) (MessageHandlerInterface, error) {
	sender, err := msg.Sender()
	if err != nil {
		return nil, err
	}
	if msg.IsComeFromGroup() {
		sender, err = msg.SenderInGroup()
	}
	userService := service.NewUserService(c, sender)
	handler := &TokenMessageHandler{
		msg:     msg,
		sender:  sender,
		service: userService,
	}

	return handler, nil
}

// handle 处理口令
func (t *TokenMessageHandler) handle() error {
	return t.ReplyText()
}

// ReplyText 回复清空口令
func (t *TokenMessageHandler) ReplyText() error {
	logger.Info("user clear token")
	t.service.ClearUserSessionContext()
	var err error
	if t.msg.IsComeFromGroup() {
		atText := "@" + t.sender.NickName + "上下文已经清空，请问下一个问题。"
		_, err = t.msg.ReplyText(atText)
	} else {
		_, err = t.msg.ReplyText("上下文已经清空，请问下一个问题。")
	}
	return err
}
