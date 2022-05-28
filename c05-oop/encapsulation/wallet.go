package wallet

import (
	"fmt"
	"github.com/pborman/uuid"
	"time"
)



type VirtualWallet struct {
	id string
	balance int
	createdTime int64
	modifiedTime int64
}

func NewVirtualWallet() *VirtualWallet {
	return &VirtualWallet{
		id: uuid.New(),
		createdTime: time.Now().Unix(),
		modifiedTime: time.Now().Unix(),
	}
}

func (v *VirtualWallet) GetId() string {
	return v.id
}

func (v *VirtualWallet) GetBalance() int {
	return v.balance
}

func (v *VirtualWallet) GetCreatedTime() int64 {
	return v.createdTime
}

func (v *VirtualWallet) GetModifiedTime() int64 {
	return v.modifiedTime
}

func (v *VirtualWallet) IncBalance(amount int) error {
	if amount < 0 {
		return fmt.Errorf("increase amount can not be minus 0")
	}
	v.balance += amount
	v.modifiedTime = time.Now().Unix()
	return nil
}

// DecBalance 数据库操作涉及并发问题
func (v *VirtualWallet) DecBalance(amount int) error {
	if amount < 0 {
		return fmt.Errorf("decrease amount can not be minus 0")
	}
	if v.balance - amount < 0 {
		return fmt.Errorf("balance is not enough")

	}
	v.balance -= amount
	v.modifiedTime = time.Now().Unix()
	return nil
}