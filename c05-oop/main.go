package main

import "fmt"

// 多态
type Shower interface {
	Print()
}
// 封装
// producedTime can not be accessed by other packages directly
type Item struct {
	Price string
	producedTime string
}
func(i *Item) Print() {
	fmt.Println("price", i.Price, "producedTime", i.producedTime)
}

func(i Item) Iso() {
	fmt.Println("独立方法")
}
// 继承, 貌似也不是继承～
type Wallet struct{
	Item
	Brand string
}

func(w *Wallet) Print() {
	fmt.Println("price", w.Price, "producedTime", w.producedTime, "brand")
}
func main() {
	var s Shower
	s = new(Item)
	s.Print()
	s = new(Wallet)
	s.Print()
	w := new(Wallet)
	w.Item.Iso()
}