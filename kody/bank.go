package main

import (
	"sync"
)

// var balance int
// func Deposit(amount int) { balance = balance + amount }
// func Balance() int { return balance }

// Ania
// go func(){
// 	bank.Deposit(200) // A1
// 	fmt.Println("=", bank.Balance()) //A2
// }
// go bank.Deposit(100) //B

// najpierw Ania			najpierwa Robert			Ania-Robert-Ania
// A1		200				B		100					A1 			200
// A2		200				A1		300					B			300
// B		300				A2		300					A2			300

// Wyscig:
// A1r	 	0
// B 		100 balance + amount
// A1w		200 balance = ....
// A2  	200

// var x []int
// go func() { x = make([]int, 10) }()
// go func() { x = make([]int, 1000000) }()
// x[999999] = 1 // UWAGA: niezdefiniowane zachowanie; możliwe uszkodzenie pamięci!

// var icons map[string]image.Image

// func loadIcon(){
// 	icons = make(map[string]image.Image)
// 	icons["spades.png"]= loadIcon("spades.png"),
// 	"hearts.png": loadIcon("hearts.png"),
// 	"diamonds.png": loadIcon("diamonds.png"),
// 	"clubs.png": loadIcon("clubs.png"),
// }

// // UWAGA: to nie jest współbieżnie bezpieczne!
// func Icon(name string) image.Image {
// 	icon, ok := icons[name]
// 	if !ok {
// 		icon = loadIcon(name)
// 		icons[name] = icon
// 	}
// 	return icon
// }

var mu sync.RWMutex
var deposits = make(chan int) // wysyłanie kwoty do wpłaty
var balances = make(chan int) // odbieranie salda
var balance int

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if Balance() < 0 {
		deposit(amount)
		return false // niewystarczające środki
	}
	return true
}

func deposit(amount int) { balance += amount }

func teller() {
	var balance int // zmienna balance jest zamknięta w funkcji goroutine teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // uruchomienie monitorującej funkcji goroutine
}
