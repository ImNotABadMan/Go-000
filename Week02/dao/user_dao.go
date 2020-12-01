package dao

import (
	"Go-000/Week02/mysql"
	"github.com/pkg/errors"
)

type User struct {
	id   int
	name string
}

func (user *User) Set(users []mysql.Data) error {
	for _, value := range users {
		switch t := value.(type) {
		case *User:
			user.id = t.id
			user.name = t.name
		}
	}
	return nil
}

func GetUsers(user *User) error {
	return errors.WithMessage(Db().GetRow(user, "select * from users"), "Dao GetUsers -- get none")
}

func Db() *mysql.MysqlDb {
	var dbConn *mysql.MysqlDb

	return func() *mysql.MysqlDb {
		if dbConn == nil {
			dbConn = mysql.Connect("127.0.0.1", 3306, "mysql", "123465", "test")
		}
		return dbConn
	}()
}
