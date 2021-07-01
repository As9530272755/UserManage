package menu

import (
	"fmt"
	"usermanagement/tool/add"
	"usermanagement/tool/del"
	"usermanagement/tool/modify"
	"usermanagement/tool/query"
)

func Menu() {
	// add.Addr()
	// query.Query()
	// del.Del()
	key := ""
	for {
		fmt.Println("欢迎使用用户管理系统！\n请输入使用功能：\n1.查询\n2.添加用户\n3.删除用户\n4.修改用户\n5.退出程序")
		fmt.Println("请输入对应的编号实现对应功能：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			query.Query()
			fmt.Println("\n")
		case "2":
			add.Addr()
			fmt.Println("\n")
		case "3":
			del.Del()
			fmt.Println("\n")
		case "4":
			modify.Modify()
			fmt.Println("\n")
		case "5":
			fmt.Println("欢迎再次使用用户管理系统！")
			return
		}
	}
}
