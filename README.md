# dingtalk-robot
[![Go Reference](https://pkg.go.dev/badge/github.com/group-robot/dingtalk-robot.svg)](https://pkg.go.dev/github.com/group-robot/dingtalk-robot) 钉钉自定义机器人

# Example

### text
```go
textMessage := NewTextMessage("我就是我, @XXX 是不一样的烟火")
textMessage.At = NewAt(true)
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
res, err := client.SendMessage(textMessage)
```

### link
```go
linkMsg := NewLinkMessage("这个即将发布的新版本，创始人xx称它为红树林。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是红树林",
"时代的火车向前开",
"https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
)
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
res, err := client.SendMessage(linkMsg)
```

### markdown
```go
markdownMsg := NewMarkdownMessage("杭州天气", "#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n")
markdownMsg.At = NewAt(true)
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
res, err := client.SendMessage(markdownMsg)
```

### card
#### single action card
```go
singleActionCardMsg := NewSingleActionCardMessage("乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
"![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png) \n ### 乔布斯 20 年前想打造的苹果咖啡厅 \n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
"阅读全文",
"https://www.dingtalk.com/",
)
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
res, err := client.SendMessage(singleActionCardMsg)
```

#### action card
```go
actionCardMsg := NewActionCardMessage(
"我 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
"![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
NewActionCardButton(
"内容不错",
"https://www.dingtalk.com/",
),
)
actionCardMsg.AddButtons(

NewActionCardButton(
"不感兴趣",
"https://www.dingtalk.com/",
),
)
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
res, err := client.SendMessage(actionCardMsg)
```

#### feed card
```go
feedCardMsg := NewFeedCardMessage(
NewLink(
"时代的火车向前开1",
"https://www.dingtalk.com/",
"https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
),
)
feedCardMsg.AddLinks(
NewLink(
"时代的火车向前开2",
"https://www.dingtalk.com/",
"https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
),
)
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
res, err := client.SendMessage(feedCardMsg)
```

# Client
```go
textMsg := `{"at":{"atMobiles":["180xxxxxx"],"atUserIds":["user123"],"isAtAll":false},"text":{"content":"我就是我, @XXX 是不一样的烟火"},"msgtype":"text"}`
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
res, err := client.SendMessageStr(textMsg)
```