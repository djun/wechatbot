# wechatbot
最近chatGPT异常火爆，想到将其接入到个人微信是件比较有趣的事，所以有了这个项目，项目基于[openwechat](https://github.com/eatmoreapple/openwechat)
开发。
![Release](https://github.com/869413421/wechatbot/releases/tag/v1.0.1)
### 目前实现了以下功能
 * 提问增加上下文，更接近官网效果 
 * 机器人群聊@回复
 * 机器人私聊回复
 * 好友添加自动通过
 
# 注册openai
chatGPT注册可以参考[这里](https://juejin.cn/post/7173447848292253704)
，注册完在控制台右上角点用户头像ViewApiKeys中新增加一个api_key。

# 安装使用
````
# 获取项目
git clone https://github.com/869413421/wechatbot.git

# 进入项目目录
cd wechatbot

# 复制配置文件
copy config.dev.json config.json

# 启动项目
go run main.go

启动前需替换config中的api_key
````