package services

import (
	"imooc-product/datamodels"
	"imooc-product/repositories"
)

type ProductService interface{
	GetProductByID(int642 int64)(*datamodels.Product,error)
	GetAllProduct()([]*datamodels.Product,error)
	DeleteProductByID(int642 int64)(bool)
	InsertProduct(product *datamodels.Product)(int64,error)
	UpdateProduct(product *datamodels.Product)(error)
}
type ProductServiceManager struct{
	ProductRepository repositories.ProductRepository
}
//初始化函数
func NewProductServiceManager(productRepository repositories.ProductRepository)(ProductService) {
	return &ProductServiceManager{ProductRepository:productRepository}
}
func(pm *ProductServiceManager)GetProductByID(int642 int64)(*datamodels.Product,error){
	return pm.ProductRepository.SelectByKey(int642)
}
func(pm *ProductServiceManager)GetAllProduct()([]*datamodels.Product,error){
	return pm.ProductRepository.SelectAll()
}
func(pm *ProductServiceManager)DeleteProductByID(int642 int64)(bool){
	return pm.ProductRepository.Delete(int642)
}
func(pm *ProductServiceManager)InsertProduct(product *datamodels.Product)(int64,error){
	return pm.ProductRepository.Insert(product)
}
func (pm *ProductServiceManager) UpdateProduct(product *datamodels.Product)(error) {
	return pm.ProductRepository.Update(product)
}