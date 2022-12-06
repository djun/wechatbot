package handlers

import (
	"github.com/eatmoreapple/openwechat"
)

// MessageHandlerInterface 消息处理接口
type MessageHandlerInterface interface {
	handle(*openwechat.Message) error
	ReplyText(*openwechat.Message) error
}

type HandlerType string

const (
	GroupHandler = "group"
)

var handlers map[HandlerType]MessageHandlerInterface

func init() {
	handlers = make(map[HandlerType]MessageHandlerInterface)
	handlers[GroupHandler] = NewGroupMessageHandler()
}

// Handler 全局处理入口
func Handler(msg *openwechat.Message) {
	//if msg.IsSendBySelf() {
	//	return
	//}
	//sender, err := msg.Sender()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	if msg.IsSendByGroup() {
		handlers[GroupHandler].handle(msg)
		return
	}
}
