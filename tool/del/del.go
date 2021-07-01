package del

import (
	"encoding/json"
	"fmt"
	"os"
	"usermanagement/inituser"
	"usermanagement/tool/add"
)

func Del() {
	user := inituser.UserList
	file, _ := os.Open(add.FilePath)
	defer file.Close()
	decode := json.NewDecoder(file)
	decode.Decode(&user)

	Y := ""
	fmt.Println("输入y/n确认是否删除 输入y删除用户信息！")
	fmt.Scanln(&Y)

	if Y == "y" {
		ID := -1
		fmt.Println("请输入需要删除的用户 ID！")
		fmt.Scanln(&ID)
		for i := 0; i < len(user); i++ {
			if ID == user[i].Id {
				ID = i
				user = append(user[:ID], user[ID+1:]...)
				file, _ = os.OpenFile(add.FilePath, os.O_TRUNC|os.O_WRONLY, 0666)
			}
		}
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(user)

}
