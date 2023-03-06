package service

import (
	"github.com/qingconglaixueit/wechatbot/config"
	"github.com/eatmoreapple/openwechat"
	"github.com/patrickmn/go-cache"
	"time"
)

// UserServiceInterface 用户业务接口
type UserServiceInterface interface {
	GetUserSessionContext() string
	SetUserSessionContext(question, reply string)
	ClearUserSessionContext()
}

var _ UserServiceInterface = (*UserService)(nil)

// UserService 用戶业务
type UserService struct {
	// 缓存
	cache *cache.Cache
	// 用户
	user *openwechat.User
}

// NewUserService 创建新的业务层
func NewUserService(cache *cache.Cache, user *openwechat.User) UserServiceInterface {
	return &UserService{
		cache: cache,
		user:  user,
	}
}

// ClearUserSessionContext 清空GTP上下文，接收文本中包含`我要问下一个问题`，并且Unicode 字符数量不超过20就清空
func (s *UserService) ClearUserSessionContext() {
	s.cache.Delete(s.user.ID())
}

// GetUserSessionContext 获取用户会话上下文文本
func (s *UserService) GetUserSessionContext() string {
	// 1.获取上次会话信息，如果没有直接返回空字符串
	sessionContext, ok := s.cache.Get(s.user.ID())
	if !ok {
		return ""
	}

	// 2.如果字符长度超过等于4000，强制清空会话（超过GPT会报错）。
	contextText := sessionContext.(string)
	if len(contextText) >= 4000 {
		s.cache.Delete(s.user.ID())
	}

	// 3.返回上文
	return contextText
}

// SetUserSessionContext 设置用户会话上下文文本，question用户提问内容，GTP回复内容
func (s *UserService) SetUserSessionContext(question, reply string) {
	value := question + "\n" + reply
	s.cache.Set(s.user.ID(), value, time.Second*config.LoadConfig().SessionTimeout)
}
