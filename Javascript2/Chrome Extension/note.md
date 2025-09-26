## 笔记：Wikipedia LocalStorage 测试与插件数据保存

### 1. 进入 LocalStorage

1. 打开 **Wikipedia** 网站
2. 按 `F12` 打开 **Inspect (开发者工具)**
3. 切换到 **Application** 选项卡
4. 在左侧找到 **LocalStorage** 并选择对应网址

---

### 2. 清除默认数据

在 **Console** 输入：

```javascript
localStorage.clear();
```

* 作用：清空当前 LocalStorage 中的所有数据
* 清除后，Wikipedia 默认的 **3 个数据** 会消失

---

### 3. 添加/修改 LocalStorage 数据

可以手动在 Console 中添加或修改数据，例如：

```javascript
localStorage.setItem("key", "value"); // 添加数据
localStorage.getItem("key"); // 读取数据
localStorage.removeItem("key"); // 删除某个数据
```

* 这些数据会保存在浏览器 LocalStorage 中
* 即使刷新页面，LocalStorage 的数据依然存在
* 但如果插件没有做保存功能，退出插件时数据不会保留

---

### 4. 默认数据重新出现的原因

* 刚才被删掉的 **3 个默认数据** 会再次出现
* 原因：**Wikipedia 会在加载时重新写入它需要的默认数据**
* 你的假数据依然存在，只是被默认数据“补上”

---

### 5. 插件数据丢失问题

* 当前插件 **没有保存到 LocalStorage**
* 退出插件后，所有用户数据会丢失
* 解决方法：在插件的 **JavaScript 文件** 中加入保存功能


### 6. manifest 配置文件中，一定要注意拼写问题！！！！！如果没有任何报错，但是API后render出undefined，请检查配置是否拼写错误！！！
