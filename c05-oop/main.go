package main

import (
	"fmt"
	wallet "github.com/LostInAurora/go-design/c05-oop/encapsulation"
	"time"
)

func main() {
	vm := wallet.NewVirtualWallet()
	fmt.Println(vm.GetCreatedTime())
	time.Sleep(time.Second)
	_ = vm.IncBalance(5)
	_ = vm.DecBalance(2)
	fmt.Println(vm.GetId(), vm.GetBalance(), vm.GetCreatedTime(), vm.GetModifiedTime())
}