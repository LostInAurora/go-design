package main

import "fmt"

type Flyable interface {
	fly()
}
// 接口
type Eggable interface {
	egg()
}

type FlyAbility struct{}

func (f *FlyAbility) fly() {
	fmt.Println("I can fly")
}

type EggAbility struct{}

func (e *EggAbility) egg() {
	fmt.Println("I can egg")
}

type Eager struct {
	f FlyAbility //组合
	e EggAbility //组合
}

func (eager *Eager) egg() {
	eager.e.egg() // 委托
}

func (eager *Eager) fly() {
	eager.f.fly() // 委托
}



func main() {
	eager := new(Eager)
	eager.egg()
	eager.fly()
}

