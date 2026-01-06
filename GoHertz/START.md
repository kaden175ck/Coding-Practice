# 🚀 开始使用 Hertz Demo

## 新手推荐流程

### 步骤 1：启动服务器 ⭐

打开终端（CMD/PowerShell/Terminal）：

```bash
cd GoHertz
go run .
```

✅ 看到 "服务器启动在 http://127.0.0.1:8080" 表示成功
⚠️ **保持这个窗口打开，不要关闭！**

---

### 步骤 2：浏览器快速验证

打开浏览器，访问：
- http://localhost:8080/

看到 JSON 响应表示服务器正常运行！

---

### 步骤 3：运行自动化测试

**打开新的终端窗口**，运行测试脚本：

**Windows 用户：**
```bash
cd GoHertz
test-api.bat
```

**Mac/Linux 用户：**
```bash
cd GoHertz
./test-api.sh
```

你会看到所有 API 的测试结果！

---

### 步骤 4：观察日志输出

回到第一个终端窗口（运行服务器的窗口），你会看到：

```
[开始] GET /users - 客户端: 127.0.0.1
[完成] GET /users - 状态码: 200 - 耗时: 123.4µs
[开始] POST /users - 客户端: 127.0.0.1
[完成] POST /users - 状态码: 201 - 耗时: 456.7µs
```

这是日志中间件的输出，显示每个请求的详细信息。

---

### 步骤 5：手动测试（可选）

如果你想手动测试某个接口：

```bash
# 新终端窗口
curl http://localhost:8080/users
```

---

## 📱 使用工具测试（进阶）

如果你有以下工具，测试会更方便：

### Postman（推荐）
1. 下载：https://www.postman.com/downloads/
2. 创建新请求
3. 输入 URL：http://localhost:8080/users
4. 点击 Send

### VS Code REST Client 插件
1. 安装插件：REST Client
2. 创建 `test.http` 文件
3. 写入：
```
GET http://localhost:8080/users
```
4. 点击 "Send Request"

---

## 🔍 常见问题

### Q: 运行 `go run .` 时报错？

**A: 确保已安装依赖：**
```bash
go mod tidy
```

### Q: curl 命令不存在？

**A: Windows 用户：**
- Windows 10/11 自带 curl
- 如果没有，使用浏览器或下载 Git Bash

### Q: 如何停止服务器？

**A: 在服务器终端按 `Ctrl + C`**

### Q: 端口 8080 被占用？

**A: 修改 main.go 中的端口：**
```go
h := server.Default(server.WithHostPorts("127.0.0.1:3000"))
```

---

## 🎓 下一步学习

1. ✅ 完成上述测试
2. 📖 阅读 `main.go` 的代码和注释
3. 📖 阅读 `handler.go` 了解请求处理
4. 📖 阅读 `middleware.go` 了解中间件
5. 🔧 尝试修改代码，添加新的路由
6. ✍️ 查看 `README.md` 了解更多细节

---

## 💡 快捷命令

```bash
# 运行服务器
go run .

# 编译项目
go build

# 运行测试
go test -v

# 格式化代码
go fmt ./...

# 停止服务器
Ctrl + C
```

开始探索吧！🎉
