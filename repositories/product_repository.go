package repositories

import (
	"database/sql"
	"imooc-product/common"
	"imooc-product/datamodels"
	"strconv"
)

type ProductRepository interface {
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int642 int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}
type ProductRepositoryManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductRepositoryManager(table string, mysqlConn *sql.DB) ProductRepository {
	return &ProductRepositoryManager{table: table, mysqlConn: mysqlConn}
}

//数据库连接
func (pm *ProductRepositoryManager) Conn() (err error) {
	if pm.mysqlConn != nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		pm.mysqlConn = mysql
	}
	if pm.table == "" {
		pm.table = "product"
	}
	return nil
}

//插入数据
func (pm *ProductRepositoryManager) Insert(product *datamodels.Product) (productId int64, err error) {
	//1、判断连接是否存在
	if err = pm.Conn(); err != nil {
		return
	}
	//2、准备sql
	sql := "INSERT product SET productName=?,productNum=?,productImage=?,productUrl=?"
	stmt, err := pm.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	//3、传入参数
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

//商品的删除
func (pm *ProductRepositoryManager) Delete(productId int64) bool {
	if err := pm.Conn(); err != nil {
		return false
	}
	sql := "delete from product where ID=?"

	stmt, err := pm.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return false
	}
	_, err = stmt.Exec(productId)
	if err != nil {
		return false
	}
	return true
}

//商品的更新
func (pm *ProductRepositoryManager) Update(product *datamodels.Product) (err error) {
	if err = pm.Conn(); err != nil {
		return
	}
	sql := "update product set productName=?,productNum=?,productImage=?,productUrl=? where ID=" + strconv.FormatInt(
		product.ID, 10)
	stmt, err := pm.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return
	}
	return
}

//查询单个商品
func (pm *ProductRepositoryManager) SelectByKey(productId int64) (productResult *datamodels.Product, err error) {
	if err = pm.Conn(); err != nil {
		return nil, err
	}
	sql := "select * from " + pm.table + " where ID=" + strconv.FormatInt(productId, 10)
	row, err := pm.mysqlConn.Query(sql)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return
	}
	productResult = &datamodels.Product{}
	common.DataToStructByTagSql(result, productResult)
	return
}

//获取所有的商品
func (pm *ProductRepositoryManager) SelectAll() (productList []*datamodels.Product, err error) {
	if err = pm.Conn(); err != nil {
		return
	}
	sql := "select * from " + pm.table
	rows, err := pm.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return
	}
	defer rows.Close()
	result := common.GetResultRows(rows)
	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		productList = append(productList, product)
	}
	return
}
