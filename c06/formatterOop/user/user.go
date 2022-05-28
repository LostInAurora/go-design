package user

import (
	"sort"
	"strconv"
	"strings"
)

type User struct {
	Name string
	Age int
	Gender string
}

func (user *User) userToString() string {
	userInfoRaw := []string{
		user.Name,
		strconv.Itoa(user.Age),
		user.Gender,
	}
	return strings.Join(userInfoRaw, "\t")
}


func parseToUser(str string) *User {
	userInfoRaw := strings.Split(str, "&")
	age, _ := strconv.Atoi(userInfoRaw[1])
	return &User{
		Name: userInfoRaw[0],
		Age: age,
		Gender: userInfoRaw[2],
	}
}

func sortByAge(users [] User) []User {
	sort.Slice(users, func(i ,j int)bool {
		return users[i].Age < users[j].Age
	})
	return users
}