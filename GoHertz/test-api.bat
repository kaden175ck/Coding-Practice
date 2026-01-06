@echo off
chcp 65001 >nul
echo ========================================
echo    Hertz API 测试脚本 (Windows)
echo ========================================
echo.

echo [1] 测试欢迎页面
curl http://localhost:8080/
echo.
echo.

echo [2] 获取用户列表
curl http://localhost:8080/users
echo.
echo.

echo [3] 获取单个用户 (ID=1)
curl http://localhost:8080/users/1
echo.
echo.

echo [4] 创建新用户
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d "{\"name\":\"张三\",\"email\":\"zhangsan@test.com\",\"age\":25}"
echo.
echo.

echo [5] 更新用户 (ID=1)
curl -X PUT http://localhost:8080/users/1 -H "Content-Type: application/json" -d "{\"name\":\"张三丰\",\"age\":30}"
echo.
echo.

echo [6] 删除用户 (ID=1)
curl -X DELETE http://localhost:8080/users/1
echo.
echo.

echo [7] 搜索用户 (Query参数)
curl "http://localhost:8080/search?name=张三&age=25"
echo.
echo.

echo [8] 访问需要认证的接口 - 无Token (会失败)
curl http://localhost:8080/api/v1/products
echo.
echo.

echo [9] 访问需要认证的接口 - 有Token (会成功)
curl http://localhost:8080/api/v1/products -H "Authorization: Bearer demo-token-123456"
echo.
echo.

echo ========================================
echo    测试完成!
echo ========================================
pause
