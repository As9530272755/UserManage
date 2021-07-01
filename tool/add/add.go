package add

import (
	"fmt"
	"usermanagement/inituser"
)

func Addr() {
	name := ""
	tel := ""
	addr := ""
	fmt.Println("请输入增加用户姓名：")
	fmt.Scanln(&name)

	fmt.Println("请输入增加用户联系方式：")
	fmt.Scanln(&tel)

	fmt.Println("请输入增加用户家庭住址：")
	fmt.Scanln(&addr)
	inituser.Users(len(inituser.UserList)+1, name, tel, addr)
}
