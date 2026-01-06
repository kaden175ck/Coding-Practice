package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

// ========== 单元测试示例 ==========
// 在 Go 中，测试文件以 _test.go 结尾
// 测试函数以 Test 开头，参数是 *testing.T

// TestGetUsers 测试获取用户列表接口
func TestGetUsers(t *testing.T) {
	// 创建一个测试用的 HTTP 引擎
	w := ut.PerformRequest(
		// 创建测试引擎并注册路由
		ut.NewEngine(),
		// 执行 HTTP 方法
		"GET",
		// 请求路径
		"/users",
		// 请求体（GET 请求为 nil）
		nil,
		// 注册测试路由
		ut.Header{Key: "Host", Value: "localhost"},
	)

	// 断言：验证响应状态码是否为 200
	// 注意：这个测试会失败，因为我们没有在测试引擎中注册路由
	// 这只是一个示例，展示如何编写测试
	t.Logf("响应状态码: %d", w.Code)
	t.Logf("响应内容: %s", w.Body.String())
}

// TestGetUserByID 测试获取单个用户接口
func TestGetUserByID(t *testing.T) {
	// 创建 Hertz 测试上下文
	c := app.NewContext(0)

	// 模拟请求参数
	c.Params = append(c.Params, app.Param{
		Key:   "id",
		Value: "1",
	})

	// 调用处理函数
	GetUserByID(context.Background(), c)

	// 获取响应体
	resp := c.Response.Body()

	// 解析 JSON 响应
	var result map[string]interface{}
	err := json.Unmarshal(resp, &result)

	// 断言：检查是否成功解析 JSON
	assert.Nil(t, err)

	// 断言：检查响应码
	if code, ok := result["code"].(float64); ok {
		assert.DeepEqual(t, 200.0, code)
		t.Logf("测试通过：响应码 = %.0f", code)
	}

	// 打印完整响应（用于调试）
	t.Logf("完整响应: %s", string(resp))
}

// TestCreateUser 测试创建用户接口
func TestCreateUser(t *testing.T) {
	// 准备测试数据
	newUser := map[string]interface{}{
		"name":  "测试用户",
		"email": "test@example.com",
		"age":   25,
	}

	// 将数据转换为 JSON
	jsonData, err := json.Marshal(newUser)
	assert.Nil(t, err)

	// 创建测试上下文
	c := app.NewContext(0)

	// 设置请求体
	c.Request.SetBody(jsonData)
	c.Request.Header.SetContentType("application/json")

	// 调用处理函数
	CreateUser(context.Background(), c)

	// 获取响应
	resp := c.Response.Body()
	var result map[string]interface{}
	json.Unmarshal(resp, &result)

	// 断言：检查响应码是否为 201（创建成功）
	if code, ok := result["code"].(float64); ok {
		assert.DeepEqual(t, 201.0, code)
		t.Logf("创建用户成功：%s", string(resp))
	}
}

// TestUserValidate 测试用户数据验证方法
func TestUserValidate(t *testing.T) {
	// 测试用例 1：有效的用户数据
	validUser := User{
		ID:    1,
		Name:  "张三",
		Email: "zhangsan@example.com",
		Age:   25,
	}

	err := validUser.Validate()
	assert.Nil(t, err)
	t.Log("测试通过：有效用户数据验证成功")

	// 测试用例 2：无效的用户数据（缺少名字）
	invalidUser := User{
		ID:    2,
		Email: "test@example.com",
		Age:   25,
	}

	err = invalidUser.Validate()
	assert.NotNil(t, err)
	t.Logf("测试通过：检测到无效数据 - %v", err)

	// 测试用例 3：无效的年龄
	invalidAgeUser := User{
		ID:    3,
		Name:  "李四",
		Email: "lisi@example.com",
		Age:   200, // 年龄超出范围
	}

	err = invalidAgeUser.Validate()
	assert.NotNil(t, err)
	t.Logf("测试通过：检测到无效年龄 - %v", err)
}

// TestProductIsInStock 测试产品库存检查方法
func TestProductIsInStock(t *testing.T) {
	// 有库存的产品
	product1 := Product{
		ID:    1,
		Name:  "笔记本电脑",
		Price: 5999.99,
		Stock: 10,
	}

	assert.True(t, product1.IsInStock())
	t.Log("测试通过：产品有库存")

	// 无库存的产品
	product2 := Product{
		ID:    2,
		Name:  "鼠标",
		Price: 99.99,
		Stock: 0,
	}

	assert.False(t, product2.IsInStock())
	t.Log("测试通过：产品无库存")
}

// BenchmarkGetUserByID 性能测试示例
// 基准测试函数以 Benchmark 开头
func BenchmarkGetUserByID(b *testing.B) {
	// b.N 会自动调整，以获得可靠的测试结果
	for i := 0; i < b.N; i++ {
		c := app.NewContext(0)
		c.Params = append(c.Params, app.Param{
			Key:   "id",
			Value: "1",
		})
		GetUserByID(context.Background(), c)
	}
}

// ========== 运行测试的命令 ==========
// 运行所有测试：go test
// 运行特定测试：go test -run TestGetUserByID
// 运行性能测试：go test -bench=.
// 查看测试覆盖率：go test -cover
// 生成覆盖率报告：go test -coverprofile=coverage.out && go tool cover -html=coverage.out
