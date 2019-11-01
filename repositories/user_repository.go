package repositories

import (
	"database/sql"
	"errors"
	"imooc-product/common"
	"imooc-product/datamodels"
	"strconv"
)

type UserRepository interface {
	Conn() error
	Select(userName string) (*datamodels.User, error)
	Insert(user *datamodels.User) (int64, error)
}
type UserRepositoryManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewUserRepositoryManager(table string, mysqlConn *sql.DB) UserRepository {
	return &UserRepositoryManager{table: table, mysqlConn: mysqlConn}
}
func (u *UserRepositoryManager) Conn() error {
	if u.mysqlConn != nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		u.mysqlConn = mysql
	}
	if u.table == "" {
		u.table = "user"
	}
	return nil
}
func (u *UserRepositoryManager) Select(userName string) (user *datamodels.User, err error) {
	if err = u.Conn(); err != nil {
		return
	}
	sql := "select * from user where userName=?"
	row, err := u.mysqlConn.Query(sql, userName)
	defer row.Close()
	if err != nil {
		return
	}
	user = &datamodels.User{}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return
	}
	common.DataToStructByTagSql(result, user)
	return
}
func (u *UserRepositoryManager) Insert(user *datamodels.User) (userID int64, err error) {
	if err = u.Conn(); err != nil {
		return
	}
	sql := "insert user set nickName=?,userName=?,password=?"

	stmt, err := u.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	result, err := stmt.Exec(user.NickName, user.UserName, user.HashPassword)
	if err != nil {
		return
	}
	return result.LastInsertId()
}
func (u *UserRepositoryManager) SelectByID(userId int64) (user *datamodels.User, err error) {
	if err = u.Conn(); err != nil {
		return &datamodels.User{}, err
	}
	sql := "select * from " + u.table + " where ID=" + strconv.FormatInt(userId, 10)
	row, errRow := u.mysqlConn.Query(sql)
	defer row.Close()
	if errRow != nil {
		return &datamodels.User{}, errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.User{}, errors.New("用户不存在！")
	}
	user = &datamodels.User{}
	common.DataToStructByTagSql(result, user)
	return
}
