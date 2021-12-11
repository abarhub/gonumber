package main

import (
	"fmt"
	"time"
)

func isPrime(n int64) bool {
	var i int64
	for i = 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func listPrime(n int64, affichage bool) {
	var i int64
	for i = 2; i <= n; i++ {
		if isPrime(i) {
			if affichage {
				println(i)
			}
		}
	}
}

func main() {
	var n int64
	var affiche bool

	n = 100
	n = 50000
	n = 100000
	n = 1_000_000

	affiche = false
	affiche = true

	debut := time.Now().UnixNano() / 1000_000
	listPrime(n, affiche)
	fin := time.Now().UnixNano() / 1000_000
	fmt.Printf("%d msec", fin-debut)
}
