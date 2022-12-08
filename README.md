# wechatbot
> 最近chatGPT异常火爆，本项目可以将个人微信化身GPT机器人，
> 项目基于[openwechat](https://github.com/eatmoreapple/openwechat) 开发。

[![Release](https://img.shields.io/github/v/release/869413421/wechatbot.svg?style=flat-square)](https://github.com/869413421/wechatbot/releases/tag/v1.0.1)
![Github stars](https://img.shields.io/github/stars/869413421/wechatbot.svg)
![Forks](https://img.shields.io/github/forks/869413421/wechatbot.svg?style=flat-square)

### 目前实现了以下功能
 * 提问增加上下文，更接近官网效果 
 * 机器人群聊@回复
 * 机器人私聊回复
 * 好友添加自动通过
 
# 使用前提
> * ~~目前只支持在windows上运行因为需要弹窗扫码登录微信，后续会支持linux~~   已支持
> * 有openai账号，并且创建好api_key，注册事项可以参考[此文章](https://juejin.cn/post/7173447848292253704) 。
> * 微信必须实名认证。

# 注意事项
> * 项目仅供娱乐，滥用可能有微信封禁的风险，请勿用于商业用途。
> * 请注意收发敏感信息，本项目不做信息过滤。

# 快速开始
> 非技术人员请直接下载release中的[压缩包](https://github.com/869413421/wechatbot/releases/tag/v1.1.1) ，解压运行。
````
# 获取项目
git clone https://github.com/869413421/wechatbot.git

# 进入项目目录
cd wechatbot

# 复制配置文件
copy config.dev.json config.json

# 启动项目
go run main.go
````

# 配置文件说明
````
{
"api_key": "your api key",
"auto_pass": true,
"session_timeout": 60
}

api_key：openai api_key
auto_pass:是否自动通过好友添加
session_timeout：会话超时时间，默认60秒，单位秒，在会话时间内所有发送给机器人的信息会作为上下文。
````

# 使用示例
### 向机器人发送`我要问下一个问题`，清空会话信息。
### 私聊
<img width="300px" src="https://raw.githubusercontent.com/869413421/study/master/static/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20221208153022.jpg"/>

### 群聊@回复
<img width="300px" src="https://raw.githubusercontent.com/869413421/study/master/static/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20221208153015.jpg"/>

### 添加微信（备注: wechabot）进群交流

**如果二维码图片没显示出来，请添加微信号 huangyanming681925**

<img width="210px"  src="https://raw.githubusercontent.com/869413421/study/master/static/qr.png" align="left">

