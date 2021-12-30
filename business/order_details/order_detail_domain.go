package order_details

import (
	"context"
	"time"
)

type OrderDetailDomain struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	OrderId   string    `json:"order_id"`
	Quantity  int       `json:"quantity"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderDetailCreateDomain struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	OrderId  string `json:"order_id"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type OrderDetailUpdateDomain struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	OrderId  string `json:"order_id"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type OrderDetailUseCaseDomain interface {
	GetAllOrderDetails(ctx context.Context) ([]OrderDetailDomain, error)
	GetByOrderId(ctx context.Context, order_id int) ([]OrderDetailDomain, error)
	GetByOrderDetailId(ctx context.Context, id int) (OrderDetailDomain, error)
	CreateOrderDetail(ctx context.Context, createOrder OrderDetailCreateDomain) (OrderDetailDomain, error)
	UpdateOrderDetail(ctx context.Context, updateOrder OrderDetailUpdateDomain) (OrderDetailDomain, error)
	DeleteOrderDetail(ctx context.Context, id int) (OrderDetailDomain, error)
}

type OrderDetailRepository interface {
	GetAllOrderDetails(ctx context.Context) (res []OrderDetailDomain, err error)
	GetByOrderId(ctx context.Context, order_id int) (res []OrderDetailDomain, err error)
	GetByOrderDetailId(ctx context.Context, id int) (res OrderDetailDomain, err error)
	CreateOrderDetail(ctx context.Context, createOrder OrderDetailCreateDomain) (res OrderDetailDomain, err error)
	UpdateOrderDetail(ctx context.Context, update OrderDetailUpdateDomain) (res OrderDetailDomain, err error)
	DeleteOrderDetail(ctx context.Context, id int) (res OrderDetailDomain, err error)
}
