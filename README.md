# wechatbot

> 最近ChatGPT异常火爆，本项目可以将个人微信化身GPT机器人，
> 项目基于[openwechat](https://github.com/eatmoreapple/openwechat) 开发。

> `友链：`[chatgpt-dingtalk](https://github.com/eryajf/chatgpt-dingtalk) 本项目可以将GPT机器人集成到钉钉群聊中。

[![Release](https://img.shields.io/github/v/release/869413421/wechatbot.svg?style=flat-square)](https://github.com/869413421/wechatbot/releases/tag/v1.1.2)
![Github stars](https://img.shields.io/github/stars/869413421/wechatbot.svg)
![Forks](https://img.shields.io/github/forks/869413421/wechatbot.svg?style=flat-square)

### 目前实现了以下功能

* GPT机器人模型热度可配置
* 提问增加上下文
* 指令清空上下文（指令：我要问下一个问题）
* 机器人群聊@回复
* 机器人私聊回复
* 私聊回复前缀设置
* 好友添加自动通过

# 实现机制
目前机器人有两种实现方式
* 逆向功能，扒取官网API，通过抓取cookie获取GPT响应信息，`优点：`效果与官网一致，`缺点：`cookie会过期需要不定时更新。
* 基于openai官网提供的API，`优点`：模型以及各种参数可以自由配置，`缺点：`效果达不到官网智能，且API收费，新账号有18元免费额度。

> 本项目基于第二种方式实现，模型之间具体差异可以参考[官方文档](https://beta.openai.com/docs/models/overview), 详细[参数示例](https://beta.openai.com/examples) 。

# 常见问题
* 如无法登录 login error: write storage.json: bad file descriptor 删除掉storage.json文件重新登录。
* 其他无法登录问题，依然尝试删除掉storage.json文件重新登录。
* ~~机器人无法正常回复，检查ApiKey能否正常使用，控制台日志中有详细错误信息~~ 新版本会机器人会直接输出，因为被问得好烦了。
* linux中二维码无法扫描，缩小命令行功能，让二维码像素尽可能清晰。（无法从代码层面解决）
* 机器人一直答非所问，可能因为上下文累积过多。切换不同问题时，发送指令：我要问下一个问题。会清空上下文

# 使用前提

> * ~~目前只支持在windows上运行因为需要弹窗扫码登录微信，后续会支持linux~~   已支持
> * 有openai账号，并且创建好api_key，注册事项可以参考[此文章](https://juejin.cn/post/7173447848292253704) 。
> * 微信必须实名认证。

# 注意事项

> * 项目仅供娱乐，滥用可能有微信封禁的风险，请勿用于商业用途。
> * 请注意收发敏感信息，本项目不做信息过滤。

# 使用docker运行

你可以使用docker快速运行本项目。

`第一种：基于环境变量运行`

```sh
# 运行项目，环境变量参考下方配置说明
$ docker run -itd --name wechatbot --restart=always -e APIKEY=xxxx -e AUTO_PASS=false -e SESSION_TIMEOUT=60s -e MODEL=text-davinci-003 -e MAX_TOKENS=512 -e TEMPREATURE=0.9 -e REPLY_PREFIX=我是来自机器人回复: docker.mirrors.sjtug.sjtu.edu.cn/qingshui869413421/wechatbot:latest

# 查看二维码
$ docker exec -it wechatbot bash 
$ tail -f -n 50 /app/run.log 
```

运行命令中映射的配置文件参考下边的配置文件说明。

`第二种：基于配置文件挂载运行`

```sh
# 复制配置文件，根据自己实际情况，调整配置里的内容
cp config.dev.json config.json  # 其中 config.dev.json 从项目的根目录获取

# 运行项目
docker run -itd --name wechatbot -v `pwd`/config.json:/app/config.json docker.mirrors.sjtug.sjtu.edu.cn/qingshui869413421/wechatbot:latest

# 查看二维码
$ docker exec -it wechatbot bash 
$ tail -f -n 50 /app/run.log 
```

其中配置文件参考下边的配置文件说明。

# 快速开始

> 非技术人员请直接下载release中的[压缩包](https://github.com/869413421/wechatbot/releases/tag/v1.1.2) ，解压运行。

````
# 获取项目
git clone https://github.com/869413421/wechatbot.git

# 进入项目目录
cd wechatbot

# 复制配置文件
copy config.dev.json config.json

# 启动项目
go run main.go

# linux编译，守护进程运行（可选）
# 编译
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o wechatbot  ./main.go
# 守护进程运行
nohup ./wechatbot > run.log &
````

# 配置文件说明

````
{
  "api_key": "your api key",
  "auto_pass": false,
  "session_timeout": 60,
  "max_tokens": 512,
  "model": "text-davinci-003",
  "temperature": 0.9,
  "reply_prefix": "来自机器人回复："
}

api_key：openai api_key
auto_pass:是否自动通过好友添加
session_timeout：会话超时时间，默认60秒，单位秒，在会话时间内所有发送给机器人的信息会作为上下文。
max_tokens: GPT响应字符数，最大2048，默认值512。max_tokens会影响接口响应速度，字符越大响应越慢。
model: GPT选用模型，默认text-davinci-003，具体选项参考官网训练场
temperature: GPT热度，0到1，默认0.9。数字越大创造力越强，但更偏离训练事实，越低越接近训练事实
reply_prefix: 私聊回复前缀
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

