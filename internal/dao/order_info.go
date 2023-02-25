// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gongzhaoweishop/internal/dao/internal"
)

// internalOrderInfoDao is internal type for wrapping internal DAO implements.
type internalOrderInfoDao = *internal.OrderInfoDao

// orderInfoDao is the data access object for table order_info.
// You can define custom methods on it to extend its functionality as you wish.
type orderInfoDao struct {
	internalOrderInfoDao
}

var (
	// OrderInfo is globally public accessible object for table order_info operations.
	OrderInfo = orderInfoDao{
		internal.NewOrderInfoDao(),
	}
)

// Fill with you ideas below.
