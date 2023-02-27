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

// todo 完成删除功能
type RotationAddRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	RotationId uint `json:"rotation_id"`
}
type RotationDeleteReq struct {
	g.Meta `path:"/backend/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图接口"`
	Id     uint `v:"min:1#请选择需要删除的轮播图" dc:"轮播图id"`
}
type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/backend/rotation/update/{Id}" method:"post" tags:"轮播图" summary:"修改轮播图接口"`
	Id     uint   `json:"id"      v:"min:1#请选择需要修改的轮播图" dc:"轮播图Id"`
	PicUrl string `json:"pic_url" v:"required#轮播图图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"跳转链接"`
}
type RotationUpdateRes struct{}
type RotationGetListReq struct {
	g.Meta `path:"/backend/rotation/list" method:"get" tags:"轮播图" summary:"轮播图列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型 默认倒序"`
	CommonPaginationReq
}
type RotationGetListRes struct {
	Data interface{}
}
