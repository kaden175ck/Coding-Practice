package main

import (
	"fmt"
	"os"
	"strconv"
)

// ========== 配置管理示例 ==========
// 在实际项目中，通常使用配置文件（如 YAML、JSON）或环境变量来管理配置
// 这里展示一个简单的配置管理方式

// Config 应用配置结构
type Config struct {
	// Server 服务器配置
	Server ServerConfig

	// Database 数据库配置（示例）
	Database DatabaseConfig

	// App 应用配置
	App AppConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	// Host 服务器监听地址
	Host string

	// Port 服务器监听端口
	Port int

	// ReadTimeout 读取超时时间（秒）
	ReadTimeout int

	// WriteTimeout 写入超时时间（秒）
	WriteTimeout int
}

// DatabaseConfig 数据库配置（示例）
type DatabaseConfig struct {
	// Driver 数据库驱动（mysql, postgres, sqlite 等）
	Driver string

	// Host 数据库主机地址
	Host string

	// Port 数据库端口
	Port int

	// Username 数据库用户名
	Username string

	// Password 数据库密码
	Password string

	// Database 数据库名称
	Database string
}

// AppConfig 应用配置
type AppConfig struct {
	// Name 应用名称
	Name string

	// Version 应用版本
	Version string

	// Debug 是否开启调试模式
	Debug bool

	// LogLevel 日志级别（debug, info, warn, error）
	LogLevel string
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host:         "127.0.0.1",
			Port:         8080,
			ReadTimeout:  30,
			WriteTimeout: 30,
		},
		Database: DatabaseConfig{
			Driver:   "mysql",
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "",
			Database: "hertz_demo",
		},
		App: AppConfig{
			Name:     "Hertz Demo",
			Version:  "1.0.0",
			Debug:    true,
			LogLevel: "debug",
		},
	}
}

// LoadConfig 加载配置
// 这个函数演示如何从环境变量加载配置
func LoadConfig() *Config {
	// 先获取默认配置
	cfg := DefaultConfig()

	// 从环境变量读取配置（如果存在）
	// 例如：export SERVER_PORT=3000

	// 读取服务器端口
	if port := os.Getenv("SERVER_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			cfg.Server.Port = p
		}
	}

	// 读取服务器主机
	if host := os.Getenv("SERVER_HOST"); host != "" {
		cfg.Server.Host = host
	}

	// 读取调试模式
	if debug := os.Getenv("APP_DEBUG"); debug != "" {
		cfg.App.Debug = debug == "true" || debug == "1"
	}

	// 读取日志级别
	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		cfg.App.LogLevel = logLevel
	}

	return cfg
}

// GetServerAddr 获取服务器地址
// 返回格式：host:port
func (c *Config) GetServerAddr() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

// GetDatabaseDSN 获取数据库连接字符串
// 返回 MySQL 格式的 DSN
func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Database,
	)
}

// Print 打印配置信息（调试用）
func (c *Config) Print() {
	fmt.Println("========== 应用配置 ==========")
	fmt.Printf("应用名称: %s\n", c.App.Name)
	fmt.Printf("应用版本: %s\n", c.App.Version)
	fmt.Printf("调试模式: %v\n", c.App.Debug)
	fmt.Printf("日志级别: %s\n", c.App.LogLevel)
	fmt.Println("\n========== 服务器配置 ==========")
	fmt.Printf("监听地址: %s\n", c.GetServerAddr())
	fmt.Printf("读取超时: %d 秒\n", c.Server.ReadTimeout)
	fmt.Printf("写入超时: %d 秒\n", c.Server.WriteTimeout)
	fmt.Println("\n========== 数据库配置 ==========")
	fmt.Printf("数据库类型: %s\n", c.Database.Driver)
	fmt.Printf("数据库地址: %s:%d\n", c.Database.Host, c.Database.Port)
	fmt.Printf("数据库名称: %s\n", c.Database.Database)
	fmt.Println("==============================")
}

// ========== 使用示例 ==========

// 如何在 main.go 中使用配置：
/*
func main() {
    // 加载配置
    cfg := LoadConfig()

    // 打印配置（可选，用于调试）
    if cfg.App.Debug {
        cfg.Print()
    }

    // 使用配置创建服务器
    h := server.Default(server.WithHostPorts(cfg.GetServerAddr()))

    // ... 其他代码
}
*/

// 如何设置环境变量：
/*
Windows (CMD):
set SERVER_PORT=3000
set APP_DEBUG=true
go run .

Windows (PowerShell):
$env:SERVER_PORT="3000"
$env:APP_DEBUG="true"
go run .

Linux/Mac:
export SERVER_PORT=3000
export APP_DEBUG=true
go run .

或者在运行时设置：
SERVER_PORT=3000 APP_DEBUG=true go run .
*/
