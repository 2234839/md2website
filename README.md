# md2website

将 MarkDown 文件转换为 html 的静态站点

[点击这里查看生成后的效果](https://2234839.github.io/md2website/) 静态文件位于 `./docs/`

## 运行方式 run

`go run .\src\main.go (sourceDir) (outDir)`

```go
go run .\src\main.go "C:\\Users\\llej\\AppData\\Local\\Programs\\SiYuan\\resources\\guide\\思源笔记用户指南" "D:\\code\\md2website\\docs"
```

## 待完成的功能点（按优先级降序排序）

| 可用 | 功能名 | 大致进度 |
| - | - | - |
| ❎🔨 | [#2 嵌入块渲染](https://github.com/2234839/md2website/issues/2) | `18%` |
| ⭕ | 菜单页面美化 |   |
| ⭕ | 页面 header 与 footer  |   |
| ⭕ | 块引用当前页面预览 |   |
| ⭕ | 块链接可 copy |   |
| ⭕ | 书签页 |   |
| ⭕ | 标签页 |   |
| ⭕ | 反链 |   |
| ✅🔨 | [#1 块引用链接](https://github.com/2234839/md2website/issues/1) | `90%` |
| ⭕ | 嵌入块查询渲染 |   |


1. ✅ 表示基本可以使用了
2. 🔨 表示还在修改中
2. ❎ 表示有一小部分功能可以使用了但还存在比较大的问题
3. ❌ 表示不可用
4. ⭕ 表示尚未开始
