package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ========== 中间件函数 ==========
// 中间件是在请求到达处理函数之前或之后执行的函数
// 可以用于：日志记录、身份验证、权限检查、请求处理等

// LoggerMiddleware 日志中间件
// 记录每个 HTTP 请求的信息：方法、路径、耗时等
func LoggerMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 记录请求开始时间
		startTime := time.Now()

		// 获取请求信息
		method := string(c.Method())       // 请求方法（GET, POST, PUT, DELETE 等）
		path := string(c.Path())           // 请求路径
		clientIP := c.ClientIP()           // 客户端 IP 地址

		// 打印请求开始信息
		fmt.Printf("[开始] %s %s - 客户端: %s\n", method, path, clientIP)

		// c.Next() 会执行下一个中间件或处理函数
		// 这是中间件链中非常重要的一步
		c.Next(ctx)

		// 请求处理完成后，会继续执行下面的代码
		// 计算请求处理耗时
		duration := time.Since(startTime)

		// 获取响应状态码
		statusCode := c.Response.StatusCode()

		// 打印请求完成信息
		fmt.Printf("[完成] %s %s - 状态码: %d - 耗时: %v\n",
			method, path, statusCode, duration)
	}
}

// AuthMiddleware 身份验证中间件
// 检查请求头中的 token，验证用户身份
func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求头中获取 Authorization token
		// 例如：Authorization: Bearer your-token-here
		token := c.GetHeader("Authorization")

		// 如果没有提供 token
		if len(token) == 0 {
			// 返回 401 未授权错误，并中止请求
			c.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "未提供认证令牌",
				"error":   "请在请求头中添加 Authorization 字段",
			})
			// c.Abort() 会停止执行后续的中间件和处理函数
			c.Abort()
			return
		}

		// 简单的 token 验证示例
		// 在实际项目中，应该验证 JWT token 或查询数据库验证 session
		validToken := "Bearer demo-token-123456"
		if string(token) != validToken {
			c.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "无效的认证令牌",
				"error":   "令牌验证失败",
			})
			c.Abort()
			return
		}

		// token 验证通过，可以在上下文中存储用户信息
		// 后续的处理函数可以通过 c.Get() 获取这些信息
		c.Set("userID", 12345)
		c.Set("username", "demo_user")

		fmt.Println("[认证] 用户验证成功 - 用户ID: 12345")

		// 继续执行下一个中间件或处理函数
		c.Next(ctx)
	}
}

// CORSMiddleware 跨域资源共享（CORS）中间件
// 允许前端应用从不同的域名访问 API
func CORSMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 设置 CORS 响应头
		// 允许所有域名访问（生产环境应该设置具体的域名）
		c.Header("Access-Control-Allow-Origin", "*")

		// 允许的 HTTP 方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 允许的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 预检请求的缓存时间（秒）
		c.Header("Access-Control-Max-Age", "86400")

		// 如果是 OPTIONS 预检请求，直接返回 204
		if string(c.Method()) == "OPTIONS" {
			c.AbortWithStatus(consts.StatusNoContent)
			return
		}

		// 继续处理请求
		c.Next(ctx)
	}
}

// RateLimitMiddleware 限流中间件（简单示例）
// 限制请求频率，防止 API 被滥用
func RateLimitMiddleware(maxRequests int) app.HandlerFunc {
	// 使用一个简单的计数器来演示（实际应该使用 Redis 或其他方案）
	var requestCount int

	return func(ctx context.Context, c *app.RequestContext) {
		requestCount++

		// 如果请求次数超过限制
		if requestCount > maxRequests {
			c.JSON(consts.StatusTooManyRequests, map[string]interface{}{
				"code":    429,
				"message": "请求过于频繁，请稍后再试",
				"error":   "超过速率限制",
			})
			c.Abort()
			return
		}

		fmt.Printf("[限流] 当前请求计数: %d/%d\n", requestCount, maxRequests)

		// 继续处理请求
		c.Next(ctx)
	}
}

// RecoveryMiddleware 恢复中间件
// 捕获 panic 错误，防止服务器崩溃
func RecoveryMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// defer 会在函数返回前执行
		defer func() {
			// recover() 可以捕获 panic
			if err := recover(); err != nil {
				fmt.Printf("[错误] 捕获到 panic: %v\n", err)

				// 返回 500 错误
				c.JSON(consts.StatusInternalServerError, map[string]interface{}{
					"code":    500,
					"message": "服务器内部错误",
					"error":   fmt.Sprintf("%v", err),
				})

				// 中止请求
				c.Abort()
			}
		}()

		// 继续执行
		c.Next(ctx)
	}
}
