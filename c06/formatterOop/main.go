package main

import (
	"github.com/LostInAurora/go-design/c06/formatterOop/user"
)

func main() {
	fm := new(user.Formatter)
	fm.Format("test.txt", "output.txt")
}