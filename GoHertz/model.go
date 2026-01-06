package main

import "fmt"

// ========== 数据模型定义 ==========
// 在 Go 中，我们使用结构体（struct）来定义数据模型

// User 用户模型
// 定义用户的数据结构
type User struct {
	// ID 用户唯一标识
	// `json:"id"` 是结构体标签（tag），用于 JSON 序列化时的字段名
	ID int `json:"id"`

	// Name 用户名
	// omitempty 表示如果字段为空值，在 JSON 序列化时会忽略该字段
	Name string `json:"name,omitempty"`

	// Email 邮箱地址
	Email string `json:"email,omitempty"`

	// Age 年龄
	Age int `json:"age,omitempty"`
}

// Product 产品模型
// 定义产品的数据结构
type Product struct {
	// ID 产品唯一标识
	ID int `json:"id"`

	// Name 产品名称
	Name string `json:"name"`

	// Price 产品价格
	// 使用 float64 来存储价格，支持小数
	Price float64 `json:"price"`

	// Stock 库存数量
	Stock int `json:"stock"`

	// Description 产品描述（可选字段）
	Description string `json:"description,omitempty"`
}

// Response 通用响应结构
// 用于统一 API 响应格式
type Response struct {
	// Code 响应状态码（业务状态码，不是 HTTP 状态码）
	Code int `json:"code"`

	// Message 响应消息
	Message string `json:"message"`

	// Data 响应数据（可以是任何类型）
	// interface{} 表示可以存储任何类型的值
	Data interface{} `json:"data,omitempty"`
}

// ErrorResponse 错误响应结构
// 用于统一错误返回格式
type ErrorResponse struct {
	// Code 错误状态码
	Code int `json:"code"`

	// Message 错误消息
	Message string `json:"message"`

	// Error 详细错误信息
	Error string `json:"error,omitempty"`
}

// UserCreateRequest 创建用户请求结构
// 用于接收创建用户的请求数据
type UserCreateRequest struct {
	// Name 用户名（必填）
	// binding:"required" 表示这是一个必填字段
	Name string `json:"name" binding:"required"`

	// Email 邮箱（必填）
	Email string `json:"email" binding:"required"`

	// Age 年龄（可选）
	Age int `json:"age"`
}

// UserUpdateRequest 更新用户请求结构
// 用于接收更新用户的请求数据
type UserUpdateRequest struct {
	// Name 用户名（可选）
	Name string `json:"name,omitempty"`

	// Email 邮箱（可选）
	Email string `json:"email,omitempty"`

	// Age 年龄（可选）
	Age int `json:"age,omitempty"`
}

// PageRequest 分页请求结构
// 用于处理分页查询
type PageRequest struct {
	// Page 页码（从 1 开始）
	Page int `json:"page" form:"page"`

	// PageSize 每页数量
	PageSize int `json:"page_size" form:"page_size"`
}

// PageResponse 分页响应结构
// 用于返回分页数据
type PageResponse struct {
	// Total 总记录数
	Total int `json:"total"`

	// Page 当前页码
	Page int `json:"page"`

	// PageSize 每页数量
	PageSize int `json:"page_size"`

	// Data 数据列表
	Data interface{} `json:"data"`
}

// ========== 模型方法示例 ==========

// Validate 验证用户数据是否合法
// 这是 User 结构体的方法
func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("用户名不能为空")
	}
	if u.Email == "" {
		return fmt.Errorf("邮箱不能为空")
	}
	if u.Age < 0 || u.Age > 150 {
		return fmt.Errorf("年龄不合法")
	}
	return nil
}

// GetFullInfo 获取用户完整信息（格式化输出）
// 返回用户的友好描述字符串
func (u *User) GetFullInfo() string {
	return fmt.Sprintf("用户 #%d: %s (%s), 年龄: %d",
		u.ID, u.Name, u.Email, u.Age)
}

// IsInStock 检查产品是否有库存
// 这是 Product 结构体的方法
func (p *Product) IsInStock() bool {
	return p.Stock > 0
}

// GetPriceString 获取格式化的价格字符串
func (p *Product) GetPriceString() string {
	return fmt.Sprintf("¥%.2f", p.Price)
}

// ========== 辅助函数 ==========

// NewResponse 创建一个新的成功响应
func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse 创建一个新的错误响应
func NewErrorResponse(code int, message string, err error) *ErrorResponse {
	errResp := &ErrorResponse{
		Code:    code,
		Message: message,
	}
	if err != nil {
		errResp.Error = err.Error()
	}
	return errResp
}
