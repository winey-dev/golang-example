package data

import "fmt"

type user struct {
	Uid  string
	Name string
	Age  int
}

var users []user

func init() {
	users = append(users, user{Uid: "smlee", Name: "smlee", Age: 30})
	users = append(users, user{Uid: "oyryu", Name: "oyryu", Age: 30})
	users = append(users, user{Uid: "kjlee", Name: "kjlee", Age: 32})
	users = append(users, user{Uid: "ekkim", Name: "ekkim", Age: 30})
}

type uuser struct{}
type UserService interface {
	GetUser(id string) (user, error)
	GetUserList() ([]user, error)
}

func User() UserService {
	return new(uuser)
}

func (u *uuser) GetUser(id string) (user, error) {
	for _, o := range users {
		if o.Uid == id {
			return o, nil
		}
	}
	return user{}, fmt.Errorf("%s 를 찾을수 없습니다.", id)
}

func (u *uuser) GetUserList() ([]user, error) {
	if len(users) == 0 {
		return nil, fmt.Errorf("등록된 유저가 없습니다.")
	}
	return users, nil
}
