# Go项目示例 (go-project-example)

## 项目介绍

这是一个基于Go语言开发的社区论坛系统示例项目，使用了Gin框架作为Web服务器，GORM作为ORM框架，MySQL作为数据库存储。该项目展示了一个标准Go项目的结构和最佳实践，包括API处理、业务逻辑、数据存储、测试以及并发编程等方面。

## 功能特性

- 用户系统：支持用户信息管理
- 话题系统：支持创建和查询话题
- 帖子系统：支持发布帖子和回复
- RESTful API：提供标准的REST风格接口
- 并发编程示例：包含goroutine和channel的示例代码
- Go编程注意事项：包含字符串处理、JSON操作、闭包和数组处理的示例

## 技术栈

- [Go](https://golang.org/) - 编程语言
- [Gin](https://github.com/gin-gonic/gin) - Web框架
- [GORM](https://gorm.io/) - ORM框架
- [MySQL](https://www.mysql.com/) - 数据库
- [Zap](https://github.com/uber-go/zap) - 日志库

## 项目结构

```
.
├── attention/            # Go编程注意事项的示例代码
│   ├── array.go          # 数组操作示例
│   ├── closure.go        # 闭包示例
│   ├── json.go           # JSON处理示例
│   └── string.go         # 字符串处理示例
├── concurrence/          # 并发编程示例
│   ├── channel.go        # 通道示例
│   └── goroutine.go      # goroutine示例
├── handler/              # HTTP请求处理层
│   ├── publish_post.go   # 发布帖子处理
│   └── query_page_info.go # 查询页面信息处理
├── repository/           # 数据持久层
│   ├── db_init.go        # 数据库初始化
│   ├── post.go           # 帖子数据操作
│   ├── topic.go          # 话题数据操作
│   └── user.go           # 用户数据操作
├── service/              # 业务逻辑层
│   ├── publish_post.go       # 发布帖子业务逻辑
│   ├── publish_post_test.go  # 发布帖子测试
│   ├── query_page_info.go    # 查询页面信息业务逻辑
│   └── query_page_info_test.go # 查询页面信息测试
├── util/                 # 工具类
│   └── logger.go         # 日志工具
├── .gitignore            # Git忽略文件
├── example.sql           # 数据库示例SQL
├── go.mod                # Go模块定义
├── go.sum                # Go依赖校验
├── LICENSE               # 许可证
├── README.md             # 项目说明文档
└── sever.go              # 主服务入口
```

## 安装和使用

### 前置要求

- Go 1.16+
- MySQL 5.7+

### 数据库设置

1. 创建MySQL数据库和表结构：

```bash
mysql -u root -p < example.sql
```

### 项目配置

1. 根据需要修改数据库连接配置（位于 `repository/db_init.go`）

```go
dsn := "root:00000000@tcp(127.0.0.1:3306)/community?charset=utf8mb4&parseTime=True&loc=Local"
```

### 运行项目

1. 克隆代码仓库

```bash
git clone https://github.com/Moonlight-Zhao/go-project-example.git
cd go-project-example
```

2. 安装依赖

```bash
go mod download
```

3. 编译和运行

```bash
go build -o app
./app
```

或者直接运行：

```bash
go run sever.go
```

4. 访问API

```
GET http://localhost:8080/ping               # 健康检查
GET http://localhost:8080/community/page/get/:id  # 获取指定ID的话题页面信息
POST http://localhost:8080/community/post/do     # 发布帖子
```

## API说明

### 获取话题页面信息

```
GET /community/page/get/:id
```

参数：
- `id`: 话题ID

返回示例：
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "topic": { ... },
    "post_list": [ ... ]
  }
}
```

### 发布帖子

```
POST /community/post/do
```

表单参数：
- `uid`: 用户ID
- `topic_id`: 话题ID
- `content`: 帖子内容

返回示例：
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "post_id": 123
  }
}
```

## 学习资源

项目包含了Go编程中的一些常见注意事项和并发编程示例，可以作为学习参考：

- `attention/`: 包含字符串处理、JSON操作、闭包和数组处理的示例和测试
- `concurrence/`: 包含goroutine和channel的示例代码和测试

## 贡献指南

1. Fork 该仓库
2. 创建新的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的修改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解更多详情。
