package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	formatUser("test.txt", "output.txt")


}

type User struct {
	Name string
	Age int
	Gender string
}

func parseToUser(str string) User {
	userInfoRaw := strings.Split(str, "&")
	age, _ := strconv.Atoi(userInfoRaw[1])
	return User{
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

func userToString(user User) string {

	userInfoRaw := []string{
		user.Name,
		strconv.Itoa(user.Age),
		user.Gender,
	}
	return strings.Join(userInfoRaw, "\t")
}

func formatUser(src, dst string) {
	// 1. parse file to Users
	file, err := os.Open(src)
	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	// defer should be after return
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	// read line by line
	Users := make([]User, 0)
	for fileScanner.Scan() {
		Users = append(Users, parseToUser(fileScanner.Text()))
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	// 2. sort User by age
	sortByAge(Users)

	// 3. write user to file
	f, err := os.Create(dst)
	if err != nil {
		log.Fatalf("Error when create file: %s", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, user := range Users {
		fmt.Fprintln(w, userToString(user))
		w.Flush()
	}

}