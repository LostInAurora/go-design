package coupon

import (
	"github.com/LostInAurora/go-design/coupon/kafka"
	"github.com/LostInAurora/go-design/coupon/orderHandler"
)

type Coupon struct {
	kafka.KafkaRecv
	orderHandler.CouponDispatcher
}