package main

import (
	"fmt"
	"time"
)

func gcdIterative(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdRecursive(b, a%b)
}

func main() {
	var a, b int

	fmt.Print("Masukkan bilangan a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan b: ")
	fmt.Scan(&b)

	startIter := time.Now()
	gcdIter := gcdIterative(a, b)
	timeIter := time.Since(startIter)

	startRec := time.Now()
	gcdRec := gcdRecursive(a, b)
	timeRec := time.Since(startRec)

	fmt.Println("\n=== HASIL PERHITUNGAN GCD ===")
	fmt.Printf("GCD Iteratif : %d\n", gcdIter)
	fmt.Printf("Waktu Iteratif : %v\n", timeIter)

	fmt.Printf("\nGCD Rekursif : %d\n", gcdRec)
	fmt.Printf("Waktu Rekursif : %v\n", timeRec)
}
