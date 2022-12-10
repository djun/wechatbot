package handlers

import (
	"errors"
	"fmt"
	"github.com/869413421/wechatbot/gtp"
	"github.com/869413421/wechatbot/pkg/logger"
	"github.com/eatmoreapple/openwechat"
	"strings"
)

var _ MessageHandlerInterface = (*GroupMessageHandler)(nil)

// GroupMessageHandler 群消息处理
type GroupMessageHandler struct {
}

// handle 处理消息
func (g *GroupMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewGroupMessageHandler 创建群消息处理器
func NewGroupMessageHandler() MessageHandlerInterface {
	return &GroupMessageHandler{}
}

// ReplyText 发送文本消息到群
func (g *GroupMessageHandler) ReplyText(msg *openwechat.Message) error {
	// 接收群消息
	sender, err := msg.Sender()
	group := openwechat.Group{User: sender}
	logger.Info(fmt.Sprintf("Received Group %v Text Msg : %v", group.NickName, msg.Content))

	// 不是@的不处理
	if !msg.IsAt() {
		return nil
	}

	// 获取@我的用户
	groupSender, err := msg.SenderInGroup()
	if err != nil {
		return errors.New(fmt.Sprintf("get sender in group error :%v ", err))
	}
	atText := "@" + groupSender.NickName + " "

	if UserService.ClearUserSessionContext(groupSender.ID(), msg.Content) {
		_, err = msg.ReplyText(atText + "上下文已经清空了，你可以问下一个问题啦。")
		if err != nil {
			return errors.New(fmt.Sprintf("response user error: %v", err))
		}
		return nil
	}

	// 替换掉@文本，设置会话上下文，然后向GPT发起请求。
	requestText := buildRequestText(sender, msg)
	if requestText == "" {
		return nil
	}
	reply, err := gtp.Completions(requestText)
	if err != nil {
		// 将GPT请求失败信息输出给用户，省得整天来问又不知道日志在哪里。
		errMsg := fmt.Sprintf("gtp request error: %v", err)
		_, err = msg.ReplyText(errMsg)
		if err != nil {
			return errors.New(fmt.Sprintf("response group error: %v ", err))
		}
		return err
	}
	if reply == "" {
		return nil
	}

	// 设置上下文
	UserService.SetUserSessionContext(sender.ID(), requestText, reply)
	replyText := atText + buildGroupReply(reply)
	_, err = msg.ReplyText(replyText)
	if err != nil {
		return errors.New(fmt.Sprintf("response group error: %v ", err))
	}
	return err
}

// buildUserReply 构建用户回复
func buildGroupReply(reply string) string {
	// 回复@我的用户
	reply = strings.Trim(strings.Trim(reply, "？"), "\n")
	if reply == "" {
		return "请求得不到任何有意义的回复，请具体提出问题。"
	}
	reply = strings.Trim(reply, "\n")
	return reply
}

// buildRequestText 构建请求GPT的文本，替换掉机器人名称，然后检查是否有上下文，如果有拼接上
func buildRequestText(sender *openwechat.User, msg *openwechat.Message) string {
	replaceText := "@" + sender.Self.NickName
	requestText := strings.TrimSpace(strings.ReplaceAll(msg.Content, replaceText, ""))
	if requestText == "" {
		return ""
	}
	requestText = UserService.GetUserSessionContext(sender.ID()) + requestText
	return requestText
}
