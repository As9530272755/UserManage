package query

import (
	"fmt"
	"usermanagement/inituser"
)

func Query() {
	if inituser.UserList == nil {
		fmt.Println("当前没有用户信息，添加一个呗！")
	} else {
		for _, v := range inituser.UserList {
			fmt.Printf("\nid: %v\n姓名: %v\n住址: %v\n联系: %v\n", v.Id, v.Name, v.Addr, v.Tel)
		}
	}

}
