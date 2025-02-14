package exercises

import (
	"encoding/json"
	"fmt"
)

type User struct{
	Name string `json:"name"`
	Id uint64 `json:"id"`
}

func (u *User) UpdateName(name string){
	u.Name = name
}

func MakeUser(name string, id uint64) {
	user := User{Name: name, Id: id}
	fmt.Println(user)
	user.UpdateName("Jeremias")
	fmt.Println(user)

	res, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}