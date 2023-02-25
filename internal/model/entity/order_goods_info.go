// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderGoodsInfo is the golang structure for table order_goods_info.
type OrderGoodsInfo struct {
	Id             int         `json:"id"             description:"商品维度的订单表"`
	OrderId        int         `json:"orderId"        description:"关联的主订单表"`
	GoodsId        int         `json:"goodsId"        description:"商品id"`
	GoodsOptionsId int         `json:"goodsOptionsId" description:"商品规格id sku id"`
	Count          int         `json:"count"          description:"商品数量"`
	Remark         string      `json:"remark"         description:"备注"`
	Price          int         `json:"price"          description:"订单金额 单位分"`
	CouponPrice    int         `json:"couponPrice"    description:"优惠券金额 单位分"`
	ActualPrice    int         `json:"actualPrice"    description:"实际支付金额 单位分"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:""`
}
