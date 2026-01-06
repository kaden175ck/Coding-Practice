# Hertz 框架学习 Demo

这是一个基于 CloudWeGo Hertz 框架的学习示例项目，包含了详细的中文注释，帮助你快速入门 Hertz 框架。

## 什么是 Hertz？

Hertz 是 CloudWeGo 开源的一个高性能、高可扩展的 Go HTTP 框架。它具有以下特点：

- 🚀 高性能：基于自研的网络库，性能优异
- 🛠️ 易扩展：提供了丰富的扩展接口
- 📦 开箱即用：提供了常用的中间件和工具
- 💡 简单易用：API 设计简洁，学习成本低

## 项目结构

```
GoHertz/
├── main.go          # 主入口文件，定义路由和启动服务器
├── handler.go       # 路由处理函数，处理各种 HTTP 请求
├── middleware.go    # 中间件函数，用于日志、认证等
├── model.go         # 数据模型定义
├── go.mod           # Go 模块配置文件
└── README.md        # 项目说明文档
```

## 安装和运行

### 1. 前置要求

- Go 1.16 或更高版本
- 已配置好 Go 环境变量

### 2. 安装依赖

```bash
cd GoHertz
go mod tidy
```

这个命令会自动下载项目所需的所有依赖包，包括 Hertz 框架。

### 3. 运行项目

```bash
go run .
```

或者

```bash
go run main.go handler.go middleware.go model.go
```

看到以下输出表示启动成功：

```
服务器启动在 http://127.0.0.1:8080
访问 http://127.0.0.1:8080 查看欢迎信息
按 Ctrl+C 停止服务器
```

## API 接口说明

### 基础接口

| 方法 | 路径 | 说明 | 示例 |
|-----|------|------|------|
| GET | `/` | 欢迎页面 | `curl http://localhost:8080/` |

### 用户管理接口

| 方法 | 路径 | 说明 | 示例 |
|-----|------|------|------|
| GET | `/users` | 获取所有用户 | `curl http://localhost:8080/users` |
| GET | `/users/:id` | 获取指定用户 | `curl http://localhost:8080/users/1` |
| POST | `/users` | 创建新用户 | 见下方示例 |
| PUT | `/users/:id` | 更新用户信息 | 见下方示例 |
| DELETE | `/users/:id` | 删除用户 | `curl -X DELETE http://localhost:8080/users/1` |
| GET | `/search` | 搜索用户 | `curl "http://localhost:8080/search?name=张三&age=25"` |

### 产品管理接口（需要认证）

| 方法 | 路径 | 说明 | 示例 |
|-----|------|------|------|
| GET | `/api/v1/products` | 获取所有产品 | 见下方示例 |
| GET | `/api/v1/products/:id` | 获取指定产品 | 见下方示例 |

## 使用示例

### 1. 获取用户列表

```bash
curl http://localhost:8080/users
```

响应：
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "name": "张三",
      "email": "zhangsan@example.com",
      "age": 25
    },
    {
      "id": 2,
      "name": "李四",
      "email": "lisi@example.com",
      "age": 30
    }
  ],
  "message": "获取用户列表成功",
  "total": 2
}
```

### 2. 创建新用户

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "王五",
    "email": "wangwu@example.com",
    "age": 28
  }'
```

响应：
```json
{
  "code": 201,
  "message": "用户创建成功",
  "data": {
    "id": 100,
    "name": "王五",
    "email": "wangwu@example.com",
    "age": 28
  }
}
```

### 3. 更新用户信息

```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三丰",
    "age": 30
  }'
```

### 4. 搜索用户（Query 参数示例）

```bash
curl "http://localhost:8080/search?name=张三&page=1"
```

### 5. 访问需要认证的接口

访问 `/api/v1/products` 需要在请求头中提供 token：

```bash
# 不带 token（会返回 401 错误）
curl http://localhost:8080/api/v1/products

# 带正确的 token
curl http://localhost:8080/api/v1/products \
  -H "Authorization: Bearer demo-token-123456"
```

## 学习要点

### 1. 路由定义（main.go）

```go
// GET 请求
h.GET("/users", GetUsers)

// 带路径参数的路由
h.GET("/users/:id", GetUserByID)

// POST 请求
h.POST("/users", CreateUser)

// 路由组
apiV1 := h.Group("/api/v1")
apiV1.GET("/products", GetProducts)
```

### 2. 请求处理（handler.go）

```go
func GetUsers(ctx context.Context, c *app.RequestContext) {
    // 获取路径参数
    id := c.Param("id")

    // 获取 Query 参数
    name := c.Query("name")

    // 解析 JSON 请求体
    var user User
    c.BindJSON(&user)

    // 返回 JSON 响应
    c.JSON(consts.StatusOK, map[string]interface{}{
        "message": "success",
        "data": user,
    })
}
```

### 3. 中间件（middleware.go）

```go
// 定义中间件
func LoggerMiddleware() app.HandlerFunc {
    return func(ctx context.Context, c *app.RequestContext) {
        // 请求前的处理

        c.Next(ctx) // 执行下一个中间件/处理函数

        // 请求后的处理
    }
}

// 使用中间件
h.Use(LoggerMiddleware())
```

### 4. 数据模型（model.go）

```go
// 定义结构体
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

// 结构体方法
func (u *User) Validate() error {
    if u.Name == "" {
        return fmt.Errorf("用户名不能为空")
    }
    return nil
}
```

## 进阶学习

1. **数据库集成**：可以集成 GORM 或其他 ORM 来操作数据库
2. **配置管理**：使用 Viper 或其他配置库管理配置
3. **日志系统**：集成结构化日志库如 zap 或 logrus
4. **参数验证**：使用 validator 库进行请求参数验证
5. **单元测试**：编写测试用例保证代码质量

## 常见问题

### Q1: 如何修改服务器端口？

在 `main.go` 中修改：
```go
h := server.Default(server.WithHostPorts("127.0.0.1:3000"))
```

### Q2: 如何添加新的路由？

1. 在 `handler.go` 中添加处理函数
2. 在 `main.go` 中注册路由

### Q3: 中间件的执行顺序是什么？

中间件按照注册顺序执行，`c.Next(ctx)` 之前的代码在请求处理前执行，之后的代码在请求处理后执行。

## 参考资料

- [Hertz 官方文档](https://www.cloudwego.io/zh/docs/hertz/)
- [Hertz GitHub](https://github.com/cloudwego/hertz)
- [Go 语言官方文档](https://golang.org/doc/)

## 许可证

本项目仅供学习使用。
