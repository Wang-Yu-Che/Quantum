package g1

import (
	"Quantum/interface/greet/enum/user_enum"
	"Quantum/interface/hello/enum/hello_enum"
	"context"

	"Quantum/restful/demo/internal/svc"
	"Quantum/restful/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req *types.Request) (resp *types.Response, err error) {
	if user_enum.UserEnum_UserEnumUnknown == user_enum.UserEnum(hello_enum.ClientEnum_ClientEnumUnknown) {

	}
	return
}
