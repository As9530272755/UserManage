package add

import (
	"encoding/json"
	"fmt"
	"os"
	"usermanagement/inituser"
)

var FilePath string = "user.json"

func Addr() {
	file, err := os.Open(FilePath)
	if err != nil {
		file, _ = os.Create(FilePath)
	} else {
		decode := json.NewDecoder(file)
		decode.Decode(&inituser.UserList)
		file, _ = os.OpenFile(FilePath, os.O_APPEND|os.O_WRONLY|os.O_TRUNC, 0666)
	}
	defer file.Close()

	name := ""
	tel := ""
	addr := ""
	fmt.Println("请输入增加用户姓名：")
	fmt.Scanln(&name)

	fmt.Println("请输入增加用户联系方式：")
	fmt.Scanln(&tel)

	fmt.Println("请输入增加用户家庭住址：")
	fmt.Scanln(&addr)

	id := 0
	for _, v := range inituser.UserList {
		id = v.Id
	}

	inituser.Users(id+1, name, tel, addr)

	encoder := json.NewEncoder(file)
	encoder.Encode(inituser.UserList)

}
