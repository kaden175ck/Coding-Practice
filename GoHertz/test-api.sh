#!/bin/bash

echo "========================================"
echo "   Hertz API 测试脚本 (Linux/Mac)"
echo "========================================"
echo

echo "[1] 测试欢迎页面"
curl http://localhost:8080/
echo -e "\n"

echo "[2] 获取用户列表"
curl http://localhost:8080/users
echo -e "\n"

echo "[3] 获取单个用户 (ID=1)"
curl http://localhost:8080/users/1
echo -e "\n"

echo "[4] 创建新用户"
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","email":"zhangsan@test.com","age":25}'
echo -e "\n"

echo "[5] 更新用户 (ID=1)"
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"张三丰","age":30}'
echo -e "\n"

echo "[6] 删除用户 (ID=1)"
curl -X DELETE http://localhost:8080/users/1
echo -e "\n"

echo "[7] 搜索用户 (Query参数)"
curl "http://localhost:8080/search?name=张三&age=25"
echo -e "\n"

echo "[8] 访问需要认证的接口 - 无Token (会失败)"
curl http://localhost:8080/api/v1/products
echo -e "\n"

echo "[9] 访问需要认证的接口 - 有Token (会成功)"
curl http://localhost:8080/api/v1/products \
  -H "Authorization: Bearer demo-token-123456"
echo -e "\n"

echo "========================================"
echo "   测试完成!"
echo "========================================"
