package query

import (
	"encoding/json"
	"fmt"
	"os"
	"usermanagement/inituser"
	"usermanagement/tool/add"
)

func Query() {
	user := inituser.UserList
	file, err := os.Open(add.FilePath)
	if err != nil {
		fmt.Println("可忽略该错误:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&user)

	if user == nil {
		fmt.Println("当前没有用户信息，添加一个呗！")
	} else {
		for _, v := range user {
			fmt.Printf("\nid: %v\n姓名: %v\n住址: %v\n联系: %v\n", v.Id, v.Name, v.Addr, v.Tel)
		}
	}
}
