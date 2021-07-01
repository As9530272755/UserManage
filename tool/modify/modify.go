package modify

import (
	"fmt"
	"usermanagement/inituser"
)

func Modify() {
	ID := 0
	fmt.Println("请输入需要修改的用户，输入 id 即可：")
	fmt.Scanln(&ID)
	for k, v := range inituser.UserList {
		if ID == v.Id {
			fmt.Printf("\nid: %v\n姓名: %v\n住址: %v\n联系: %v\n", v.Id, v.Name, v.Addr, v.Tel)
		}
		y := ""
		fmt.Println("输入y/n确认是否修改 输入y修改用户信息")
		fmt.Scanln(&y)
		if y == "y" {
			name := ""
			tel := ""
			addr := ""
			fmt.Println("请输入修改后的用户姓名：")
			fmt.Scanln(&name)

			fmt.Println("请输入修改后的用户联系方式：")
			fmt.Scanln(&tel)

			fmt.Println("请输入修改后的用户家庭住址：")
			fmt.Scanln(&addr)
			inituser.UserList[k].Name = name
			inituser.UserList[k].Addr = addr
			inituser.UserList[k].Tel = tel
			return
		}
	}

}
