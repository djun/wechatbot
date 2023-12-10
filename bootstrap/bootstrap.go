package bootstrap

import (
	"github.com/eatmoreapple/openwechat"
	"io"
	"log"
	"wechatbot/handlers"
)

func Run() {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = handlers.Handler
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")

	defer func(reloadStorage io.ReadWriteCloser) {
		err := reloadStorage.Close()
		if err != nil {
			log.Printf("storage.json close error: %v \n", err)
		}
	}(reloadStorage)

	// 执行热登录
	err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Printf("login error: %v \n", err)
			return
		}
	}
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
