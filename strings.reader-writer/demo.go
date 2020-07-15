package main

import (
	"fmt"
	"sync"
	"time"
)

type UserTable struct {
	sync.RWMutex
	//deadlock.RWMutex

	table map[int]string
}

func NewUserTable() *UserTable {
	return &UserTable{
		table: make(map[int]string),
	}
}

func (userTable *UserTable) GetUser(i int) string {
	userTable.RLock()
	defer userTable.RUnlock()
	if name, exist := userTable.table[i]; exist {
		return name
	}
	return ""
}

// 并发安全操作 map
func (userTable *UserTable) GetTable() map[int]string {
	userTable.RLock()
	defer userTable.RUnlock()
	newMap := make(map[int]string)
	for seatID := range userTable.table {
		newMap[seatID] = userTable.GetUser(seatID)
	}
	return newMap
}

func (userTable *UserTable) SetUser(i int, name string) {
	userTable.Lock()
	defer userTable.Unlock()
	if len(userTable.table) >= 100 {
		userTable.table = make(map[int]string)
	}
	userTable.table[i] = name
}

func main() {
	userTable := NewUserTable()
	go func() {
		i := 0
		for {
			if i%1000 == 0 {
				fmt.Printf("%v\n", userTable.GetTable())
			}
			userTable.SetUser(i, fmt.Sprintf("%v", i))
			i++
		}
	}()

	go func() {
		for {
			userTable.GetTable()
			time.Sleep(100 * time.Millisecond)
		}

	}()

	//select {}

	time.Sleep(2 * time.Second)
}
