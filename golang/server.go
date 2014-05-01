package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func root(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "OK")
}

func tetra(w http.ResponseWriter, r *http.Request) {
	val, _ := strconv.ParseFloat(r.URL.Path[len("/tetra/"):], 0)
	//bigint := big.NewInt(val).Exp(big.NewInt(val), big.NewInt(val), big.NewInt(1))
	ret := math.Pow(val, val)
	fmt.Fprintf(w, "%.0f", ret)
}

func available(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yes")
}

// Prime Factors (from http://edapx.com/2014/04/12/how-to-get-the-prime-factors-of-a-number-in-golang/)

func prime(w http.ResponseWriter, r *http.Request) {
	val, _ := strconv.ParseInt(r.URL.Path[len("/prime/"):], 0, 64)
	result := CalcPrimeFactors(int(val))
	fmt.Fprintf(w, "%v", result)
}

// Generate numbers until the limit max.
// after the 2, all the prime numbers are odd
// Send a channel signal when the limit is reached
func Generate(max int, ch chan<- int) {
	ch <- 2
	for i := 3; i <= max; i += 2 {
		ch <- i
	}
	ch <- -1 // signal that the limit is reached
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for i := <-in; i != -1; i = <-in {
		if i%prime != 0 {
			out <- i
		}
	}
	out <- -1
}

func CalcPrimeFactors(number_to_factorize int) []int {
	rv := []int{}
	ch := make(chan int)
	go Generate(number_to_factorize, ch)
	for prime := <-ch; (prime != -1) && (number_to_factorize > 1); prime = <-ch {
		for number_to_factorize%prime == 0 {
			number_to_factorize = number_to_factorize / prime
			rv = append(rv, prime)
		}
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return rv
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/tetra/", tetra)
	http.HandleFunc("/available/", available)
	http.HandleFunc("/prime/", prime)
	http.ListenAndServe(":3000", nil)
}
