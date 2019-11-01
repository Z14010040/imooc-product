package repositories

import (
	"database/sql"
	"imooc-product/common"
	"imooc-product/datamodels"
	"strconv"
)

type OrderRepository interface {
	Conn() error
	Insert(order *datamodels.Order) (int64, error)
	Delete(int642 int64) bool
	Update(order *datamodels.Order) error
	SelectByKey(int642 int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}
type OrderRepositoryManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewOrderRepositoryManager(table string, sql *sql.DB) OrderRepository {
	return &OrderRepositoryManager{table: table, mysqlConn: sql}
}
func (o *OrderRepositoryManager) Conn() error {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}
func (o *OrderRepositoryManager) Insert(order *datamodels.Order) (productId int64, err error) {
	if err = o.Conn(); err != nil {
		return
	}
	sql := "insert `order` set userID=?,productID=?,orderStatus=?"
	stmt, err := o.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	result, err := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		return
	}
	return result.LastInsertId()
}
func (o *OrderRepositoryManager) Delete(orderID int64) bool {
	if err := o.Conn(); err != nil {
		return false
	}
	sql := "delete from " + o.table + " where ID=" + strconv.FormatInt(orderID, 10)
	stmt, err := o.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return false
	}
	_, err = stmt.Exec(orderID)
	if err != nil {
		return false
	}
	return true
}
func (o *OrderRepositoryManager) Update(order *datamodels.Order) (err error) {
	if err = o.Conn(); err != nil {
		return
	}
	sql := "update " + o.table + " set userID=?, productID=?, orderStatus=? where ID=" + strconv.FormatInt(order.ID, 10)

	stmt, err := o.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	_, err = stmt.Exec(order.UserId, order.ProductId, order.OrderStatus, order.ID)
	return
}
func (o *OrderRepositoryManager) SelectByKey(orderID int64) (order *datamodels.Order, err error) {
	if err = o.Conn(); err != nil {
		return
	}
	sql := "select * from " + o.table + " where ID=" + strconv.FormatInt(orderID, 10)
	row, err := o.mysqlConn.Query(sql)
	defer row.Close()
	if err != nil {
		return
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return
	}
	order = &datamodels.Order{}
	common.DataToStructByTagSql(result, order)
	return
}
func (o *OrderRepositoryManager) SelectAll() (orders []*datamodels.Order, err error) {
	if err = o.Conn(); err != nil {
		return
	}
	sql := "select * from " + o.table
	rows, err := o.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return
	}
	for _, v := range result {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v, order)
		orders = append(orders, order)
	}
	return
}
func (o *OrderRepositoryManager) SelectAllWithInfo() (orderMap map[int]map[string]string, err error) {
	if err = o.Conn(); err != nil {
		return
	}
	sql := "select o.ID, p.productName, o.orderStatus from order as o left join on o.productID=p.ID"
	rows, err := o.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return
	}
	orderMap = common.GetResultRows(rows)
	return
}
