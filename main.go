package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func affichageListePrime(n int64, affiche bool) {

	listPrime(n, affiche)
}

func enregistreListePrime(n int64, chemin string) {

	f, err := os.Create(chemin)
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	var i int64
	for i = 2; i <= n; i++ {
		if isPrime(i) {

			s := strconv.FormatInt(i, 10)
			//println(s)
			_, err := w.WriteString(s + "\n")
			check(err)
		}
	}
	err = w.Flush()
	check(err)

}

func enregistreListeDoublePrime(chemin string, cheminDest string) {

	f, err := os.Open(chemin)
	check(err)
	defer f.Close()

	f2, err2 := os.Create(cheminDest)
	check(err2)
	defer f2.Close()

	var buf []int64

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		check(scanner.Err())
		//fmt.Println(scanner.Text())
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			check(err)
		}
		buf = append(buf, i)
	}

	w := bufio.NewWriter(f2)
	//var i int64
	for _, v := range buf {
		for _, v2 := range buf {
			n := v * v2
			s := fmt.Sprintf("%d*%d=%d", v, v2, n)
			_, err := w.WriteString(s + "\n")
			check(err)
		}
	}
	//for i = 2; i <= n; i++ {
	//	if isPrime(i) {
	//
	//		s := strconv.FormatInt(i, 10)
	//		//println(s)
	//		_, err := w.WriteString(s + "\n")
	//		check(err)
	//	}
	//}
	err = w.Flush()
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var n int64
	//var affiche bool

	n = 100
	n = 50000
	n = 100000
	n = 1_000_000
	//n = 1_000_000_000

	//affiche = false
	//affiche = true
	methode := 3

	if methode == 1 {
		debut := time.Now().UnixNano() / 1000_000
		affichageListePrime(n, true)
		fin := time.Now().UnixNano() / 1000_000
		fmt.Printf("%d msec", fin-debut)
	} else if methode == 2 {
		filename := "d:\\temp\\primes.txt"
		fmt.Printf("Enregistrement du fichier %s\n", filename)
		debut := time.Now().UnixNano() / 1000_000
		enregistreListePrime(n, filename)
		fin := time.Now().UnixNano() / 1000_000
		fmt.Printf("%d msec\n", fin-debut)
	} else if methode == 3 {
		filename := "data/primes_1_000_000.txt"
		filename2 := "data/double_primes.txt"
		fmt.Printf("Enregistrement du fichier %s\n", filename)
		debut := time.Now().UnixNano() / 1000_000
		enregistreListeDoublePrime(filename, filename2)
		fin := time.Now().UnixNano() / 1000_000
		fmt.Printf("%d msec\n", fin-debut)
	}
}
