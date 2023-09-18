package main

import (
	"fmt"
	"sync"
	"time"
)

type updateBalance struct {
	mx      sync.Mutex
	balance int
}

func (u *updateBalance) deposit(wg *sync.WaitGroup, amount int) {
	u.mx.Lock()
	defer u.mx.Unlock()
	defer wg.Done()
	u.balance += amount
	time.Sleep(time.Second)
}

func (u *updateBalance) getBalance(wg *sync.WaitGroup) {
	u.mx.Lock()
	defer u.mx.Unlock()
	defer wg.Done()
	fmt.Printf("Current balance is %d \n", u.balance)
}

func main() {
	var wg sync.WaitGroup
	u := updateBalance{balance: 0}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go u.deposit(&wg, 1000)
	}
	go u.getBalance(&wg)
	wg.Wait()
	fmt.Println("Done")
}
