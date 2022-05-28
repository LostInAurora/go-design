package user

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Formatter struct {

}

func (fm *Formatter) Format(src, dst string) {
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
		Users = append(Users, *parseToUser(fileScanner.Text()))
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	// 2. sort User by age
	sortByAge(Users)

	// 3.
	f, err := os.Create(dst)
	if err != nil {
		log.Fatalf("Error when create file: %s", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, user := range Users {
		fmt.Fprintln(w, user.userToString())
		w.Flush()
	}
}
