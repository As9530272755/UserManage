package menu

import (
	"crypto/md5"
	"fmt"
	"usermanagement/tool/add"
	"usermanagement/tool/del"
	"usermanagement/tool/modify"
	"usermanagement/tool/query"

	"github.com/spf13/cobra"
)

func Menu() {
	n := 3

	for i := 0; i < n; i++ {
		pwd := ""
		fmt.Println("请输入密码登录用户管理系统：")
		fmt.Printf("当前登录次数：%v\n", n-i)
		fmt.Scanln(&pwd)
		NewMd5 := md5.Sum([]byte("666"))
		PawMd5 := md5.Sum([]byte(pwd))
		if PawMd5 == NewMd5 {
			fmt.Println("登录成功!!\n")
			// 定义 q 命令
			var cmd1 = &cobra.Command{
				Use:   "q",
				Short: "query user info!",
				Run: func(cmd *cobra.Command, args []string) {
					query.Query()
				},
			}

			// 定义 a 命令
			var cmd2 = &cobra.Command{
				Use:   "a",
				Short: "add user info!",
				Run: func(cmd *cobra.Command, args []string) {
					add.Addr()
				},
			}

			// 定义 d 命令
			var cmd3 = &cobra.Command{
				Use:   "d",
				Short: "delete user",
				Run: func(cmd *cobra.Command, args []string) {
					del.Del()
				},
			}

			// 定义 m 命令
			var cmd4 = &cobra.Command{
				Use:   "m",
				Short: "modify user info!",
				Run: func(cmd *cobra.Command, args []string) {
					modify.Modify()
				},
			}

			// 定义根命令
			var rootCmd1 = &cobra.Command{Use: "user"}

			// 加入子命令
			rootCmd1.AddCommand(cmd1, cmd2, cmd3, cmd4)

			// 初始化cobra
			rootCmd1.Execute()
			break
		} else {
			fmt.Println("输入密码错误!\n")
		}
	}

}

// fmt.Println("欢迎使用用户管理系统！\n请输入使用功能：\n1.查询\n2.添加用户\n3.删除用户\n4.修改用户\n5.退出程序")
// fmt.Println("请输入对应的编号实现对应功能：")
