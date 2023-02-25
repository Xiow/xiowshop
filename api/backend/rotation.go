package backend

import "github.com/gogf/gf/v2/frame/g"

//required 代表该字段时候为必填字段,如果请求不带该字段,返回报错根据#后面进行返回
//dc descrition 描述
type RotationAddReq struct {
	g.Meta `path:"/backend/rotation/add" tags:"Rotation" method:"post" summary:"添加轮播图"`
	PicUrl string `json:"pic_url"    v:"required#轮播图不能为空" dc:"轮播图"`
	Link   string `json:"link"    v:"required#跳转链接不能为空"`
	Sort   int    `json:"sort"`
}

// todo
type RotationAddRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	RotationId uint `json:"rotation_id"`
}
