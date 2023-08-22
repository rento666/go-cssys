# go-cssys 使用Go语言开发的一款客服系统

本项目全称`Go-CustomerServiceSystem`。

## 项目目录介绍

```text
.
├── README.md
├── app
│   ├── client
│   │   ├── enter.go        // 客户端的入口文件，client->main方法调用这里的方法。
│   │   └── write           // client端发送消息
│   └── server
│       └── enter.go        // 同上
├── cmd                     // 存放main.go
│   ├── client              // 客户端的main.go
│   │   └── main.go
│   └── server
│       └── main.go
├── config                  // 存放配置信息
│   └── conf.go
└── go.mod                  // go mod 包管理

```

