package services

import (
	"imooc-product/datamodels"
	"imooc-product/repositories"
)

type OrderService interface {
	GetOrderByID(int642 int64) (*datamodels.Order, error)
	DeleteOrderByID(int642 int64) bool
	UpdateOrder(order *datamodels.Order) error
	InsertOrder(order *datamodels.Order) (int64, error)
	GetAllOrder() ([]*datamodels.Order, error)
	GetAllOrderInfo() (map[int]map[string]string, error)
}
type OrderServiceManage struct {
	OrderRepository repositories.OrderRepository
}

func NewOrderServiceManager(repository repositories.OrderRepository) OrderService {
	return &OrderServiceManage{OrderRepository: repository}
}
func (o *OrderServiceManage) GetOrderByID(int642 int64) (order *datamodels.Order, err error) {
	return o.OrderRepository.SelectByKey(int642)
}
func (o *OrderServiceManage) DeleteOrderByID(int642 int64) bool {
	isOk := o.OrderRepository.Delete(int642)
	return isOk
}
func (o *OrderServiceManage) UpdateOrder(order *datamodels.Order) error {
	err := o.OrderRepository.Update(order)
	return err
}
func (o *OrderServiceManage) InsertOrder(order *datamodels.Order) (int64, error) {
	id, err := o.OrderRepository.Insert(order)
	return id, err
}
func (o *OrderServiceManage) GetAllOrder() ([]*datamodels.Order, error) {
	orders, err := o.OrderRepository.SelectAll()
	return orders, err
}
func (o *OrderServiceManage) GetAllOrderInfo() (map[int]map[string]string, error) {
	ordersInfo, err := o.OrderRepository.SelectAllWithInfo()
	return ordersInfo, err
}
