package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
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
	var j int64 = 0
	for i = 2; i <= n; i++ {
		if isPrime(i) {

			s := strconv.FormatInt(i, 10)
			//println(s)
			_, err := w.WriteString(s + "\n")
			check(err)
			j++
			if j%10 == 0 {
				err = w.Flush()
				check(err)
			}
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
	for _, v := range buf {
		for _, v2 := range buf {
			n := v * v2
			s := fmt.Sprintf("%d*%d=%d", v, v2, n)
			_, err := w.WriteString(s + "\n")
			check(err)
		}
	}
	err = w.Flush()
	check(err)

}

// IntPow calculates n to the mth power. Since the result is an int, it is assumed that m is a positive power
func IntPow(n int64, m int) int64 {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func filtreListeDoublePrime(n int, chemin string, cheminDest string) {

	f, err := os.Open(chemin)
	check(err)
	defer f.Close()

	f2, err2 := os.Create(cheminDest)
	check(err2)
	defer f2.Close()

	//var buf []int64

	scanner := bufio.NewScanner(f)
	w := bufio.NewWriter(f2)
	for scanner.Scan() {
		check(scanner.Err())
		s := scanner.Text()
		pos := strings.LastIndex(s, "=")
		if pos >= 0 {
			i, err := strconv.ParseInt(s[pos+1:], 10, 64)
			if err != nil {
				check(err)
			}
			if i >= IntPow(10, n-1) && i < IntPow(10, n) {
				_, err := w.WriteString(s + "\n")
				check(err)
			}
		}

	}

	err = w.Flush()
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func racineCarre(n *big.Int) *big.Int {
	var racineCarre, tmp *big.Int
	racineCarre = big.NewInt(0)
	tmp = big.NewInt(0)
	racineCarre.Sqrt(n)
	tmp.Mul(racineCarre, racineCarre)
	if tmp.Cmp(n) == 0 {
		return racineCarre
	} else {
		return big.NewInt(0)
	}
}

func factorisationFermat(n *big.Int) {

	var a0, b0 *big.Int
	trouve := false
	zero := big.NewInt(0)
	un := big.NewInt(1)
	deux := big.NewInt(2)
	n2 := n
	var max *big.Int = big.NewInt(0)
	max.Add(n2, un)
	tmp := big.NewInt(0)
	var tmp2 *big.Int = big.NewInt(0)
	for i := big.NewInt(0); i.Cmp(max) <= 0; {

		a := i
		tmp.Exp(a, deux, nil)
		tmp2.Sub(tmp, n)

		if tmp2.Cmp(zero) < 0 {
			// tmp2 <0 => on passe au suivant
		} else if tmp2.Cmp(zero) == 0 {
			a0 = a
			b0 = a
			trouve = true
			break
		} else {
			tmp3 := racineCarre(tmp2)
			if tmp3.Cmp(zero) != 0 {
				a1 := a
				b1 := tmp3
				a0 = big.NewInt(0)
				b0 = big.NewInt(0)
				a0.Add(a1, b1)
				b0.Sub(a1, b1)
				trouve = true
				break
			}
		}

		i = i.Add(i, un)
	}
	if trouve {
		fmt.Printf("trouve %s : %s * %s\n", n, a0, b0)
	} else {
		fmt.Printf("pas trouve (%s)\n", n)
	}
}

func main() {
	var n int64
	//var affiche bool

	n = 100
	n = 50000
	n = 100000
	n = 1_000_000
	n = 10_000_000
	//n = 1_000_000_000

	//affiche = false
	//affiche = true
	//methode := 5
	methode := 6

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
	} else if methode == 4 {
		filename := "d:\\temp\\double_primes_1_000_000.txt"
		filename2 := "d:\\temp\\double_primes_100_000_bis.txt"
		fmt.Printf("Enregistrement du fichier %s\n", filename2)
		debut := time.Now().UnixNano() / 1000_000
		filtreListeDoublePrime(5, filename, filename2)
		fin := time.Now().UnixNano() / 1000_000
		fmt.Printf("%d msec\n", fin-debut)
	} else if methode == 5 {
		filename := "d:\\temp\\primes2.txt"
		fmt.Printf("Enregistrement du fichier %s\n", filename)
		debut := time.Now().UnixNano() / 1000_000
		enregistreListePrime(n, filename)
		fin := time.Now().UnixNano() / 1000_000
		fmt.Printf("%d msec\n", fin-debut)
	} else if methode == 6 {
		//filename := "d:\\temp\\primes2.txt"
		//fmt.Printf("Enregistrement du fichier %s\n", filename)

		var s string
		//s = "15"  // 3*5
		//s = "115" // 5*23
		//s = "9409" // 97*97
		//s = "28741"                // 41*701
		//s = "99400891"             // 9967*9973
		//s = "2479541989"           //49789*49801 <1s
		s = "99998800003591" // <1s
		//s = "26382685634187504697" // 4248200851*6210319747 > plusieurs minutes

		x, res2 := new(big.Int).SetString(s, 10)
		if !res2 {
			println("Erreur")
		}
		debut := time.Now().UnixNano() / 1000_000
		factorisationFermat(x)
		fin := time.Now().UnixNano() / 1000_000
		fmt.Printf("%d msec\n", fin-debut)
	}
}
