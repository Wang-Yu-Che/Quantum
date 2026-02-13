# Quantum

## 工程维度（目录结构设计）

```markdown
.
├── consumer            # 队列消费服务
├── go.mod              # Go 模块依赖文件
├── internal            # 工程内部可访问的公共模块
│   └── model           # GORM 数据库模型
├── job                 # cron job 服务
├── pkg                 # 工程外部可访问的公共模块
├── restful             # HTTP 服务目录，下存放以服务为维度的微服务
├── script              # 脚本服务目录，下存放以脚本为维度的服务
├── service             # gRPC 服务目录，下存放以服务为维度的微服务
└── Makefile            # 快速构建脚本
```

## 本地启动事项
1. 切换工作目录到所在目录
2. api设定etcd地址
3. 确保启动etcd(配置中心，可切换为consul和nacos) 
   1. 在 go-zero 中，支持 etcd 服务注册和直连模式，我们仅对 etc 目录下的静态配置文件稍作调整即可。
4. 启动api服务和rpc服务