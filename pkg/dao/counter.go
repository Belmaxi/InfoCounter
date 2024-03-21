package dao

import (
	"InfoCounter/pkg/entity"
	"context"
)

// 操作数据库的接口
type CountNumDAO interface {
	//添加一个
	AddNumInfo(ctx context.Context, info entity.UserInfo) bool
}
