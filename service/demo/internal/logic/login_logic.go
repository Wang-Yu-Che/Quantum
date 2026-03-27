package logic

import (
	"context"
	"errors"
	"time"

	"Quantum/interface/demo/pb/demo"
	"Quantum/service/demo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(in *demo.UserLoginReq) (*demo.UserLoginResp, error) {
	// Mock 逻辑：验证用户名和密码
	// 实际开发中这里会查询数据库并校验密码哈希
	if in.Username == "admin" && in.Password == "123456" {
		return &demo.UserLoginResp{
			AccessToken: "mock-access-token-shard-123456",
			Expire:      time.Now().Add(time.Hour * 24).Unix(), // 24小时后过期
		}, nil
	}

	// 模拟账号密码错误
	return nil, errors.New("用户名或密码错误")
}
