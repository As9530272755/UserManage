package inituser

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Tel  string `json:"tel"`
	Addr string `json:"addr"`
}

var UserList []User

func Users(id int, name string, tel string, addr string) {

	users := []User{
		{
			id,
			name,
			tel,
			addr,
		},
	}
	UserList = append(UserList, users...)

}
