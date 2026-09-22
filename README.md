# alan-content-platform-api

alan-content-platform 的后端 API 服务，基于 Go + Gin 构建。

## 开发

```bash
go run ./cmd/api
```

服务默认监听 `:8080`，可通过 `GET /healthz` 检查存活状态。

## 目录结构

```
cmd/api/          # 程序入口
internal/server/  # HTTP server 与路由注册
internal/handler/ # 请求处理器
```
