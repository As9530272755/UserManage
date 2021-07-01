package inituser

type User struct {
	Id   int
	Name string
	Tel  string
	Addr string
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
