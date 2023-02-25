package controller

import (
	"context"
	"gongzhaoweishop/api/backend"
	"gongzhaoweishop/internal/model"
	"gongzhaoweishop/internal/service"
)

var (
	Rotation = cRotation{}
)

type cRotation struct{}

func (a *cRotation) Create(ctx context.Context, req *backend.RotationAddReq) (res *backend.RotationAddRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationAddRes{RotationId: out.RotationId}, nil
}
