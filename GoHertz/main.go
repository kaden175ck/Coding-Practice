package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	// 创建一个 Hertz 服务器实例
	// WithHostPorts 设置服务器监听的地址和端口
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// ========== 注册全局中间件 ==========
	// 使用自定义的日志中间件，记录每个请求的信息
	h.Use(LoggerMiddleware())

	// ========== 基础路由示例 ==========

	// 1. 简单的 GET 请求 - 返回欢迎信息
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, map[string]interface{}{
			"message": "欢迎使用 Hertz 框架！",
			"version": "1.0.0",
		})
	})

	// 2. GET 请求 - 获取用户列表
	h.GET("/users", GetUsers)

	// 3. GET 请求 - 根据 ID 获取单个用户（路径参数示例）
	h.GET("/users/:id", GetUserByID)

	// 4. POST 请求 - 创建新用户（请求体示例）
	h.POST("/users", CreateUser)

	// 5. PUT 请求 - 更新用户信息
	h.PUT("/users/:id", UpdateUser)

	// 6. DELETE 请求 - 删除用户
	h.DELETE("/users/:id", DeleteUser)

	// 7. Query 参数示例 - 搜索用户
	h.GET("/search", SearchUsers)

	// ========== 路由组示例 ==========
	// 创建一个路由组，所有路由都以 /api/v1 开头
	apiV1 := h.Group("/api/v1")
	{
		// 可以为路由组添加特定的中间件
		apiV1.Use(AuthMiddleware())

		// /api/v1/products
		apiV1.GET("/products", GetProducts)

		// /api/v1/products/:id
		apiV1.GET("/products/:id", GetProductByID)
	}

	// ========== 错误处理示例 ==========
	h.GET("/error", func(ctx context.Context, c *app.RequestContext) {
		// 模拟错误场景
		c.JSON(consts.StatusInternalServerError, map[string]interface{}{
			"error":   "服务器内部错误",
			"code":    500,
			"message": "这是一个错误示例",
		})
	})

	// 启动服务器
	fmt.Println("服务器启动在 http://127.0.0.1:8080")
	fmt.Println("访问 http://127.0.0.1:8080 查看欢迎信息")
	fmt.Println("按 Ctrl+C 停止服务器")

	// Spin 会阻塞当前协程，直到服务器关闭
	h.Spin()
}
