# Quantum

## 工程维度（目录结构设计）

```markdown
.
├── docs                # 文档说明
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

## k8s命令提示

### 查看所有pod
```shell
kubectl get pods -A
```

### 关闭服务
```shell
kubectl delete deployment 服务名
```

### 强制删除该 Pod
```shell
kubectl delete pod <pod名字> --grace-period=0 --force
```

### 查看 Ingress(对外开放接口)
```shell
kubectl get ingress
```

### 查看 Service
```shell
kubectl get svc
```

### 查看本地镜像
```shell
k3s crictl images
```

### 删除镜像
```shell
k3s crictl rmi 镜像名/镜像ID
```