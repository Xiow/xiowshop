package model

import "github.com/gogf/gf/v2/os/gtime"

// RotationCreateUpdateBase 创建/修改内容基类
type RotationCreateUpdateBase struct {
	PicUrl string // 轮播图地址
	Sort   int    // 排序
	Link   string // 跳转链接
}

// RotationCreateInput 创建内容
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建内容返回结果
type RotationCreateOutput struct {
	RotationId uint `json:"rotation_id"`
}

// RotationUpdateInput 修改内容
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}

// RotationGetListInput 获取内容列表
type RotationGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RotationGetListOutput 查询列表结果
type RotationGetListOutput struct {
	List  []RotationGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

//todo 讲讲这里 多级json
type RotationGetListOutputItem struct {
	//指定item的键名的方式
	//Rotation *RotationListItem `json:"Rotation"`
	//不指定item键名的方式
	Id        uint        `json:"id"`         // 自增ID
	PicUrl    string      `json:"pic_url"`    //图片链接
	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Link      string      `json:"brief"`      // 跳转链接
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
