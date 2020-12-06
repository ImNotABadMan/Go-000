package main

import (
	"Go-000/Week02/dao"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	homework()
	user := dao.User{}
	users := dao.GetUsers(&user)
	fmt.Println("这里应该是进行Wrap错误，比如：errors.Wrapf(code.NotFound, \\\"sql error: %v\\\", err)；\\\n")
	fmt.Println("在业务层再通过errors.Is(err, code.NotFound) 进行判断，而不是依赖底层错误")
	fmt.Printf("The cause is [%T], %v\n", errors.Cause(users), errors.Cause(users))
	fmt.Printf("Stack trace:\n%+v\n", users)

}

func homework() {
	fmt.Println("作业：我们在数据库操作的时候，",
		"比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？")
}
