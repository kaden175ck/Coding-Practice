# 快速开始指南

## 立即开始

### 第一步：运行服务器

```bash
cd GoHertz
go run .
```

看到以下输出表示成功：
```
服务器启动在 http://127.0.0.1:8080
访问 http://127.0.0.1:8080 查看欢迎信息
按 Ctrl+C 停止服务器
```

### 第二步：测试 API

打开新的终端窗口，执行以下命令测试各个接口：

#### 1. 欢迎页面
```bash
curl http://localhost:8080/
```

#### 2. 获取用户列表
```bash
curl http://localhost:8080/users
```

#### 3. 获取单个用户
```bash
curl http://localhost:8080/users/1
```

#### 4. 创建用户
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"张三\",\"email\":\"zhangsan@test.com\",\"age\":25}"
```

#### 5. 搜索用户（Query 参数）
```bash
curl "http://localhost:8080/search?name=张三&age=25"
```

#### 6. 访问需要认证的接口

不带 token（会返回 401 错误）：
```bash
curl http://localhost:8080/api/v1/products
```

带正确的 token：
```bash
curl http://localhost:8080/api/v1/products \
  -H "Authorization: Bearer demo-token-123456"
```

## 学习路径

1. **理解基础**
   - 阅读 `main.go` - 了解如何创建服务器和注册路由
   - 阅读 `model.go` - 了解数据结构定义

2. **学习路由处理**
   - 阅读 `handler.go` - 了解如何处理不同类型的 HTTP 请求
   - 重点关注：
     - 路径参数：`c.Param("id")`
     - Query 参数：`c.Query("name")`
     - 请求体：`c.BindJSON(&user)`
     - JSON 响应：`c.JSON(status, data)`

3. **掌握中间件**
   - 阅读 `middleware.go` - 了解中间件的工作原理
   - 尝试添加自己的中间件

4. **编写测试**
   - 阅读 `handler_test.go` - 了解如何编写单元测试
   - 运行测试：`go test -v`

## 常用命令

```bash
# 运行项目
go run .

# 编译项目
go build -o app.exe .

# 运行编译后的程序
./app.exe  # Linux/Mac
app.exe    # Windows

# 运行测试
go test

# 运行测试并显示详细输出
go test -v

# 运行性能测试
go test -bench=.

# 查看测试覆盖率
go test -cover

# 格式化代码
go fmt ./...

# 检查代码问题
go vet ./...
```

## 下一步

- 尝试添加新的路由和处理函数
- 集成数据库（如 MySQL、PostgreSQL）
- 添加更多中间件（如 CORS、限流等）
- 学习 Hertz 的高级特性

详细内容请查看 `README.md`
