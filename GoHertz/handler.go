package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ========== 用户相关的处理函数 ==========

// GetUsers 获取所有用户列表
// 这是一个简单的 GET 请求处理函数
func GetUsers(ctx context.Context, c *app.RequestContext) {
	// 返回模拟的用户数据列表
	// 在实际项目中，这里应该从数据库查询数据
	users := []User{
		{ID: 1, Name: "张三", Email: "zhangsan@example.com", Age: 25},
		{ID: 2, Name: "李四", Email: "lisi@example.com", Age: 30},
		{ID: 3, Name: "王五", Email: "wangwu@example.com", Age: 28},
	}

	// 返回 JSON 格式的响应
	// consts.StatusOK 是 HTTP 200 状态码
	c.JSON(consts.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "获取用户列表成功",
		"data":    users,
		"total":   len(users),
	})
}

// GetUserByID 根据 ID 获取单个用户
// 这个函数演示如何获取路径参数（URL 中的参数）
func GetUserByID(ctx context.Context, c *app.RequestContext) {
	// 从 URL 路径中获取参数 :id
	// 例如：/users/1 中的 "1"
	idStr := c.Param("id")

	// 将字符串类型的 ID 转换为整数
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// 如果转换失败，返回 400 错误
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "无效的用户 ID",
			"error":   err.Error(),
		})
		return
	}

	// 模拟从数据库查找用户
	// 在实际项目中，这里应该查询数据库
	user := User{
		ID:    id,
		Name:  fmt.Sprintf("用户%d", id),
		Email: fmt.Sprintf("user%d@example.com", id),
		Age:   20 + id,
	}

	// 返回找到的用户数据
	c.JSON(consts.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "获取用户成功",
		"data":    user,
	})
}

// CreateUser 创建新用户
// 这个函数演示如何处理 POST 请求和解析请求体（Request Body）
func CreateUser(ctx context.Context, c *app.RequestContext) {
	// 定义一个变量来接收请求体中的数据
	var newUser User

	// 将请求体中的 JSON 数据绑定到 newUser 变量
	// BindJSON 会自动解析 JSON 并填充到结构体中
	if err := c.BindJSON(&newUser); err != nil {
		// 如果解析失败，返回 400 错误
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "请求数据格式错误",
			"error":   err.Error(),
		})
		return
	}

	// 验证数据（简单示例）
	if newUser.Name == "" {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "用户名不能为空",
		})
		return
	}

	if newUser.Email == "" {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "邮箱不能为空",
		})
		return
	}

	// 模拟保存到数据库，生成一个 ID
	newUser.ID = 100 // 在实际项目中，这个 ID 应该由数据库自动生成

	// 返回创建成功的响应
	// consts.StatusCreated 是 HTTP 201 状态码，表示资源创建成功
	c.JSON(consts.StatusCreated, map[string]interface{}{
		"code":    201,
		"message": "用户创建成功",
		"data":    newUser,
	})
}

// UpdateUser 更新用户信息
// 这个函数演示如何同时处理路径参数和请求体
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	// 获取要更新的用户 ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "无效的用户 ID",
		})
		return
	}

	// 解析请求体中的更新数据
	var updateData User
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "请求数据格式错误",
			"error":   err.Error(),
		})
		return
	}

	// 设置 ID（确保更新的是正确的用户）
	updateData.ID = id

	// 模拟更新数据库操作
	// 在实际项目中，这里应该执行 UPDATE SQL 语句

	// 返回更新后的数据
	c.JSON(consts.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "用户更新成功",
		"data":    updateData,
	})
}

// DeleteUser 删除用户
// 这个函数演示 DELETE 请求的处理
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	// 获取要删除的用户 ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "无效的用户 ID",
		})
		return
	}

	// 模拟删除操作
	// 在实际项目中，这里应该执行 DELETE SQL 语句

	// 返回删除成功的响应
	c.JSON(consts.StatusOK, map[string]interface{}{
		"code":    200,
		"message": fmt.Sprintf("用户 %d 删除成功", id),
	})
}

// SearchUsers 搜索用户
// 这个函数演示如何获取 Query 参数（URL 中 ? 后面的参数）
func SearchUsers(ctx context.Context, c *app.RequestContext) {
	// 获取查询参数
	// 例如：/search?name=张三&age=25
	name := c.Query("name")    // 获取 name 参数
	ageStr := c.Query("age")   // 获取 age 参数
	page := c.DefaultQuery("page", "1") // 获取 page 参数，如果不存在则默认为 "1"

	// 构建响应信息
	result := map[string]interface{}{
		"code":    200,
		"message": "搜索完成",
		"query": map[string]string{
			"name": name,
			"age":  ageStr,
			"page": page,
		},
	}

	// 如果提供了 name 参数，模拟搜索结果
	if name != "" {
		result["data"] = []User{
			{ID: 1, Name: name, Email: name + "@example.com", Age: 25},
		}
		result["total"] = 1
	} else {
		result["data"] = []User{}
		result["total"] = 0
		result["message"] = "请提供搜索参数 name"
	}

	c.JSON(consts.StatusOK, result)
}

// ========== 产品相关的处理函数 ==========

// GetProducts 获取产品列表
// 这个函数用于演示路由组的使用
func GetProducts(ctx context.Context, c *app.RequestContext) {
	products := []Product{
		{ID: 1, Name: "笔记本电脑", Price: 5999.99, Stock: 50},
		{ID: 2, Name: "无线鼠标", Price: 99.99, Stock: 200},
		{ID: 3, Name: "机械键盘", Price: 399.99, Stock: 100},
	}

	c.JSON(consts.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "获取产品列表成功",
		"data":    products,
		"total":   len(products),
	})
}

// GetProductByID 根据 ID 获取产品详情
func GetProductByID(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "无效的产品 ID",
		})
		return
	}

	// 模拟产品数据
	product := Product{
		ID:    id,
		Name:  fmt.Sprintf("产品%d", id),
		Price: 99.99 * float64(id),
		Stock: 100,
	}

	c.JSON(consts.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "获取产品成功",
		"data":    product,
	})
}
