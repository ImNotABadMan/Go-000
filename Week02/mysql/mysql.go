package mysql

import (
	"fmt"
	"github.com/pkg/errors"
)

type MysqlDb struct {
	host     string
	port     int
	username string
	password string
	database string
}

// sqlError
type ErrNoRows struct {
	sql string
}

// dao层实现接口
type Data interface {
	Set([]Data) error
}

func Connect(host string, port int, username string, password string, database string) *MysqlDb {
	return &MysqlDb{
		host:     host,
		port:     port,
		username: username,
		password: password,
		database: database,
	}
}

// error no rows实现error接口
func (err *ErrNoRows) Error() string {
	return fmt.Sprint("Mysql "+err.sql, " -- SQL get no rows")
}

// 获取数据库记录
func (mysql *MysqlDb) GetRow(data Data, sqlStr string) error {
	rows, err := query(sqlStr)
	if err != nil {
		return errors.Wrap(err, "main GetRow -- get none")

	}

	return data.Set(rows)
}

func query(sql string) ([]Data, error) {
	return []Data{}, &ErrNoRows{sql: sql}
}
