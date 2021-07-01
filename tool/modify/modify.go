package modify

import (
	"encoding/json"
	"fmt"
	"os"
	"usermanagement/inituser"
	"usermanagement/tool/add"
)

func Modify() {
	user := inituser.UserList
	file, _ := os.Open(add.FilePath)
	defer file.Close()

	decode := json.NewDecoder(file)
	decode.Decode(&user)

	ID := 0
	fmt.Println("请输入需要修改的用户，输入 id 即可：")
	fmt.Scanln(&ID)
	for i := 0; i < len(user); i++ {
		if ID == user[i].Id {
			fmt.Printf("\nid: %v\n姓名: %v\n住址: %v\n联系: %v\n", user[i].Id, user[i].Name, user[i].Addr, user[i].Tel)
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
				user[i].Name = name
				user[i].Addr = addr
				user[i].Tel = tel
				Newfile, _ := os.OpenFile(add.FilePath, os.O_RDWR|os.O_TRUNC, 0666)
				encode := json.NewEncoder(Newfile)
				encode.Encode(user)
				return
			}
		}
	}
}
