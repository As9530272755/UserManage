package del

import (
	"fmt"
	"usermanagement/inituser"
)

func Del() {
	Y := ""
	fmt.Println("输入y/n确认是否删除 输入y删除用户信息！")
	fmt.Scanln(&Y)

	if Y == "y" {
		var ID int
		fmt.Println("请输入需要删除的用户 ID！")
		fmt.Scanln(&ID)
		for _, v := range inituser.UserList {
			if ID == v.Id {
				inituser.UserList = append(inituser.UserList[:ID-1], inituser.UserList[ID:]...)
				fmt.Println("删除用户成功！")
			} else {
				fmt.Println("输入需要删除的用户 id 错误！")
			}
		}
	} else {
		fmt.Println("取消删除！")
	}

}
